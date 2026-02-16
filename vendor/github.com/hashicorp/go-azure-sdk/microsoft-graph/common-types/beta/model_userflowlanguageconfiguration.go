package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserFlowLanguageConfiguration{}

type UserFlowLanguageConfiguration struct {
	// Collection of pages with the default content to display in a user flow for a specified language. This collection
	// doesn't allow any kind of modification.
	DefaultPages *[]UserFlowLanguagePage `json:"defaultPages,omitempty"`

	// The language name to display. This property is read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the language is enabled within the user flow.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Collection of pages with the overrides messages to display in a user flow for a specified language. This collection
	// only allows to modify the content of the page, any other modification isn't allowed (creation or deletion of pages).
	OverridesPages *[]UserFlowLanguagePage `json:"overridesPages,omitempty"`

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

func (s UserFlowLanguageConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserFlowLanguageConfiguration{}

func (s UserFlowLanguageConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper UserFlowLanguageConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserFlowLanguageConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserFlowLanguageConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userFlowLanguageConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserFlowLanguageConfiguration: %+v", err)
	}

	return encoded, nil
}
