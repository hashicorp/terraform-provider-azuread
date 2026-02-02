package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettings interface {
	MobileAppAssignmentSettings
	Win32LobAppAssignmentSettings() BaseWin32LobAppAssignmentSettingsImpl
}

var _ Win32LobAppAssignmentSettings = BaseWin32LobAppAssignmentSettingsImpl{}

type BaseWin32LobAppAssignmentSettingsImpl struct {
	// The auto-update settings to apply for this app assignment.
	AutoUpdateSettings *Win32LobAppAutoUpdateSettings `json:"autoUpdateSettings,omitempty"`

	// Contains value for delivery optimization priority.
	DeliveryOptimizationPriority *Win32LobAppDeliveryOptimizationPriority `json:"deliveryOptimizationPriority,omitempty"`

	// The install time settings to apply for this app assignment.
	InstallTimeSettings *MobileAppInstallTimeSettings `json:"installTimeSettings,omitempty"`

	// Contains value for notification status.
	Notifications *Win32LobAppNotification `json:"notifications,omitempty"`

	// The reboot settings to apply for this app assignment.
	RestartSettings *Win32LobAppRestartSettings `json:"restartSettings,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWin32LobAppAssignmentSettingsImpl) Win32LobAppAssignmentSettings() BaseWin32LobAppAssignmentSettingsImpl {
	return s
}

func (s BaseWin32LobAppAssignmentSettingsImpl) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Win32LobAppAssignmentSettings = RawWin32LobAppAssignmentSettingsImpl{}

// RawWin32LobAppAssignmentSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWin32LobAppAssignmentSettingsImpl struct {
	win32LobAppAssignmentSettings BaseWin32LobAppAssignmentSettingsImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawWin32LobAppAssignmentSettingsImpl) Win32LobAppAssignmentSettings() BaseWin32LobAppAssignmentSettingsImpl {
	return s.win32LobAppAssignmentSettings
}

func (s RawWin32LobAppAssignmentSettingsImpl) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return s.win32LobAppAssignmentSettings.MobileAppAssignmentSettings()
}

var _ json.Marshaler = BaseWin32LobAppAssignmentSettingsImpl{}

func (s BaseWin32LobAppAssignmentSettingsImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWin32LobAppAssignmentSettingsImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWin32LobAppAssignmentSettingsImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWin32LobAppAssignmentSettingsImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWin32LobAppAssignmentSettingsImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWin32LobAppAssignmentSettingsImplementation(input []byte) (Win32LobAppAssignmentSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppAssignmentSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32CatalogAppAssignmentSettings") {
		var out Win32CatalogAppAssignmentSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32CatalogAppAssignmentSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseWin32LobAppAssignmentSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWin32LobAppAssignmentSettingsImpl: %+v", err)
	}

	return RawWin32LobAppAssignmentSettingsImpl{
		win32LobAppAssignmentSettings: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
