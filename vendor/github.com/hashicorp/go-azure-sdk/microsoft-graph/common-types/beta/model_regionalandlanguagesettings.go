package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RegionalAndLanguageSettings{}

type RegionalAndLanguageSettings struct {
	// Prioritized list of languages the user reads and authors in.Returned by default. Not nullable.
	AuthoringLanguages *[]LocaleInfo `json:"authoringLanguages,omitempty"`

	// The user's preferred user interface language (menus, buttons, ribbons, warning messages) for Microsoft web
	// applications.Returned by default. Not nullable.
	DefaultDisplayLanguage *LocaleInfo `json:"defaultDisplayLanguage,omitempty"`

	// The locale that drives the default date, time, and calendar formatting.Returned by default.
	DefaultRegionalFormat *LocaleInfo `json:"defaultRegionalFormat,omitempty"`

	// The language a user expected to use as input for text to speech scenarios.Returned by default.
	DefaultSpeechInputLanguage *LocaleInfo `json:"defaultSpeechInputLanguage,omitempty"`

	// The language a user expects to have documents, emails, and messages translated into.Returned by default.
	DefaultTranslationLanguage *LocaleInfo `json:"defaultTranslationLanguage,omitempty"`

	// Allows a user to override their defaultRegionalFormat with field specific formats.Returned by default.
	RegionalFormatOverrides *RegionalFormatOverrides `json:"regionalFormatOverrides,omitempty"`

	// The user's preferred settings when consuming translated documents, emails, messages, and websites.Returned by
	// default. Not nullable.
	TranslationPreferences *TranslationPreferences `json:"translationPreferences,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RegionalAndLanguageSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RegionalAndLanguageSettings{}

func (s RegionalAndLanguageSettings) MarshalJSON() ([]byte, error) {
	type wrapper RegionalAndLanguageSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegionalAndLanguageSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegionalAndLanguageSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.regionalAndLanguageSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegionalAndLanguageSettings: %+v", err)
	}

	return encoded, nil
}
