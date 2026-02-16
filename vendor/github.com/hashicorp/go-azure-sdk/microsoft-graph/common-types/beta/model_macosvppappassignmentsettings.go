package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppAssignmentSettings = MacOsVppAppAssignmentSettings{}

type MacOsVppAppAssignmentSettings struct {
	// When TRUE, indicates that the app should not be automatically updated with the latest version from Apple app store.
	// When FALSE, indicates that the app may be auto updated. By default, this property is set to null which internally is
	// treated as FALSE.
	PreventAutoAppUpdate nullable.Type[bool] `json:"preventAutoAppUpdate,omitempty"`

	// When TRUE, indicates that the app should not be backed up to iCloud. When FALSE, indicates that the app may be backed
	// up to iCloud. By default, this property is set to null which internally is treated as FALSE.
	PreventManagedAppBackup nullable.Type[bool] `json:"preventManagedAppBackup,omitempty"`

	// Whether or not to uninstall the app when device is removed from Intune.
	UninstallOnDeviceRemoval nullable.Type[bool] `json:"uninstallOnDeviceRemoval,omitempty"`

	// Whether or not to use device licensing.
	UseDeviceLicensing *bool `json:"useDeviceLicensing,omitempty"`

	// Fields inherited from MobileAppAssignmentSettings

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MacOsVppAppAssignmentSettings) MobileAppAssignmentSettings() BaseMobileAppAssignmentSettingsImpl {
	return BaseMobileAppAssignmentSettingsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOsVppAppAssignmentSettings{}

func (s MacOsVppAppAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper MacOsVppAppAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOsVppAppAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOsVppAppAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOsVppAppAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOsVppAppAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
