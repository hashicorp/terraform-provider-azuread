// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func expandRequestorSettings(input []interface{}) *msgraph.RequestorSettings {
	if len(input) == 0 {
		return nil
	}
	in := input[0].(map[string]interface{})
	result := msgraph.RequestorSettings{
		ScopeType:      in["scope_type"].(string),
		AcceptRequests: utils.Bool(in["requests_accepted"].(bool)),
	}
	result.AllowedRequestors = expandUserSets(in["requestor"].([]interface{}))

	return &result
}

func flattenRequestorSettings(input *msgraph.RequestorSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"requests_accepted": input.AcceptRequests,
		"scope_type":        input.ScopeType,
		"requestor":         flattenUserSets(input.AllowedRequestors),
	}}
}

func expandApprovalSettings(input []interface{}) *msgraph.ApprovalSettings {
	if len(input) == 0 {
		return nil
	}

	in := input[0].(map[string]interface{})

	result := msgraph.ApprovalSettings{
		IsApprovalRequired:               utils.Bool(in["approval_required"].(bool)),
		IsApprovalRequiredForExtension:   utils.Bool(in["approval_required_for_extension"].(bool)),
		IsRequestorJustificationRequired: utils.Bool(in["requestor_justification_required"].(bool)),
	}

	approvalStages := make([]msgraph.ApprovalStage, 0)
	for _, v := range in["approval_stage"].([]interface{}) {
		v_map := v.(map[string]interface{})

		stage := msgraph.ApprovalStage{
			ApprovalStageTimeOutInDays:      utils.Int32(int32(v_map["approval_timeout_in_days"].(int))),
			EscalationTimeInMinutes:         utils.Int32(int32(v_map["enable_alternative_approval_in_days"].(int) * 24 * 60)),
			IsApproverJustificationRequired: utils.Bool(v_map["approver_justification_required"].(bool)),
			IsEscalationEnabled:             utils.Bool(v_map["alternative_approval_enabled"].(bool)),
		}

		stage.PrimaryApprovers = expandUserSets(v_map["primary_approver"].([]interface{}))
		stage.EscalationApprovers = expandUserSets(v_map["alternative_approver"].([]interface{}))

		approvalStages = append(approvalStages, stage)
	}

	result.ApprovalStages = &approvalStages

	return &result
}

func flattenApprovalSettings(input *msgraph.ApprovalSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	result := []map[string]interface{}{{
		"approval_required":                input.IsApprovalRequired,
		"approval_required_for_extension":  input.IsApprovalRequiredForExtension,
		"requestor_justification_required": input.IsRequestorJustificationRequired,
	}}

	approvalStages := make([]interface{}, 0)
	for _, v := range *input.ApprovalStages {
		approvalStage := map[string]interface{}{
			"approval_timeout_in_days":            v.ApprovalStageTimeOutInDays,
			"approver_justification_required":     v.IsApproverJustificationRequired,
			"alternative_approval_enabled":        v.IsEscalationEnabled,
			"enable_alternative_approval_in_days": *v.EscalationTimeInMinutes / 60 / 24,
			"primary_approver":                    flattenUserSets(v.PrimaryApprovers),
			"alternative_approver":                flattenUserSets(v.EscalationApprovers),
		}

		approvalStages = append(approvalStages, approvalStage)
	}
	result[0]["approval_stage"] = approvalStages

	return result
}

func expandAssignmentReviewSettings(input []interface{}) (*msgraph.AssignmentReviewSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}

	in := input[0].(map[string]interface{})

	result := msgraph.AssignmentReviewSettings{
		AccessReviewTimeoutBehavior:     in["access_review_timeout_behavior"].(string),
		DurationInDays:                  utils.Int32(int32(in["duration_in_days"].(int))),
		IsAccessRecommendationEnabled:   utils.Bool(in["access_recommendation_enabled"].(bool)),
		IsApprovalJustificationRequired: utils.Bool(in["approver_justification_required"].(bool)),
		IsEnabled:                       utils.Bool(in["enabled"].(bool)),
		RecurrenceType:                  in["review_frequency"].(string),
		ReviewerType:                    in["review_type"].(string),
	}

	startOnDate := in["starting_on"].(string)
	if startOnDate != "" {
		startOn, err := time.Parse(time.RFC3339, startOnDate)
		if err != nil {
			return nil, fmt.Errorf("converting starting date %q to a valid date: %q", in["starting_on"].(string), err)
		}

		result.StartDateTime = &startOn
	}

	result.Reviewers = expandUserSets(in["reviewer"].([]interface{}))

	return &result, nil
}

func flattenAssignmentReviewSettings(input *msgraph.AssignmentReviewSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"access_recommendation_enabled":   input.IsAccessRecommendationEnabled,
		"access_review_timeout_behavior":  input.AccessReviewTimeoutBehavior,
		"approver_justification_required": input.IsApprovalJustificationRequired,
		"duration_in_days":                input.DurationInDays,
		"enabled":                         input.IsEnabled,
		"review_frequency":                input.RecurrenceType,
		"review_type":                     input.ReviewerType,
		"reviewer":                        flattenUserSets(input.Reviewers),
		"starting_on":                     input.StartDateTime.Format(time.RFC3339),
	}}
}

func expandUserSets(input []interface{}) *[]msgraph.UserSet {
	userSets := make([]msgraph.UserSet, 0)
	for _, v := range input {
		v_map := v.(map[string]interface{})
		oDataType, needId := userSetODataType(v_map["subject_type"].(string))
		userSet := msgraph.UserSet{
			ODataType: oDataType,
			IsBackup:  utils.Bool(v_map["backup"].(bool)),
		}
		if needId {
			userSet.ID = utils.String(v_map["object_id"].(string))
		}

		userSets = append(userSets, userSet)
	}

	return &userSets
}

