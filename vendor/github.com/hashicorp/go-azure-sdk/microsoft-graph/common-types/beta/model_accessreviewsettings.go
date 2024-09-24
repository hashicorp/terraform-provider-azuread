package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewSettings interface {
	AccessReviewSettings() BaseAccessReviewSettingsImpl
}

var _ AccessReviewSettings = BaseAccessReviewSettingsImpl{}

type BaseAccessReviewSettingsImpl struct {
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

func (s BaseAccessReviewSettingsImpl) AccessReviewSettings() BaseAccessReviewSettingsImpl {
	return s
}

var _ AccessReviewSettings = RawAccessReviewSettingsImpl{}

// RawAccessReviewSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessReviewSettingsImpl struct {
	accessReviewSettings BaseAccessReviewSettingsImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawAccessReviewSettingsImpl) AccessReviewSettings() BaseAccessReviewSettingsImpl {
	return s.accessReviewSettings
}

func UnmarshalAccessReviewSettingsImplementation(input []byte) (AccessReviewSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.businessFlowSettings") {
		var out BusinessFlowSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessFlowSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewSettingsImpl: %+v", err)
	}

	return RawAccessReviewSettingsImpl{
		accessReviewSettings: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
