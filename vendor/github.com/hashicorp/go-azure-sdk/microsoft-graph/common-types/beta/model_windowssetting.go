package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsSetting{}

type WindowsSetting struct {
	// A collection of setting values for a given windowsSetting.
	Instances *[]WindowsSettingInstance `json:"instances,omitempty"`

	// The type of setting payloads contained in the instances navigation property.
	PayloadType nullable.Type[string] `json:"payloadType,omitempty"`

	SettingType *WindowsSettingType `json:"settingType,omitempty"`

	// A unique identifier for the device the setting might belong to if it is of the settingType backup.
	WindowsDeviceId nullable.Type[string] `json:"windowsDeviceId,omitempty"`

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

func (s WindowsSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsSetting{}

func (s WindowsSetting) MarshalJSON() ([]byte, error) {
	type wrapper WindowsSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsSetting: %+v", err)
	}

	return encoded, nil
}
