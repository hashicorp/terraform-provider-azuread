// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
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

// scopeTypeToAllowedTargetScope maps the existing Terraform schema scope_type values (using beta naming)
// to the stable API AllowedTargetScope enum. Unknown values are passed through as-is.
func scopeTypeToAllowedTargetScope(scopeType string) stable.AllowedTargetScope {
	mapping := map[string]stable.AllowedTargetScope{
		"AllConfiguredConnectedOrganizationSubjects": stable.AllowedTargetScope_AllConfiguredConnectedOrganizationUsers,
		"AllExistingConnectedOrganizationSubjects":   stable.AllowedTargetScope_AllConfiguredConnectedOrganizationUsers,
		"AllExistingDirectoryMemberUsers":            stable.AllowedTargetScope_AllMemberUsers,
		"AllExistingDirectorySubjects":               stable.AllowedTargetScope_AllDirectoryUsers,
		"AllExternalSubjects":                        stable.AllowedTargetScope_AllExternalUsers,
		"NoSubjects":                                 stable.AllowedTargetScope_NotSpecified,
		"SpecificConnectedOrganizationSubjects":      stable.AllowedTargetScope_SpecificConnectedOrganizationUsers,
		"SpecificDirectorySubjects":                  stable.AllowedTargetScope_SpecificDirectoryUsers,
	}
	if v, ok := mapping[scopeType]; ok {
		return v
	}
	return stable.AllowedTargetScope(scopeType)
}

// allowedTargetScopeToScopeType maps stable API AllowedTargetScope enum values back to the Terraform
// schema scope_type values used in the resource schema.
func allowedTargetScopeToScopeType(scope stable.AllowedTargetScope) string {
	mapping := map[stable.AllowedTargetScope]string{
		stable.AllowedTargetScope_AllConfiguredConnectedOrganizationUsers: "AllConfiguredConnectedOrganizationSubjects",
		stable.AllowedTargetScope_AllMemberUsers:                          "AllExistingDirectoryMemberUsers",
		stable.AllowedTargetScope_AllDirectoryUsers:                       "AllExistingDirectorySubjects",
		stable.AllowedTargetScope_AllExternalUsers:                        "AllExternalSubjects",
		stable.AllowedTargetScope_NotSpecified:                            "NoSubjects",
		stable.AllowedTargetScope_SpecificConnectedOrganizationUsers:      "SpecificConnectedOrganizationSubjects",
		stable.AllowedTargetScope_SpecificDirectoryUsers:                  "SpecificDirectorySubjects",
	}
	if v, ok := mapping[scope]; ok {
		return v
	}
	return string(scope)
}

// reviewFrequencyToRecurrencePattern maps Terraform review_frequency values to the stable recurrence
// pattern type and interval.
func reviewFrequencyToRecurrencePattern(freq string) (stable.RecurrencePatternType, int64) {
	switch freq {
	case "weekly":
		return stable.RecurrencePatternType_Weekly, 1
	case "monthly":
		return stable.RecurrencePatternType_AbsoluteMonthly, 1
	case "quarterly":
		return stable.RecurrencePatternType_AbsoluteMonthly, 3
	case "halfyearly":
		return stable.RecurrencePatternType_AbsoluteMonthly, 6
	case "annual":
		return stable.RecurrencePatternType_AbsoluteYearly, 1
	}
	return stable.RecurrencePatternType_Weekly, 1
}

// recurrencePatternToReviewFrequency maps stable recurrence pattern type + interval back to the
// Terraform review_frequency schema value.
func recurrencePatternToReviewFrequency(patternType stable.RecurrencePatternType, interval int64) string {
	switch patternType {
	case stable.RecurrencePatternType_Weekly:
		return "weekly"
	case stable.RecurrencePatternType_AbsoluteMonthly:
		switch interval {
		case 3:
			return "quarterly"
		case 6:
			return "halfyearly"
		default:
			return "monthly"
		}
	case stable.RecurrencePatternType_AbsoluteYearly:
		return "annual"
	}
	return "weekly"
}