func flattenUserSets(input *[]msgraph.UserSet) []interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	userSets := make([]interface{}, 0)
	for _, v := range *input {
		userSet := map[string]interface{}{
			"subject_type": userSetShortType(*v.ODataType),
			"backup":       v.IsBackup,
			"object_id":    v.ID,
		}

		userSets = append(userSets, userSet)
	}

	return userSets
}

func userSetODataType(in string) (*string, bool) {
	odataType := odata.TypeSingleUser
	needId := true
	switch in {
	case odata.ShortTypeGroupMembers:
		odataType = odata.TypeGroupMembers
	case odata.ShortTypeConnectedOrganizationMembers:
		odataType = odata.TypeConnectedOrganizationMembers
	case odata.ShortTypeRequestorManager:
		odataType = odata.TypeRequestorManager
		needId = false
	case odata.ShortTypeInternalSponsors:
		odataType = odata.TypeInternalSponsors
		needId = false
	case odata.ShortTypeExternalSponsors:
		odataType = odata.TypeExternalSponsors
		needId = false
	}

	return &odataType, needId
}

func userSetShortType(in string) *string {
	shortType := odata.ShortTypeSingleUser
	switch in {
	case odata.TypeGroupMembers:
		shortType = odata.ShortTypeGroupMembers
	case odata.TypeConnectedOrganizationMembers:
		shortType = odata.ShortTypeConnectedOrganizationMembers
	case odata.TypeRequestorManager:
		shortType = odata.ShortTypeRequestorManager
	case odata.TypeInternalSponsors:
		shortType = odata.ShortTypeInternalSponsors
	case odata.TypeExternalSponsors:
		shortType = odata.ShortTypeExternalSponsors
	}

	return &shortType
}

func expandAccessPackageQuestions(questions []interface{}) *[]msgraph.AccessPackageQuestion {
	result := make([]msgraph.AccessPackageQuestion, 0)

	for _, v := range questions {
		v_map := v.(map[string]interface{})
		v_text_list := v_map["text"].([]interface{})
		v_text := v_text_list[0].(map[string]interface{})

		q := msgraph.AccessPackageQuestion{
			IsRequired: utils.Bool(v_map["required"].(bool)),
			Sequence:   utils.Int32(int32(v_map["sequence"].(int))),
			Text:       expandAccessPackageLocalizedContent(v_text),
		}

		v_map_choices := v_map["choice"].([]interface{})
		q.ODataType = utils.String(odata.TypeAccessPackageTextInputQuestion)
		if len(v_map_choices) > 0 {
			q.ODataType = utils.String(odata.TypeAccessPackageMultipleChoiceQuestion)
			choices := make([]msgraph.AccessPackageMultipleChoiceQuestions, 0)

			for _, c := range v_map_choices {
				c_map := c.(map[string]interface{})
				c_map_display_value := c_map["display_value"].([]interface{})
				choices = append(choices, msgraph.AccessPackageMultipleChoiceQuestions{
					ActualValue:  utils.String(c_map["actual_value"].(string)),
					DisplayValue: expandAccessPackageLocalizedContent(c_map_display_value[0].(map[string]interface{})),
				})
			}

			q.Choices = &choices
		}

		result = append(result, q)
	}

	return &result
}

func flattenAccessPackageQuestions(input *[]msgraph.AccessPackageQuestion) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	questions := make([]map[string]interface{}, 0)

	for _, v := range *input {
		question := map[string]interface{}{
			"required": v.IsRequired,
			"sequence": v.Sequence,
			"text":     flattenAccessPackageLocalizedContent(v.Text),
		}

		if c_array := v.Choices; c_array != nil && len(*c_array) > 0 {
			choices := make([]map[string]interface{}, 0)

			for _, c := range *c_array {
				choice := map[string]interface{}{
					"actual_value":  c.ActualValue,
					"display_value": flattenAccessPackageLocalizedContent(c.DisplayValue),
				}

				choices = append(choices, choice)
			}

			question["choice"] = choices
		}

		questions = append(questions, question)
	}

	return questions
}

func expandAccessPackageLocalizedContent(input map[string]interface{}) *msgraph.AccessPackageLocalizedContent {
	result := msgraph.AccessPackageLocalizedContent{
		DefaultText: utils.String(input["default_text"].(string)),
	}

	texts := make([]msgraph.AccessPackageLocalizedTexts, 0)

	for _, v := range input["localized_text"].([]interface{}) {
		v_map := v.(map[string]interface{})
		texts = append(texts, msgraph.AccessPackageLocalizedTexts{
			LanguageCode: utils.String(v_map["language_code"].(string)),
			Text:         utils.String(v_map["content"].(string)),
		})
	}

	result.LocalizedTexts = &texts

	return &result
}

func flattenAccessPackageLocalizedContent(input *msgraph.AccessPackageLocalizedContent) []map[string]interface{} {
	result := []map[string]interface{}{{
		"default_text": input.DefaultText,
	}}

	texts := make([]map[string]interface{}, 0)

	for _, v := range *input.LocalizedTexts {
		text := map[string]interface{}{
			"language_code": v.LanguageCode,
			"content":       v.Text,
		}

		texts = append(texts, text)
	}

	result[0]["localized_text"] = texts

	return result
}
