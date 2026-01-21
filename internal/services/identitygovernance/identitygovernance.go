// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetAccessPackageResourcesRoleScope(ctx context.Context, client *entitlementmanagementaccesspackage.EntitlementManagementAccessPackageClient, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeId) (*beta.AccessPackageResourceRoleScope, error) {
	accessPackageId := beta.NewIdentityGovernanceEntitlementManagementAccessPackageID(id.AccessPackageId)
	options := entitlementmanagementaccesspackage.GetEntitlementManagementAccessPackageOperationOptions{
		Expand: &odata.Expand{
			Relationship: "accessPackageResourceRoleScopes($expand=accessPackageResourceRole,accessPackageResourceScope)",
		},
	}
	resp, err := client.GetEntitlementManagementAccessPackage(ctx, accessPackageId, options)
	if err != nil {
		return nil, fmt.Errorf("retrieving %s: %v", accessPackageId, err)
	}

	if resp.Model == nil {
		return nil, fmt.Errorf("retrieving %s: model was nil", accessPackageId)
	}

	if resp.Model.AccessPackageResourceRoleScopes == nil {
		return nil, fmt.Errorf("retrieving %s: AccessPackageResourceRoleScopes was nil", accessPackageId)
	}

	// There is only a select and expand method on this endpoint, we iterate the result to find the RoleScope
	for _, roleScope := range *resp.Model.AccessPackageResourceRoleScopes {
		if roleScope.Id != nil && *roleScope.Id == id.AccessPackageResourceRoleScopeId {
			return &roleScope, nil
		}
	}

	return nil, nil
}

func expandRequestorSettings(input []interface{}) (*beta.RequestorSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}
	in := input[0].(map[string]interface{})
	result := beta.RequestorSettings{
		ScopeType:      nullable.NoZero(in["scope_type"].(string)),
		AcceptRequests: nullable.NoZero(in["requests_accepted"].(bool)),
	}

	if _, ok := in["requestor"]; ok {
		allowedRequestors, err := expandUserSets(in["requestor"].([]interface{}))
		if err != nil {
			return nil, fmt.Errorf("building `requestor`: %v", err)
		}
		result.AllowedRequestors = allowedRequestors
	}

	return &result, nil
}

func flattenRequestorSettings(input *beta.RequestorSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"requests_accepted": input.AcceptRequests.GetOrZero(),
		"scope_type":        input.ScopeType.GetOrZero(),
		"requestor":         flattenUserSets(input.AllowedRequestors),
	}}
}

func expandApprovalSettings(input []interface{}) (*beta.ApprovalSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}

	in := input[0].(map[string]interface{})

	result := beta.ApprovalSettings{
		IsApprovalRequired:               nullable.Value(in["approval_required"].(bool)),
		IsApprovalRequiredForExtension:   nullable.Value(in["approval_required_for_extension"].(bool)),
		IsRequestorJustificationRequired: nullable.Value(in["requestor_justification_required"].(bool)),
	}

	approvalStages := make([]beta.ApprovalStage, 0)
	if _, ok := in["approval_stage"]; ok {
		for _, raw := range in["approval_stage"].([]interface{}) {
			v := raw.(map[string]interface{})

			stage := beta.ApprovalStage{
				ApprovalStageTimeOutInDays:      nullable.NoZero(int64(v["approval_timeout_in_days"].(int))),
				EscalationTimeInMinutes:         nullable.NoZero(int64(v["enable_alternative_approval_in_days"].(int) * 24 * 60)),
				IsApproverJustificationRequired: nullable.NoZero(v["approver_justification_required"].(bool)),
				IsEscalationEnabled:             nullable.NoZero(v["alternative_approval_enabled"].(bool)),
			}

			if _, ok := v["primary_approver"]; ok {
				primaryApprovers, err := expandUserSets(v["primary_approver"].([]interface{}))
				if err != nil {
					return nil, fmt.Errorf("building `primary_approver`: %v", err)
				}
				stage.PrimaryApprovers = primaryApprovers
			}

			if _, ok := v["alternative_approver"]; ok {
				escalationApprovers, err := expandUserSets(v["alternative_approver"].([]interface{}))
				if err != nil {
					return nil, fmt.Errorf("building `alternative_approver`: %v", err)
				}
				stage.EscalationApprovers = escalationApprovers
			}

			approvalStages = append(approvalStages, stage)
		}
	}

	result.ApprovalStages = &approvalStages

	return &result, nil
}

