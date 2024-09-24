package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = IosStoreAppAssignmentSettings{}

type IosStoreAppAssignmentSettings struct {
	// When TRUE, indicates that the app can be uninstalled by the user. When FALSE, indicates that the app cannot be
	// uninstalled by the user. By default, this property is set to null which internally is treated as TRUE.
	IsRemovable nullable.Type[bool] `json:"isRemovable,omitempty"`

	// When TRUE, indicates that the app should be uninstalled when the device is removed from Intune. When FALSE, indicates
	// that the app will not be uninstalled when the device is removed from Intune. By default, property is set to null
	// which internally is treated as TRUE.
	UninstallOnDeviceRemoval nullable.Type[bool] `json:"uninstallOnDeviceRemoval,omitempty"`

	// This is the unique identifier (Id) of the VPN Configuration to apply to the app.
	VpnConfigurationId nullable.Type[string] `json:"vpnConfigurationId,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosStoreAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosStoreAppAssignmentSettings{}

func (s IosStoreAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper IosStoreAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosStoreAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosStoreAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosStoreAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosStoreAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