// expandRequestorSettings maps the requestor_settings Terraform block to the stable API fields
// AllowedTargetScope and SpecificAllowedTargets on the AccessPackageAssignmentPolicy.
// requests_accepted has no stable equivalent and is silently ignored.
func expandRequestorSettings(input []interface{}) (*stable.AllowedTargetScope, *[]stable.SubjectSet, error) {
	if len(input) == 0 {
		return nil, nil, nil
	}
	in := input[0].(map[string]interface{})

	scope := scopeTypeToAllowedTargetScope(in["scope_type"].(string))

	var targets *[]stable.SubjectSet
	if _, ok := in["requestor"]; ok {
		var err error
		targets, err = expandSubjectSets(in["requestor"].([]interface{}))
		if err != nil {
			return nil, nil, fmt.Errorf("building `requestor`: %v", err)
		}
	}

	return &scope, targets, nil
}

// flattenRequestorSettings maps the stable API AllowedTargetScope and SpecificAllowedTargets fields
// back to the requestor_settings Terraform block. requests_accepted always returns false.
func flattenRequestorSettings(scope *stable.AllowedTargetScope, targets *[]stable.SubjectSet) []map[string]interface{} {
	scopeType := ""
	if scope != nil {
		scopeType = allowedTargetScopeToScopeType(*scope)
	}
	return []map[string]interface{}{{
		"requests_accepted": false,
		"scope_type":        scopeType,
		"requestor":         flattenSubjectSets(targets),
	}}
}

func expandApprovalSettings(input []interface{}) (*stable.AccessPackageAssignmentApprovalSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}

	in := input[0].(map[string]interface{})

	result := stable.AccessPackageAssignmentApprovalSettings{
		IsApprovalRequiredForAdd:         nullable.Value(in["approval_required"].(bool)),
		IsApprovalRequiredForUpdate:      nullable.Value(in["approval_required_for_extension"].(bool)),
		IsRequestorJustificationRequired: nullable.Value(in["requestor_justification_required"].(bool)),
	}

	stages := make([]stable.AccessPackageApprovalStage, 0)
	if _, ok := in["approval_stage"]; ok {
		for _, raw := range in["approval_stage"].([]interface{}) {
			v := raw.(map[string]interface{})

			timeoutDays := v["approval_timeout_in_days"].(int)
			escalationDays := v["enable_alternative_approval_in_days"].(int)

			stage := stable.AccessPackageApprovalStage{
				IsApproverJustificationRequired: nullable.NoZero(v["approver_justification_required"].(bool)),
				IsEscalationEnabled:             nullable.NoZero(v["alternative_approval_enabled"].(bool)),
			}

			if timeoutDays > 0 {
				stage.DurationBeforeAutomaticDenial = nullable.Value(fmt.Sprintf("P%dD", timeoutDays))
			}
			if escalationDays > 0 {
				stage.DurationBeforeEscalation = nullable.Value(fmt.Sprintf("P%dD", escalationDays))
			}

			if _, ok := v["primary_approver"]; ok {
				primaryApprovers, err := expandSubjectSets(v["primary_approver"].([]interface{}))
				if err != nil {
					return nil, fmt.Errorf("building `primary_approver`: %v", err)
				}
				stage.PrimaryApprovers = primaryApprovers
			}

			if _, ok := v["alternative_approver"]; ok {
				escalationApprovers, err := expandSubjectSets(v["alternative_approver"].([]interface{}))
				if err != nil {
					return nil, fmt.Errorf("building `alternative_approver`: %v", err)
				}
				stage.EscalationApprovers = escalationApprovers
			}

			stages = append(stages, stage)
		}
	}

	result.Stages = &stages

	return &result, nil
}

func flattenApprovalSettings(input *stable.AccessPackageAssignmentApprovalSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	result := []map[string]interface{}{{
		"approval_required":                input.IsApprovalRequiredForAdd.GetOrZero(),
		"approval_required_for_extension":  input.IsApprovalRequiredForUpdate.GetOrZero(),
		"requestor_justification_required": input.IsRequestorJustificationRequired.GetOrZero(),
	}}

	approvalStages := make([]interface{}, 0)
	if input.Stages != nil {
		for _, v := range *input.Stages {
			timeoutDays := parseISO8601Days(v.DurationBeforeAutomaticDenial.GetOrZero())
			escalationDays := parseISO8601Days(v.DurationBeforeEscalation.GetOrZero())

			approvalStage := map[string]interface{}{
				"approval_timeout_in_days":            timeoutDays,
				"approver_justification_required":     v.IsApproverJustificationRequired.GetOrZero(),
				"alternative_approval_enabled":        v.IsEscalationEnabled.GetOrZero(),
				"enable_alternative_approval_in_days": escalationDays,
				"primary_approver":                    flattenSubjectSets(v.PrimaryApprovers),
				"alternative_approver":                flattenSubjectSets(v.EscalationApprovers),
			}

			approvalStages = append(approvalStages, approvalStage)
		}
	}

	result[0]["approval_stage"] = approvalStages

	return result
}

