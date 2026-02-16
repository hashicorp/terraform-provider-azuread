package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DirectorySetting{}

type DirectorySetting struct {
	// Display name of this group of settings, which comes from the associated template. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique identifier for the template used to create this group of settings. Read-only.
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

	// Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced
	// directorySettingTemplates object.
	Values *[]SettingValue `json:"values,omitempty"`

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

func (s DirectorySetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectorySetting{}

func (s DirectorySetting) MarshalJSON() ([]byte, error) {
	type wrapper DirectorySetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectorySetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectorySetting: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "templateId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directorySetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectorySetting: %+v", err)
	}

	return encoded, nil
}
