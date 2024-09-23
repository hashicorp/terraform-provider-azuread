package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceAccount interface {
	WindowsDeviceAccount() BaseWindowsDeviceAccountImpl
}

var _ WindowsDeviceAccount = BaseWindowsDeviceAccountImpl{}

type BaseWindowsDeviceAccountImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Not yet documented
	Password nullable.Type[string] `json:"password,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsDeviceAccountImpl) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return s
}

var _ WindowsDeviceAccount = RawWindowsDeviceAccountImpl{}

// RawWindowsDeviceAccountImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsDeviceAccountImpl struct {
	windowsDeviceAccount BaseWindowsDeviceAccountImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawWindowsDeviceAccountImpl) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return s.windowsDeviceAccount
}

func UnmarshalWindowsDeviceAccountImplementation(input []byte) (WindowsDeviceAccount, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsDeviceAccount into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDeviceADAccount") {
		var out WindowsDeviceADAccount
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDeviceADAccount: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDeviceAzureADAccount") {
		var out WindowsDeviceAzureADAccount
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDeviceAzureADAccount: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsDeviceAccountImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsDeviceAccountImpl: %+v", err)
	}

	return RawWindowsDeviceAccountImpl{
		windowsDeviceAccount: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