func expandAssignmentReviewSettings(input []interface{}) (*stable.AccessPackageAssignmentReviewSettings, error) {
	if len(input) == 0 {
		return nil, nil
	}

	in := input[0].(map[string]interface{})

	isEnabled := in["enabled"].(bool)
	durationInDays := in["duration_in_days"].(int)
	reviewFrequency := in["review_frequency"].(string)
	reviewType := in["review_type"].(string)
	startingOn := in["starting_on"].(string)
	timeoutBehavior := in["access_review_timeout_behavior"].(string)
	accessRecommendation := in["access_recommendation_enabled"].(bool)
	approverJustification := in["approver_justification_required"].(bool)

	if !isEnabled && durationInDays == 0 && reviewFrequency == "" && reviewType == "" {
		return nil, nil
	}

	result := stable.AccessPackageAssignmentReviewSettings{
		IsEnabled:                       nullable.Value(isEnabled),
		IsRecommendationEnabled:         nullable.Value(accessRecommendation),
		IsReviewerJustificationRequired: nullable.Value(approverJustification),
	}

	if timeoutBehavior != "" {
		result.ExpirationBehavior = pointer.To(stable.AccessReviewExpirationBehavior(timeoutBehavior))
	}

	isSelf := reviewType == AccessReviewReviewerTypeSelf
	result.IsSelfReview = nullable.Value(isSelf)

	// Build the schedule
	if reviewFrequency != "" || durationInDays > 0 || startingOn != "" {
		patternType, interval := reviewFrequencyToRecurrencePattern(reviewFrequency)
		rangeType := stable.RecurrenceRangeType_NoEnd

		schedule := stable.EntitlementManagementSchedule{
			StartDateTime: nullable.NoZero(startingOn),
			Recurrence: &stable.PatternedRecurrence{
				Pattern: &stable.RecurrencePattern{
					Type:     patternType,
					Interval: interval,
				},
				Range: &stable.RecurrenceRange{
					Type: rangeType,
				},
			},
		}
		if durationInDays > 0 {
			schedule.Expiration = &stable.ExpirationPattern{
				Duration: nullable.Value(fmt.Sprintf("P%dD", durationInDays)),
				Type:     pointer.To(stable.ExpirationPatternType_AfterDuration),
			}
		}
		result.Schedule = &schedule
	}

	if reviewType != AccessReviewReviewerTypeSelf {
		if _, ok := in["reviewer"]; ok {
			reviewers, err := expandSubjectSets(in["reviewer"].([]interface{}))
			if err != nil {
				return nil, fmt.Errorf("building `reviewer`: %v", err)
			}
			result.PrimaryReviewers = reviewers
		}
	}

	return &result, nil
}

