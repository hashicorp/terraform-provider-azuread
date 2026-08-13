package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdatableAssetError interface {
	WindowsUpdatesUpdatableAssetError() BaseWindowsUpdatesUpdatableAssetErrorImpl
}

var _ WindowsUpdatesUpdatableAssetError = BaseWindowsUpdatesUpdatableAssetErrorImpl{}

type BaseWindowsUpdatesUpdatableAssetErrorImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsUpdatesUpdatableAssetErrorImpl) WindowsUpdatesUpdatableAssetError() BaseWindowsUpdatesUpdatableAssetErrorImpl {
	return s
}

var _ WindowsUpdatesUpdatableAssetError = RawWindowsUpdatesUpdatableAssetErrorImpl{}

// RawWindowsUpdatesUpdatableAssetErrorImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawWindowsUpdatesUpdatableAssetErrorImpl struct {
	windowsUpdatesUpdatableAssetError BaseWindowsUpdatesUpdatableAssetErrorImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawWindowsUpdatesUpdatableAssetErrorImpl) WindowsUpdatesUpdatableAssetError() BaseWindowsUpdatesUpdatableAssetErrorImpl {
	return s.windowsUpdatesUpdatableAssetError
}

func (s RawWindowsUpdatesUpdatableAssetErrorImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalWindowsUpdatesUpdatableAssetErrorImplementation(input []byte) (WindowsUpdatesUpdatableAssetError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesUpdatableAssetError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.azureADDeviceRegistrationError") {
		var out WindowsUpdatesAzureADDeviceRegistrationError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesAzureADDeviceRegistrationError: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesUpdatableAssetErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesUpdatableAssetErrorImpl: %+v", err)
	}

	return RawWindowsUpdatesUpdatableAssetErrorImpl{
		windowsUpdatesUpdatableAssetError: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
