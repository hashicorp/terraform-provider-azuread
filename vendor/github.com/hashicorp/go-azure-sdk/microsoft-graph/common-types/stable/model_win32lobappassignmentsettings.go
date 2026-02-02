package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = Win32LobAppAssignmentSettings{}

type Win32LobAppAssignmentSettings struct {
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

func (s Win32LobAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32LobAppAssignmentSettings{}

func (s Win32LobAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