func flattenAssignmentReviewSettings(input *stable.AccessPackageAssignmentReviewSettings) []map[string]interface{} {
	if input == nil {
		return nil
	}

	reviewType := ""
	if input.IsSelfReview.GetOrZero() {
		reviewType = AccessReviewReviewerTypeSelf
	} else if input.PrimaryReviewers != nil && len(*input.PrimaryReviewers) > 0 {
		// Check if the primary reviewer is requestorManager
		for _, r := range *input.PrimaryReviewers {
			if _, ok := r.(stable.RequestorManager); ok {
				reviewType = AccessReviewReviewerTypeManager
				break
			}
		}
		if reviewType == "" {
			reviewType = AccessReviewReviewerTypeReviewers
		}
	}

	timeoutBehavior := ""
	if input.ExpirationBehavior != nil {
		timeoutBehavior = string(*input.ExpirationBehavior)
	}

	durationInDays := 0
	reviewFrequency := ""
	startingOn := ""

	if input.Schedule != nil {
		if input.Schedule.Expiration != nil {
			durationInDays = parseISO8601Days(input.Schedule.Expiration.Duration.GetOrZero())
		}
		if input.Schedule.StartDateTime.GetOrZero() != "" {
			startingOn = input.Schedule.StartDateTime.GetOrZero()
		}
		if input.Schedule.Recurrence != nil && input.Schedule.Recurrence.Pattern != nil {
			reviewFrequency = recurrencePatternToReviewFrequency(
				input.Schedule.Recurrence.Pattern.Type,
				input.Schedule.Recurrence.Pattern.Interval,
			)
		}
	}

	return []map[string]interface{}{{
		"access_recommendation_enabled":   input.IsRecommendationEnabled.GetOrZero(),
		"access_review_timeout_behavior":  timeoutBehavior,
		"approver_justification_required": input.IsReviewerJustificationRequired.GetOrZero(),
		"duration_in_days":                durationInDays,
		"enabled":                         input.IsEnabled.GetOrZero(),
		"review_frequency":                reviewFrequency,
		"review_type":                     reviewType,
		"reviewer":                        flattenSubjectSets(input.PrimaryReviewers),
		"starting_on":                     startingOn,
	}}
}

// expandSubjectSets converts the Terraform schema user set blocks to stable SubjectSet slices.
// The backup field has no equivalent in the stable API and is silently ignored.
func expandSubjectSets(input []interface{}) (*[]stable.SubjectSet, error) {
	subjectSets := make([]stable.SubjectSet, 0)
	for _, raw := range input {
		v := raw.(map[string]interface{})

		objectId := v["object_id"].(string)
		odataType := formatODataType(v["subject_type"].(string))

		var subjectSet stable.SubjectSet
		switch odataType {
		case "ConnectedOrganizationMembers":
			subjectSet = stable.ConnectedOrganizationMembers{
				ConnectedOrganizationId: nullable.NoZero(objectId),
			}
		case "ExternalSponsors":
			subjectSet = stable.ExternalSponsors{}
		case "GroupMembers":
			subjectSet = stable.GroupMembers{
				GroupId: nullable.NoZero(objectId),
			}
		case "InternalSponsors":
			subjectSet = stable.InternalSponsors{}
		case "RequestorManager":
			subjectSet = stable.RequestorManager{}
		case "SingleUser":
			subjectSet = stable.SingleUser{
				UserId: nullable.NoZero(objectId),
			}
		case "TargetUserSponsors":
			subjectSet = stable.TargetUserSponsors{}
		default:
			return nil, fmt.Errorf("unknown `subject_type`: %s", odataType)
		}

		subjectSets = append(subjectSets, subjectSet)
	}

	return &subjectSets, nil
}

// flattenSubjectSets converts stable SubjectSet slices to Terraform schema user set blocks.
// The backup field is always set to false as the stable API does not support it.
func flattenSubjectSets(input *[]stable.SubjectSet) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	results := make([]map[string]interface{}, 0)
	for _, raw := range *input {
		base := raw.SubjectSet()
		var objectId string

		switch impl := raw.(type) {
		case stable.ConnectedOrganizationMembers:
			objectId = impl.ConnectedOrganizationId.GetOrZero()
		case stable.GroupMembers:
			objectId = impl.GroupId.GetOrZero()
		case stable.SingleUser:
			objectId = impl.UserId.GetOrZero()
		}

		results = append(results, map[string]interface{}{
			"subject_type": formatODataType(pointer.From(base.ODataType)),
			"backup":       false,
			"object_id":    objectId,
		})
	}

	return results
}

