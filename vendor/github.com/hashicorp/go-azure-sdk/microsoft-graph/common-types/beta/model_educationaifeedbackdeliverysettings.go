package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAiFeedbackDeliverySettings struct {
	// Indicates whether the student should receive feedback on their rhetorical techniques from the AI feedback.
	AreRhetoricalTechniquesEnabled *bool `json:"areRhetoricalTechniquesEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their language use from the AI feedback.
	IsLanguageUseEnabled *bool `json:"isLanguageUseEnabled,omitempty"`

	// Indicates whether the student should receive feedback on their style from the AI feedback.
	IsStyleEnabled *bool `json:"isStyleEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
