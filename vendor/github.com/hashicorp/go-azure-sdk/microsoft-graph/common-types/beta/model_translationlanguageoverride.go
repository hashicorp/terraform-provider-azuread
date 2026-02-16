package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationLanguageOverride struct {
	// The language to apply the override.Returned by default. Not nullable.
	LanguageTag *string `json:"languageTag,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The translation override behavior for the language, if any.Returned by default. Not nullable.
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`
}