func expandAccessPackageQuestions(input []interface{}) *[]stable.AccessPackageQuestion {
	result := make([]stable.AccessPackageQuestion, 0)

	for _, raw := range input {
		v := raw.(map[string]interface{})

		textList := v["text"].([]interface{})
		if len(textList) == 0 {
			continue
		}
		textBlock := textList[0].(map[string]interface{})
		defaultText := textBlock["default_text"].(string)
		localizations := expandAccessPackageLocalizedTexts(textBlock)

		var question stable.AccessPackageQuestion

		if choicesRaw, ok := v["choice"].([]interface{}); ok && len(choicesRaw) > 0 {
			choices := make([]stable.AccessPackageAnswerChoice, 0)
			for _, choiceRaw := range choicesRaw {
				choice := choiceRaw.(map[string]interface{})
				choiceText := ""
				var choiceLocalizations *[]stable.AccessPackageLocalizedText
				if dv := choice["display_value"].([]interface{}); len(dv) > 0 {
					dvBlock := dv[0].(map[string]interface{})
					choiceText = dvBlock["default_text"].(string)
					choiceLocalizations = expandAccessPackageLocalizedTexts(dvBlock)
				}

				choices = append(choices, stable.AccessPackageAnswerChoice{
					ActualValue:   nullable.NoZero(choice["actual_value"].(string)),
					Text:          nullable.NoZero(choiceText),
					Localizations: choiceLocalizations,
				})
			}

			question = stable.AccessPackageMultipleChoiceQuestion{
				Choices:       &choices,
				IsRequired:    nullable.Value(v["required"].(bool)),
				Sequence:      nullable.Value(int64(v["sequence"].(int))),
				Text:          nullable.NoZero(defaultText),
				Localizations: localizations,
			}
		} else {
			question = stable.AccessPackageTextInputQuestion{
				IsRequired:    nullable.Value(v["required"].(bool)),
				Sequence:      nullable.Value(int64(v["sequence"].(int))),
				Text:          nullable.NoZero(defaultText),
				Localizations: localizations,
			}
		}

		result = append(result, question)
	}

	return &result
}

func flattenAccessPackageQuestions(input *[]stable.AccessPackageQuestion) []map[string]interface{} {
	if input == nil || len(*input) == 0 {
		return nil
	}

	questions := make([]map[string]interface{}, 0)

	for _, raw := range *input {
		v := raw.AccessPackageQuestion()

		question := map[string]interface{}{
			"required": v.IsRequired.GetOrZero(),
			"sequence": int(v.Sequence.GetOrZero()),
			"text":     flattenAccessPackageLocalizedTexts(v.Text.GetOrZero(), v.Localizations),
		}

		if impl, ok := raw.(stable.AccessPackageMultipleChoiceQuestion); ok {
			choices := make([]map[string]interface{}, 0)
			for _, choice := range pointer.From(impl.Choices) {
				choices = append(choices, map[string]interface{}{
					"actual_value":  choice.ActualValue.GetOrZero(),
					"display_value": flattenAccessPackageLocalizedTexts(choice.Text.GetOrZero(), choice.Localizations),
				})
			}
			question["choice"] = choices
		}

		questions = append(questions, question)
	}

	return questions
}

// expandAccessPackageLocalizedTexts converts the localized_text schema block to a
// stable AccessPackageLocalizedText slice.
func expandAccessPackageLocalizedTexts(input map[string]interface{}) *[]stable.AccessPackageLocalizedText {
	texts := make([]stable.AccessPackageLocalizedText, 0)

	if raw, ok := input["localized_text"]; ok {
		for _, item := range raw.([]interface{}) {
			v := item.(map[string]interface{})
			texts = append(texts, stable.AccessPackageLocalizedText{
				LanguageCode: v["language_code"].(string),
				Text:         nullable.NoZero(v["content"].(string)),
			})
		}
	}

	if len(texts) == 0 {
		return nil
	}
	return &texts
}

// flattenAccessPackageLocalizedTexts converts stable localized text fields back to the
// schema's localized_content block (default_text + localized_text list).
func flattenAccessPackageLocalizedTexts(defaultText string, localizations *[]stable.AccessPackageLocalizedText) []map[string]interface{} {
	texts := make([]map[string]interface{}, 0)
	if localizations != nil {
		for _, v := range *localizations {
			texts = append(texts, map[string]interface{}{
				"language_code": v.LanguageCode,
				"content":       v.Text.GetOrZero(),
			})
		}
	}
	return []map[string]interface{}{{
		"default_text":   defaultText,
		"localized_text": texts,
	}}
}

// parseISO8601Days parses a simple ISO 8601 duration of the form "P{n}D" and returns n.
// Returns 0 for empty or unparseable strings.
func parseISO8601Days(duration string) int {
	if duration == "" {
		return 0
	}
	var days int
	fmt.Sscanf(duration, "P%dD", &days)
	return days
}

func formatODataType(in string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.TrimPrefix(in, "#microsoft.graph."))
}
