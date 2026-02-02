package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingInstance interface {
	Entity
	DeviceManagementSettingInstance() BaseDeviceManagementSettingInstanceImpl
}

var _ DeviceManagementSettingInstance = BaseDeviceManagementSettingInstanceImpl{}

type BaseDeviceManagementSettingInstanceImpl struct {
	// The ID of the setting definition for this instance
	DefinitionId *string `json:"definitionId,omitempty"`

	// JSON representation of the value
	ValueJson nullable.Type[string] `json:"valueJson,omitempty"`

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

func (s BaseDeviceManagementSettingInstanceImpl) DeviceManagementSettingInstance() BaseDeviceManagementSettingInstanceImpl {
	return s
}

func (s BaseDeviceManagementSettingInstanceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementSettingInstance = RawDeviceManagementSettingInstanceImpl{}

// RawDeviceManagementSettingInstanceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementSettingInstanceImpl struct {
	deviceManagementSettingInstance BaseDeviceManagementSettingInstanceImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawDeviceManagementSettingInstanceImpl) DeviceManagementSettingInstance() BaseDeviceManagementSettingInstanceImpl {
	return s.deviceManagementSettingInstance
}

func (s RawDeviceManagementSettingInstanceImpl) Entity() BaseEntityImpl {
	return s.deviceManagementSettingInstance.Entity()
}

var _ json.Marshaler = BaseDeviceManagementSettingInstanceImpl{}

func (s BaseDeviceManagementSettingInstanceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementSettingInstanceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementSettingInstanceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementSettingInstanceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementSettingInstanceImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceManagementSettingInstanceImplementation(input []byte) (DeviceManagementSettingInstance, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingInstance into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementAbstractComplexSettingInstance") {
		var out DeviceManagementAbstractComplexSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAbstractComplexSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementBooleanSettingInstance") {
		var out DeviceManagementBooleanSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementBooleanSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCollectionSettingInstance") {
		var out DeviceManagementCollectionSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCollectionSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplexSettingInstance") {
		var out DeviceManagementComplexSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplexSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntegerSettingInstance") {
		var out DeviceManagementIntegerSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntegerSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementStringSettingInstance") {
		var out DeviceManagementStringSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementStringSettingInstance: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementSettingInstanceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementSettingInstanceImpl: %+v", err)
	}

	return RawDeviceManagementSettingInstanceImpl{
		deviceManagementSettingInstance: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
