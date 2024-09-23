package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupSetting{}

type GroupSetting struct {
	// Display name of this group of settings, which comes from the associated template.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level
	// settings object. Read-only.
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

	// Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced
	// groupSettingTemplates object.
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

func (s GroupSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupSetting{}

func (s GroupSetting) MarshalJSON() ([]byte, error) {
	type wrapper GroupSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupSetting: %+v", err)
	}

	delete(decoded, "templateId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupSetting: %+v", err)
	}

	return encoded, nil
}
