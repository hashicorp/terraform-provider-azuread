package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewSettings = BusinessFlowSettings{}

type BusinessFlowSettings struct {
	DurationInDays *int64 `json:"durationInDays,omitempty"`

	// Fields inherited from AccessReviewSettings

	// Indicates whether showing recommendations to reviewers is enabled.
	AccessRecommendationsEnabled *bool `json:"accessRecommendationsEnabled,omitempty"`

	// The number of days of user activities to show to reviewers.
	ActivityDurationInDays *int64 `json:"activityDurationInDays,omitempty"`

	// Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.
	// If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyReviewResultsEnabled *bool `json:"autoApplyReviewResultsEnabled,omitempty"`

	// Indicates whether a decision should be set if the reviewer didn't supply one. For use when, auto-apply is enabled. If
	// you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
	AutoReviewEnabled *bool `json:"autoReviewEnabled,omitempty"`

	// Detailed settings for how the feature should set the review decision. For use when, auto-apply is enabled.
	AutoReviewSettings *AutoReviewSettings `json:"autoReviewSettings,omitempty"`

	// Indicates whether reviewers are required to provide a justification when reviewing access.
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`

	// Indicates whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Detailed settings for recurrence.
	RecurrenceSettings *AccessReviewRecurrenceSettings `json:"recurrenceSettings,omitempty"`

	// Indicates whether sending reminder emails to reviewers is enabled.
	RemindersEnabled *bool `json:"remindersEnabled,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BusinessFlowSettings) AccessReviewSettings() BaseAccessReviewSettingsImpl {
	return BaseAccessReviewSettingsImpl{
		AccessRecommendationsEnabled:    s.AccessRecommendationsEnabled,
		ActivityDurationInDays:          s.ActivityDurationInDays,
		AutoApplyReviewResultsEnabled:   s.AutoApplyReviewResultsEnabled,
		AutoReviewEnabled:               s.AutoReviewEnabled,
		AutoReviewSettings:              s.AutoReviewSettings,
		JustificationRequiredOnApproval: s.JustificationRequiredOnApproval,
		MailNotificationsEnabled:        s.MailNotificationsEnabled,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
		RecurrenceSettings:              s.RecurrenceSettings,
		RemindersEnabled:                s.RemindersEnabled,
	}
}

var _ json.Marshaler = BusinessFlowSettings{}

func (s BusinessFlowSettings) MarshalJSON() ([]byte, error) {
	type wrapper BusinessFlowSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BusinessFlowSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessFlowSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.businessFlowSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BusinessFlowSettings: %+v", err)
	}

	return encoded, nil
}
