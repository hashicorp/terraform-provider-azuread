package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeFolderItem interface {
	AndroidDeviceOwnerKioskModeHomeScreenItem
	AndroidDeviceOwnerKioskModeFolderItem() BaseAndroidDeviceOwnerKioskModeFolderItemImpl
}

var _ AndroidDeviceOwnerKioskModeFolderItem = BaseAndroidDeviceOwnerKioskModeFolderItemImpl{}

type BaseAndroidDeviceOwnerKioskModeFolderItemImpl struct {

	// Fields inherited from AndroidDeviceOwnerKioskModeHomeScreenItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAndroidDeviceOwnerKioskModeFolderItemImpl) AndroidDeviceOwnerKioskModeFolderItem() BaseAndroidDeviceOwnerKioskModeFolderItemImpl {
	return s
}

func (s BaseAndroidDeviceOwnerKioskModeFolderItemImpl) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AndroidDeviceOwnerKioskModeFolderItem = RawAndroidDeviceOwnerKioskModeFolderItemImpl{}

// RawAndroidDeviceOwnerKioskModeFolderItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAndroidDeviceOwnerKioskModeFolderItemImpl struct {
	androidDeviceOwnerKioskModeFolderItem BaseAndroidDeviceOwnerKioskModeFolderItemImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawAndroidDeviceOwnerKioskModeFolderItemImpl) AndroidDeviceOwnerKioskModeFolderItem() BaseAndroidDeviceOwnerKioskModeFolderItemImpl {
	return s.androidDeviceOwnerKioskModeFolderItem
}

func (s RawAndroidDeviceOwnerKioskModeFolderItemImpl) AndroidDeviceOwnerKioskModeHomeScreenItem() BaseAndroidDeviceOwnerKioskModeHomeScreenItemImpl {
	return s.androidDeviceOwnerKioskModeFolderItem.AndroidDeviceOwnerKioskModeHomeScreenItem()
}

var _ json.Marshaler = BaseAndroidDeviceOwnerKioskModeFolderItemImpl{}

func (s BaseAndroidDeviceOwnerKioskModeFolderItemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAndroidDeviceOwnerKioskModeFolderItemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAndroidDeviceOwnerKioskModeFolderItemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAndroidDeviceOwnerKioskModeFolderItemImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerKioskModeFolderItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAndroidDeviceOwnerKioskModeFolderItemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAndroidDeviceOwnerKioskModeFolderItemImplementation(input []byte) (AndroidDeviceOwnerKioskModeFolderItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerKioskModeFolderItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeApp") {
		var out AndroidDeviceOwnerKioskModeApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidDeviceOwnerKioskModeWeblink") {
		var out AndroidDeviceOwnerKioskModeWeblink
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidDeviceOwnerKioskModeWeblink: %+v", err)
		}
		return out, nil
	}

	var parent BaseAndroidDeviceOwnerKioskModeFolderItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAndroidDeviceOwnerKioskModeFolderItemImpl: %+v", err)
	}

	return RawAndroidDeviceOwnerKioskModeFolderItemImpl{
		androidDeviceOwnerKioskModeFolderItem: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
