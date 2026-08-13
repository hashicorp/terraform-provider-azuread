package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
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

	Password nullable.Type[string] `json:"password,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsDeviceAccountImpl) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return s
}

var _ WindowsDeviceAccount = RawWindowsDeviceAccountImpl{}

// RawWindowsDeviceAccountImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawWindowsDeviceAccountImpl struct {
	windowsDeviceAccount BaseWindowsDeviceAccountImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawWindowsDeviceAccountImpl) WindowsDeviceAccount() BaseWindowsDeviceAccountImpl {
	return s.windowsDeviceAccount
}

func (s RawWindowsDeviceAccountImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
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
