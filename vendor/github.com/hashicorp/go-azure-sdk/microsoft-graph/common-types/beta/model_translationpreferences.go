package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslationPreferences struct {
	// Translation override behavior for languages, if any.Returned by default.
	LanguageOverrides *[]TranslationLanguageOverride `json:"languageOverrides,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The user's preferred translation behavior.Returned by default. Not nullable.
	TranslationBehavior *TranslationBehavior `json:"translationBehavior,omitempty"`

	// The list of languages the user does not need translated. This is computed from the authoringLanguages collection in
	// regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies
	// neutral culture values that include the language code without any country or region association. For example, it
	// would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by
	// default. Read only.
	UntranslatedLanguages *[]string `json:"untranslatedLanguages,omitempty"`
}
