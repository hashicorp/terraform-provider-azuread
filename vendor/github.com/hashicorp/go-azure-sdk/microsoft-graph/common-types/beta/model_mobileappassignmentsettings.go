package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppAssignmentSettings interface {
	MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl
}

var _ MobileAppAssignmentSettings = BaseMobileAppAssignmentSettingsImpl{}

type BaseMobileAppAssignmentSettingsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMobileAppAssignmentSettingsImpl) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return s
}

var _ MobileAppAssignmentSettings = RawMobileAppAssignmentSettingsImpl{}

// RawMobileAppAssignmentSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileAppAssignmentSettingsImpl struct {
	mobileAppAssignmentSettings BaseMobileAppAssignmentSettingsImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawMobileAppAssignmentSettingsImpl) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return s.mobileAppAssignmentSettings
}

func UnmarshalMobileAppAssignmentSettingsImplementation(input []byte) (MobileAppAssignmentSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppAssignmentSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidManagedStoreAppAssignmentSettings") {
		var out AndroidManagedStoreAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidManagedStoreAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosDdmLobAppAssignmentSettings") {
		var out IosDdmLobAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosDdmLobAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosLobAppAssignmentSettings") {
		var out IosLobAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosLobAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosStoreAppAssignmentSettings") {
		var out IosStoreAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosStoreAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosVppAppAssignmentSettings") {
		var out IosVppAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosVppAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOsLobAppAssignmentSettings") {
		var out MacOsLobAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOsLobAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOsVppAppAssignmentSettings") {
		var out MacOsVppAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOsVppAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftStoreForBusinessAppAssignmentSettings") {
		var out MicrosoftStoreForBusinessAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftStoreForBusinessAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppAssignmentSettings") {
		var out Win32LobAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.winGetAppAssignmentSettings") {
		var out WinGetAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WinGetAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsAppXAppAssignmentSettings") {
		var out WindowsAppXAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsAppXAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUniversalAppXAppAssignmentSettings") {
		var out WindowsUniversalAppXAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUniversalAppXAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileAppAssignmentSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileAppAssignmentSettingsImpl: %+v", err)
	}

	return RawMobileAppAssignmentSettingsImpl{
		mobileAppAssignmentSettings: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
