package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = WindowsUniversalAppXAppAssignmentSettings{}

type WindowsUniversalAppXAppAssignmentSettings struct {
	// If true, uses device execution context for Windows Universal AppX mobile app. Device-context install is not allowed
	// when this type of app is targeted with Available intent. Defaults to false.
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUniversalAppXAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUniversalAppXAppAssignmentSettings{}

func (s WindowsUniversalAppXAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUniversalAppXAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUniversalAppXAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUniversalAppXAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUniversalAppXAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUniversalAppXAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