func flattenApprovalSettings(input *beta.ApprovalSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	result := []map[string]interface{}{{
		"approval_required":                input.IsApprovalRequired.GetOrZero(),
		"approval_required_for_extension":  input.IsApprovalRequiredForExtension.GetOrZero(),
		"requestor_justification_required": input.IsRequestorJustificationRequired.GetOrZero(),
	}}

	approvalStages := make([]interface{}, 0)
	for _, v := range *input.ApprovalStages {
		var alternativeApprovalInDays int
		if w := v.EscalationTimeInMinutes.GetOrZero(); w > 0 {
			alternativeApprovalInDays = int(w) / 60 / 24
		}
		approvalStage := map[string]interface{}{
			"approval_timeout_in_days":            int(v.ApprovalStageTimeOutInDays.GetOrZero()),
			"approver_justification_required":     v.IsApproverJustificationRequired.GetOrZero(),
			"alternative_approval_enabled":        v.IsEscalationEnabled.GetOrZero(),
			"enable_alternative_approval_in_days": alternativeApprovalInDays,
			"primary_approver":                    flattenUserSets(v.PrimaryApprovers),
			"alternative_approver":                flattenUserSets(v.EscalationApprovers),
		}

		approvalStages = append(approvalStages, approvalStage)
	}

	result[0]["approval_stage"] = approvalStages

	return result
}

func expandAssignmentReviewSettings(input []interface{}) (*beta.AssignmentReviewSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}

	in := input[0].(map[string]interface{})

	result := beta.AssignmentReviewSettings{
		AccessReviewTimeoutBehavior:     pointer.To(beta.AccessReviewTimeoutBehavior(in["access_review_timeout_behavior"].(string))),
		DurationInDays:                  nullable.Value(int64(in["duration_in_days"].(int))),
		IsAccessRecommendationEnabled:   nullable.Value(in["access_recommendation_enabled"].(bool)),
		IsApprovalJustificationRequired: nullable.Value(in["approver_justification_required"].(bool)),
		IsEnabled:                       nullable.Value(in["enabled"].(bool)),
		RecurrenceType:                  nullable.NoZero(in["review_frequency"].(string)),
		ReviewerType:                    nullable.NoZero(in["review_type"].(string)),
		StartDateTime:                   nullable.NoZero(in["starting_on"].(string)),
	}

	if _, ok := in["reviewer"]; ok {
		reviewers, err := expandUserSets(in["reviewer"].([]interface{}))
		if err != nil {
			return nil, fmt.Errorf("building `reviewer`: %v", err)
		}
		result.Reviewers = reviewers
	}

	if pointer.From(result.AccessReviewTimeoutBehavior) == "" && result.DurationInDays.GetOrZero() == 0 &&
		!result.IsAccessRecommendationEnabled.GetOrZero() && !result.IsApprovalJustificationRequired.GetOrZero() &&
		!result.IsEnabled.GetOrZero() && result.RecurrenceType.GetOrZero() == "" &&
		result.ReviewerType.GetOrZero() == "" && len(pointer.From(result.Reviewers)) == 0 {
		return nil, nil
	}

	return &result, nil
}

func flattenAssignmentReviewSettings(input *beta.AssignmentReviewSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	return []map[string]interface{}{{
		"access_recommendation_enabled":   input.IsAccessRecommendationEnabled.GetOrZero(),
		"access_review_timeout_behavior":  pointer.From(input.AccessReviewTimeoutBehavior),
		"approver_justification_required": input.IsApprovalJustificationRequired.GetOrZero(),
		"duration_in_days":                input.DurationInDays.GetOrZero(),
		"enabled":                         input.IsEnabled.GetOrZero(),
		"review_frequency":                input.RecurrenceType.GetOrZero(),
		"review_type":                     input.ReviewerType.GetOrZero(),
		"reviewer":                        flattenUserSets(input.Reviewers),
		"starting_on":                     input.StartDateTime.GetOrZero(),
	}}
}

func expandUserSets(input []interface{}) (*[]beta.UserSet, error) {
	userSets := make([]beta.UserSet, 0)
	for _, raw := range input {
		v := raw.(map[string]interface{})

		isBackup := v["backup"].(bool)
		objectId := v["object_id"].(string)
		odataType := formatODataType(v["subject_type"].(string))

		var userSet beta.UserSet
		switch odataType {
		case "ConnectedOrganizationMembers":
			userSet = beta.ConnectedOrganizationMembers{
				Id:       nullable.Value(objectId),
				IsBackup: nullable.Value(isBackup),
			}
		case "ExternalSponsors":
			userSet = beta.ExternalSponsors{
				IsBackup: nullable.Value(isBackup),
			}
		case "GroupMembers":
			userSet = beta.GroupMembers{
				Id:       nullable.Value(objectId),
				IsBackup: nullable.Value(isBackup),
			}
		case "InternalSponsors":
			userSet = beta.InternalSponsors{
				IsBackup: nullable.Value(isBackup),
			}
		case "RequestorManager":
			userSet = beta.RequestorManager{
				IsBackup: nullable.Value(isBackup),
			}
		case "SingleUser":
			userSet = beta.SingleUser{
				Id:       nullable.Value(objectId),
				IsBackup: nullable.Value(isBackup),
			}
		case "TargetUserSponsors":
			userSet = beta.TargetUserSponsors{
				IsBackup: nullable.Value(isBackup),
			}
		default:
			return nil, fmt.Errorf("unknown `subject_type`: %s", odataType)
		}

		userSets = append(userSets, userSet)
	}

	return &userSets, nil
}

