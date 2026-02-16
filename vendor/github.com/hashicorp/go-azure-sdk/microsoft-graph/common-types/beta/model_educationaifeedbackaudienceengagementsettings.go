package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAiFeedbackAudienceEngagementSettings struct {
	// Indicates whether the student should receive feedback on their engagement strategies from the AI feedback.
	AreEngagementStrategiesEnabled *bool `json:"areEngagementStrategiesEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their call to action from the AI feedback.
	IsCallToActionEnabled *bool `json:"isCallToActionEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their emotional and intellectual appeal from the AI
	// feedback.
	IsEmotionalAndIntellectualAppealEnabled *bool `json:"isEmotionalAndIntellectualAppealEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
