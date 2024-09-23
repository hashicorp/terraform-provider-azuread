package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = DirectorySettingTemplate{}

type DirectorySettingTemplate struct {
	// Description of the template. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the template. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Collection of settingTemplateValues that list the set of available settings, defaults, and types that make up this
	// template. Read-only.
	Values *[]SettingTemplateValue `json:"values,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s DirectorySettingTemplate) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s DirectorySettingTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectorySettingTemplate{}

func (s DirectorySettingTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DirectorySettingTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectorySettingTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectorySettingTemplate: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "values")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directorySettingTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectorySettingTemplate: %+v", err)
	}

	return encoded, nil
}
