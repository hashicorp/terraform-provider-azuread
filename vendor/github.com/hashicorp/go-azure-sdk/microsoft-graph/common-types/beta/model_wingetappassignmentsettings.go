package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = WinGetAppAssignmentSettings{}

type WinGetAppAssignmentSettings struct {
	// The install time settings to apply for this app assignment.
	InstallTimeSettings *WinGetAppInstallTimeSettings `json:"installTimeSettings,omitempty"`

	// Contains value for notification status.
	Notifications *WinGetAppNotification `json:"notifications,omitempty"`

	// The reboot settings to apply for this app assignment.
	RestartSettings *WinGetAppRestartSettings `json:"restartSettings,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WinGetAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WinGetAppAssignmentSettings{}

func (s WinGetAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper WinGetAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WinGetAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WinGetAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.winGetAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WinGetAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
