package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSpeakerCoachDeliverySettings struct {
	// Indicates whether the student should receive feedback on their filler words from the Speaker Coach.
	AreFillerWordsEnabled *bool `json:"areFillerWordsEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their pace from the Speaker Coach.
	IsPaceEnabled *bool `json:"isPaceEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their pitch from the Speaker Coach.
	IsPitchEnabled *bool `json:"isPitchEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their pronunciation from the Speaker Coach. This is
	// automatically enabled if isAiFeedbackEnabled is set to true on the educationSpeakerProgressResource, or if
	// spokenLanguageLocale is set to a value besides en-US on the educationSpeakerProgressResource.
	IsPronunciationEnabled *bool `json:"isPronunciationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
