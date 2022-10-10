package identitygovernance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
	"github.com/manicminer/hamilton/odata"
)

const idDelimitor = ":"

func buildAssignmentPolicyRequestorSettings(input []interface{}) *msgraph.RequestorSettings {
	if len(input) == 0 {
		return nil
	}
	in := input[0].(map[string]interface{})
	result := msgraph.RequestorSettings{
		ScopeType:      in["scope_type"].(string),
		AcceptRequests: utils.Bool(in["accept_requests"].(bool)),
	}
	result.AllowedRequestors = buildUserSet(in["requestor"].([]interface{}))

	return &result
}

func flattenRequestorSettings(input *msgraph.RequestorSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"accept_requests": input.AcceptRequests,
		"scope_type":      input.ScopeType,
		"requestor":       flattenUserSet(input.AllowedRequestors),
	}}
}

func buildAssignmentPolicyApprovalSettings(input []interface{}) *msgraph.ApprovalSettings {
	if len(input) == 0 {
		return nil
	}
	in := input[0].(map[string]interface{})
	result := msgraph.ApprovalSettings{
		IsApprovalRequired:               utils.Bool(in["is_approval_required"].(bool)),
		IsApprovalRequiredForExtension:   utils.Bool(in["is_approval_required_for_extension"].(bool)),
		IsRequestorJustificationRequired: utils.Bool(in["is_requestor_justification_required"].(bool)),
	}
	approvalStages := make([]msgraph.ApprovalStage, 0)
	for _, v := range in["approval_stage"].([]interface{}) {
		v_map := v.(map[string]interface{})
		stage := msgraph.ApprovalStage{
			ApprovalStageTimeOutInDays:      utils.Int32(int32(v_map["approval_timeout_in_days"].(int))),
			IsApproverJustificationRequired: utils.Bool(v_map["is_approver_justification_required"].(bool)),
			IsEscalationEnabled:             utils.Bool(v_map["is_alternative_approval_enabled"].(bool)),
			EscalationTimeInMinutes:         utils.Int32((int32(v_map["enable_alternative_approval_in_days"].(int) * 24 * 60))),
		}
		stage.PrimaryApprovers = buildUserSet(v_map["primary_approver"].([]interface{}))
		stage.EscalationApprovers = buildUserSet(v_map["alternative_approver"].([]interface{}))

		approvalStages = append(approvalStages, stage)
	}
	result.ApprovalStages = &approvalStages

	return &result
}

func falttenApprovalSettings(input *msgraph.ApprovalSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	result := []map[string]interface{}{{
		"is_approval_required":                input.IsApprovalRequired,
		"is_approval_required_for_extension":  input.IsApprovalRequiredForExtension,
		"is_requestor_justification_required": input.IsRequestorJustificationRequired,
	}}

	approvalStages := make([]interface{}, 0)
	for _, v := range *input.ApprovalStages {
		approvalStage := map[string]interface{}{
			"approval_timeout_in_days":            v.ApprovalStageTimeOutInDays,
			"is_approver_justification_required":  v.IsApproverJustificationRequired,
			"is_alternative_approval_enabled":     v.IsEscalationEnabled,
			"enable_alternative_approval_in_days": *v.EscalationTimeInMinutes / 60 / 24,
			"primary_approver":                    flattenUserSet(v.PrimaryApprovers),
			"alternative_approver":                flattenUserSet(v.EscalationApprovers),
		}

		approvalStages = append(approvalStages, approvalStage)
	}
	result[0]["approval_stage"] = approvalStages

	return result
}

func buildAssignmentPolicyReviewSettings(input []interface{}) (*msgraph.AssignmentReviewSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}
	in := input[0].(map[string]interface{})

	result := msgraph.AssignmentReviewSettings{
		IsEnabled:                       utils.Bool(in["is_enabled"].(bool)),
		RecurrenceType:                  in["review_frequency"].(string),
		ReviewerType:                    in["review_type"].(string),
		DurationInDays:                  utils.Int32(int32(in["duration_in_days"].(int))),
		IsAccessRecommendationEnabled:   utils.Bool(in["is_access_recommendation_enabled"].(bool)),
		IsApprovalJustificationRequired: utils.Bool(in["is_approver_justification_required"].(bool)),
		AccessReviewTimeoutBehavior:     in["access_review_timeout_behavior"].(string),
	}

	startOnDate := in["starting_on"].(string)
	if startOnDate != "" {
		startOn, err := time.Parse(time.RFC3339, startOnDate)
		if err != nil {
			return nil, fmt.Errorf("Error converting starting date %q to a valid date: %q", in["starting_on"].(string), err)
		}
		result.StartDateTime = &startOn
	}

	result.Reviewers = buildUserSet(in["reviewer"].([]interface{}))

	return &result, nil
}

