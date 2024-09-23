package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewScheduleSettings struct {
	// Optional field. Describes the actions to take once a review is complete. There are two types that are currently
	// supported: removeAccessApplyAction (default) and disableAndDeleteUserApplyAction. Field only needs to be specified in
	// the case of disableAndDeleteUserApplyAction.
	ApplyActions *[]AccessReviewApplyAction `json:"applyActions,omitempty"`

	// Indicates whether decisions are automatically applied. When set to false, an admin must apply the decisions manually
	// once the reviewer completes the access review. When set to true, decisions are applied automatically after the access
	// review instance duration ends, whether or not the reviewers have responded. Default value is false. CAUTION: If both
	// autoApplyDecisionsEnabled and defaultDecisionEnabled are true, all access for the principals to the resource risks
	// being revoked if the reviewers fail to respond.
	AutoApplyDecisionsEnabled *bool `json:"autoApplyDecisionsEnabled,omitempty"`

	// Indicates whether decisions on previous access review stages are available for reviewers on an accessReviewInstance
	// with multiple subsequent stages. If not provided, the default is disabled (false).
	DecisionHistoriesForReviewersEnabled nullable.Type[bool] `json:"decisionHistoriesForReviewersEnabled,omitempty"`

	// Decision chosen if defaultDecisionEnabled is enabled. Can be one of Approve, Deny, or Recommendation.
	DefaultDecision nullable.Type[string] `json:"defaultDecision,omitempty"`

	// Indicates whether the default decision is enabled or disabled when reviewers do not respond. Default value is false.
	// CAUTION: If both autoApplyDecisionsEnabled and defaultDecisionEnabled are true, all access for the principals to the
	// resource risks being revoked if the reviewers fail to respond.
	DefaultDecisionEnabled *bool `json:"defaultDecisionEnabled,omitempty"`

	// Duration of an access review instance in days. NOTE: If the stageSettings of the accessReviewScheduleDefinition
	// object is defined, its durationInDays setting will be used instead of the value of this property.
	InstanceDurationInDays *int64 `json:"instanceDurationInDays,omitempty"`

	// Indicates whether reviewers are required to provide justification with their decision. Default value is false.
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`

	// Indicates whether emails are enabled or disabled. Default value is false.
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. Describes the types of insights that aid reviewers to make access review decisions. NOTE: If the
	// stageSettings of the accessReviewScheduleDefinition object is defined, its recommendationInsightSettings setting will
	// be used instead of the value of this property.
	RecommendationInsightSettings *[]AccessReviewRecommendationInsightSetting `json:"recommendationInsightSettings,omitempty"`

	// Optional field. Indicates the period of inactivity (with respect to the start date of the review instance) that
	// recommendations will be configured from. The recommendation will be to deny if the user is inactive during the
	// look-back duration. For reviews of groups and Microsoft Entra roles, any duration is accepted. For reviews of
	// applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: If the stageSettings
	// of the accessReviewScheduleDefinition object is defined, its recommendationLookBackDuration setting will be used
	// instead of the value of this property.
	RecommendationLookBackDuration nullable.Type[string] `json:"recommendationLookBackDuration,omitempty"`

	// Indicates whether decision recommendations are enabled or disabled. NOTE: If the stageSettings of the
	// accessReviewScheduleDefinition object is defined, its recommendationsEnabled setting will be used instead of the
	// value of this property.
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`

	// Detailed settings for recurrence using the standard Outlook recurrence object. Note: Only dayOfMonth, interval, and
	// type (weekly, absoluteMonthly) properties are supported. Use the property startDate on recurrenceRange to determine
	// the day the review starts.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// Indicates whether reminders are enabled or disabled. Default value is false.
	ReminderNotificationsEnabled *bool `json:"reminderNotificationsEnabled,omitempty"`
}

var _ json.Unmarshaler = &AccessReviewScheduleSettings{}

func (s *AccessReviewScheduleSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AutoApplyDecisionsEnabled            *bool                 `json:"autoApplyDecisionsEnabled,omitempty"`
		DecisionHistoriesForReviewersEnabled nullable.Type[bool]   `json:"decisionHistoriesForReviewersEnabled,omitempty"`
		DefaultDecision                      nullable.Type[string] `json:"defaultDecision,omitempty"`
		DefaultDecisionEnabled               *bool                 `json:"defaultDecisionEnabled,omitempty"`
		InstanceDurationInDays               *int64                `json:"instanceDurationInDays,omitempty"`
		JustificationRequiredOnApproval      *bool                 `json:"justificationRequiredOnApproval,omitempty"`
		MailNotificationsEnabled             *bool                 `json:"mailNotificationsEnabled,omitempty"`
		ODataId                              *string               `json:"@odata.id,omitempty"`
		ODataType                            *string               `json:"@odata.type,omitempty"`
		RecommendationLookBackDuration       nullable.Type[string] `json:"recommendationLookBackDuration,omitempty"`
		RecommendationsEnabled               *bool                 `json:"recommendationsEnabled,omitempty"`
		Recurrence                           *PatternedRecurrence  `json:"recurrence,omitempty"`
		ReminderNotificationsEnabled         *bool                 `json:"reminderNotificationsEnabled,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AutoApplyDecisionsEnabled = decoded.AutoApplyDecisionsEnabled
	s.DecisionHistoriesForReviewersEnabled = decoded.DecisionHistoriesForReviewersEnabled
	s.DefaultDecision = decoded.DefaultDecision
	s.DefaultDecisionEnabled = decoded.DefaultDecisionEnabled
	s.InstanceDurationInDays = decoded.InstanceDurationInDays
	s.JustificationRequiredOnApproval = decoded.JustificationRequiredOnApproval
	s.MailNotificationsEnabled = decoded.MailNotificationsEnabled
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecommendationLookBackDuration = decoded.RecommendationLookBackDuration
	s.RecommendationsEnabled = decoded.RecommendationsEnabled
	s.Recurrence = decoded.Recurrence
	s.ReminderNotificationsEnabled = decoded.ReminderNotificationsEnabled

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewScheduleSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["applyActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ApplyActions into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewApplyAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewApplyActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ApplyActions' for 'AccessReviewScheduleSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ApplyActions = &output
	}

	if v, ok := temp["recommendationInsightSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RecommendationInsightSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessReviewRecommendationInsightSetting, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessReviewRecommendationInsightSettingImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RecommendationInsightSettings' for 'AccessReviewScheduleSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RecommendationInsightSettings = &output
	}

	return nil
}
