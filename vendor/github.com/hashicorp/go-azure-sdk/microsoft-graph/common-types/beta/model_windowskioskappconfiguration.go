package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppConfiguration interface {
	WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl
}

var _ WindowsKioskAppConfiguration = BaseWindowsKioskAppConfigurationImpl{}

type BaseWindowsKioskAppConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsKioskAppConfigurationImpl) WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl {
	return s
}

var _ WindowsKioskAppConfiguration = RawWindowsKioskAppConfigurationImpl{}

// RawWindowsKioskAppConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsKioskAppConfigurationImpl struct {
	windowsKioskAppConfiguration BaseWindowsKioskAppConfigurationImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawWindowsKioskAppConfigurationImpl) WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl {
	return s.windowsKioskAppConfiguration
}

func UnmarshalWindowsKioskAppConfigurationImplementation(input []byte) (WindowsKioskAppConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskAppConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskMultipleApps") {
		var out WindowsKioskMultipleApps
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskMultipleApps: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskSingleUWPApp") {
		var out WindowsKioskSingleUWPApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskSingleUWPApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskSingleWin32App") {
		var out WindowsKioskSingleWin32App
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskSingleWin32App: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsKioskAppConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsKioskAppConfigurationImpl: %+v", err)
	}

	return RawWindowsKioskAppConfigurationImpl{
		windowsKioskAppConfiguration: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