func flattenReviewSettings(input *msgraph.AssignmentReviewSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"is_enabled":                         input.IsEnabled,
		"review_frequency":                   input.RecurrenceType,
		"review_type":                        input.ReviewerType,
		"starting_on":                        input.StartDateTime.Format(time.RFC3339),
		"duration_in_days":                   input.DurationInDays,
		"reviewer":                           flattenUserSet(input.Reviewers),
		"is_access_recommendation_enabled":   input.IsAccessRecommendationEnabled,
		"is_approver_justification_required": input.IsApprovalJustificationRequired,
		"access_review_timeout_behavior":     input.AccessReviewTimeoutBehavior,
	}}
}

func buildUserSet(input []interface{}) *[]msgraph.UserSet {
	userSets := make([]msgraph.UserSet, 0)
	for _, v := range input {
		v_map := v.(map[string]interface{})
		userSet := msgraph.UserSet{
			ODataType: userSetODataType(v_map["subject_type"].(string)),
			ID:        utils.String(v_map["object_id"].(string)),
			IsBackup:  utils.Bool(v_map["is_backup"].(bool)),
		}

		userSets = append(userSets, userSet)
	}

	return &userSets
}

func flattenUserSet(input *[]msgraph.UserSet) []interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	userSets := make([]interface{}, 0)
	for _, v := range *input {
		userSet := map[string]interface{}{
			"subject_type": userSetShortType(*v.ODataType),
			"is_backup":    v.IsBackup,
			"object_id":    v.ID,
		}
		userSets = append(userSets, userSet)
	}
	return userSets
}

func userSetODataType(in string) *string {
	odataType := odata.TypeSingleUser
	switch in {
	case odata.ShortTypeGroupMembers:
		odataType = odata.TypeGroupMembers
	case odata.ShortTypeConnectedOrganizationMembers:
		odataType = odata.TypeConnectedOrganizationMembers
	case odata.ShortTypeRequestorManager:
		odataType = odata.TypeRequestorManager
	case odata.ShortTypeInternalSponsors:
		odataType = odata.TypeInternalSponsors
	case odata.ShortTypeExternalSponsors:
		odataType = odata.TypeExternalSponsors
	}

	return &odataType
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

func expandAccessPakcageAssignmentPolicyQuestions(questions []interface{}) *[]msgraph.AccessPackageQuestion {
	result := make([]msgraph.AccessPackageQuestion, 0)

	for _, v := range questions {
		v_map := v.(map[string]interface{})
		v_text_list := v_map["text"].([]interface{})
		v_text := v_text_list[0].(map[string]interface{})

		q := msgraph.AccessPackageQuestion{
			IsRequired: utils.Bool(v_map["is_required"].(bool)),
			Sequence:   utils.Int32(int32(v_map["sequence"].(int))),
			Text:       expandAccessPakcageAssignmentPolicyQuestionContent(v_text),
		}

		v_map_chocies := v_map["choice"].([]interface{})
		q.ODataType = utils.String(odata.TypeAccessPackageTextInputQuestion)
		if len(v_map_chocies) > 0 {
			q.ODataType = utils.String(odata.TypeAccessPackageMultipleChoiceQuestion)
			choices := make([]msgraph.AccessPackageMultipleChoiceQuestions, 0)
			for _, c := range v_map_chocies {
				c_map := c.(map[string]interface{})
				c_map_display_value := c_map["display_value"].([]interface{})
				choices = append(choices, msgraph.AccessPackageMultipleChoiceQuestions{
					ActualValue:  utils.String(c_map["actual_value"].(string)),
					DisplayValue: expandAccessPakcageAssignmentPolicyQuestionContent(c_map_display_value[0].(map[string]interface{})),
				})
			}
			q.Choices = &choices
		}

		result = append(result, q)
	}

	return &result
}

func flattenAssignmentPolicyQuestions(input *[]msgraph.AccessPackageQuestion) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	questions := make([]map[string]interface{}, 0)
	for _, v := range *input {
		question := map[string]interface{}{
			"is_required": v.IsRequired,
			"sequence":    v.Sequence,
			"text":        flattenAssignmentPolicyQuestionContent(v.Text),
		}

		if c_array := v.Choices; c_array != nil && len(*c_array) > 0 {
			choices := make([]map[string]interface{}, 0)
			for _, c := range *c_array {
				choice := map[string]interface{}{
					"actual_value":  c.ActualValue,
					"display_value": flattenAssignmentPolicyQuestionContent(c.DisplayValue),
				}

				choices = append(choices, choice)
			}
			question["choice"] = choices
		}

		questions = append(questions, question)
	}

	return questions
}

func expandAccessPakcageAssignmentPolicyQuestionContent(input map[string]interface{}) *msgraph.AccessPackageLocalizedContent {
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

func flattenAssignmentPolicyQuestionContent(input *msgraph.AccessPackageLocalizedContent) []map[string]interface{} {
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
