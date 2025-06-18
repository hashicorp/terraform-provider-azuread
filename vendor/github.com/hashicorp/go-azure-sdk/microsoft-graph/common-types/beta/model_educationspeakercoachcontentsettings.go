package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSpeakerCoachContentSettings struct {
	// Indicates whether the student should receive feedback on their inclusiveness from the Speaker Coach.
	IsInclusivenessEnabled *bool `json:"isInclusivenessEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their repetitive language from the Speaker Coach.
	IsRepetitiveLanguageEnabled *bool `json:"isRepetitiveLanguageEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
