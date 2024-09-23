package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdatableAsset interface {
	Entity
	WindowsUpdatesUpdatableAsset() BaseWindowsUpdatesUpdatableAssetImpl
}

var _ WindowsUpdatesUpdatableAsset = BaseWindowsUpdatesUpdatableAssetImpl{}

type BaseWindowsUpdatesUpdatableAssetImpl struct {

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

func (s BaseWindowsUpdatesUpdatableAssetImpl) WindowsUpdatesUpdatableAsset() BaseWindowsUpdatesUpdatableAssetImpl {
	return s
}

func (s BaseWindowsUpdatesUpdatableAssetImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesUpdatableAsset = RawWindowsUpdatesUpdatableAssetImpl{}

// RawWindowsUpdatesUpdatableAssetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesUpdatableAssetImpl struct {
	windowsUpdatesUpdatableAsset BaseWindowsUpdatesUpdatableAssetImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawWindowsUpdatesUpdatableAssetImpl) WindowsUpdatesUpdatableAsset() BaseWindowsUpdatesUpdatableAssetImpl {
	return s.windowsUpdatesUpdatableAsset
}

func (s RawWindowsUpdatesUpdatableAssetImpl) Entity() BaseEntityImpl {
	return s.windowsUpdatesUpdatableAsset.Entity()
}

var _ json.Marshaler = BaseWindowsUpdatesUpdatableAssetImpl{}

func (s BaseWindowsUpdatesUpdatableAssetImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesUpdatableAssetImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesUpdatableAssetImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesUpdatableAssetImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.updatableAsset"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesUpdatableAssetImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesUpdatableAssetImplementation(input []byte) (WindowsUpdatesUpdatableAsset, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdatableAsset into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.azureADDevice") {
		var out WindowsUpdatesAzureADDevice
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesAzureADDevice: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.updatableAssetGroup") {
		var out WindowsUpdatesUpdatableAssetGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesUpdatableAssetGroup: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesUpdatableAssetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesUpdatableAssetImpl: %+v", err)
	}

	return RawWindowsUpdatesUpdatableAssetImpl{
		windowsUpdatesUpdatableAsset: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
