package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAiFeedbackContentSettings struct {
	// Indicates whether the student should receive feedback on their message clarity from the AI feedback.
	IsMessageClarityEnabled *bool `json:"isMessageClarityEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their quality of information from the AI feedback.
	IsQualityOfInformationEnabled *bool `json:"isQualityOfInformationEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their speech organization from the AI feedback.
	IsSpeechOrganizationEnabled *bool `json:"isSpeechOrganizationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