func flattenUserSets(input *[]beta.UserSet) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	userSets := make([]map[string]interface{}, 0)
	for _, raw := range *input {
		v := raw.UserSet()
		var id *string

		switch impl := raw.(type) {
		case beta.ConnectedOrganizationMembers:
			id = impl.Id.Get()
		case beta.GroupMembers:
			id = impl.Id.Get()
		case beta.SingleUser:
			id = impl.Id.Get()
		}

		userSet := map[string]interface{}{
			"subject_type": formatODataType(pointer.From(v.ODataType)),
			"backup":       v.IsBackup.GetOrZero(),
			"object_id":    pointer.From(id),
		}

		userSets = append(userSets, userSet)
	}

	return userSets
}

func expandAccessPackageQuestions(input []interface{}) *[]beta.AccessPackageQuestion {
	result := make([]beta.AccessPackageQuestion, 0)

	for _, raw := range input {
		v := raw.(map[string]interface{})

		textList := v["text"].([]interface{})
		if len(textList) == 0 {
			continue
		}
		text := textList[0].(map[string]interface{})

		var question beta.AccessPackageQuestion

		if choicesRaw, ok := v["choice"].([]interface{}); ok && len(choicesRaw) > 0 {
			choices := make([]beta.AccessPackageAnswerChoice, 0)
			for _, choiceRaw := range choicesRaw {
				choice := choiceRaw.(map[string]interface{})
				displayValue := make(map[string]interface{})
				if v := choice["display_value"].([]interface{}); len(v) > 0 {
					displayValue = v[0].(map[string]interface{})
				}

				choices = append(choices, beta.AccessPackageAnswerChoice{
					ActualValue:  nullable.NoZero(choice["actual_value"].(string)),
					DisplayValue: pointer.From(expandAccessPackageLocalizedContent(displayValue)),
				})
			}

			question = beta.AccessPackageMultipleChoiceQuestion{
				Choices:    &choices,
				IsRequired: nullable.Value(v["required"].(bool)),
				Sequence:   nullable.Value(int64(v["sequence"].(int))),
				Text:       expandAccessPackageLocalizedContent(text),
			}
		} else {
			question = beta.AccessPackageTextInputQuestion{
				IsRequired: nullable.Value(v["required"].(bool)),
				Sequence:   nullable.Value(int64(v["sequence"].(int))),
				Text:       expandAccessPackageLocalizedContent(text),
			}
		}

		result = append(result, question)
	}

	return &result
}

func flattenAccessPackageQuestions(input *[]beta.AccessPackageQuestion) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	questions := make([]map[string]interface{}, 0)

	for _, raw := range *input {
		v := raw.AccessPackageQuestion()

		question := map[string]interface{}{
			"required": v.IsRequired.GetOrZero(),
			"sequence": int(v.Sequence.GetOrZero()),
			"text":     flattenAccessPackageLocalizedContent(v.Text),
		}

		if impl, ok := raw.(beta.AccessPackageMultipleChoiceQuestion); ok {
			choices := make([]map[string]interface{}, 0)
			for _, choice := range pointer.From(impl.Choices) {
				choices = append(choices, map[string]interface{}{
					"actual_value":  choice.ActualValue.GetOrZero(),
					"display_value": flattenAccessPackageLocalizedContent(&choice.DisplayValue),
				})
			}
			question["choice"] = choices
		}

		questions = append(questions, question)
	}

	return questions
}

func expandAccessPackageLocalizedContent(input map[string]interface{}) *beta.AccessPackageLocalizedContent {
	if len(input) == 0 {
		return nil
	}

	result := beta.AccessPackageLocalizedContent{
		DefaultText: nullable.NoZero(input["default_text"].(string)),
	}

	texts := make([]beta.AccessPackageLocalizedText, 0)

	if _, ok := input["localized_text"]; ok {
		for _, raw := range input["localized_text"].([]interface{}) {
			v := raw.(map[string]interface{})

			texts = append(texts, beta.AccessPackageLocalizedText{
				LanguageCode: nullable.NoZero(v["language_code"].(string)),
				Text:         nullable.NoZero(v["content"].(string)),
			})
		}
	}

	result.LocalizedTexts = &texts

	return &result
}

func flattenAccessPackageLocalizedContent(input *beta.AccessPackageLocalizedContent) []map[string]interface{} {
	result := []map[string]interface{}{{
		"default_text": input.DefaultText.GetOrZero(),
	}}

	texts := make([]map[string]interface{}, 0)

	for _, v := range *input.LocalizedTexts {
		text := map[string]interface{}{
			"language_code": v.LanguageCode.GetOrZero(),
			"content":       v.Text.GetOrZero(),
		}

		texts = append(texts, text)
	}

	result[0]["localized_text"] = texts

	return result
}

func formatODataType(in string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimPrefix(in, "#microsoft.graph."))
}
