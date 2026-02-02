package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingApplicability = DeviceManagementConfigurationExchangeOnlineSettingApplicability{}

type DeviceManagementConfigurationExchangeOnlineSettingApplicability struct {

	// Fields inherited from DeviceManagementConfigurationSettingApplicability

	// description of the setting
	Description nullable.Type[string] `json:"description,omitempty"`

	// Describes applicability for the mode the device is in
	DeviceMode *DeviceManagementConfigurationDeviceMode `json:"deviceMode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Supported platform types.
	Platform *DeviceManagementConfigurationPlatforms `json:"platform,omitempty"`

	// Describes which technology this setting can be deployed with
	Technologies *DeviceManagementConfigurationTechnologies `json:"technologies,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationExchangeOnlineSettingApplicability) DeviceManagementConfigurationSettingApplicability() BaseDeviceManagementConfigurationSettingApplicabilityImpl {
	return BaseDeviceManagementConfigurationSettingApplicabilityImpl{
		Description:  s.Description,
		DeviceMode:   s.DeviceMode,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
		Platform:     s.Platform,
		Technologies: s.Technologies,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationExchangeOnlineSettingApplicability{}

func (s DeviceManagementConfigurationExchangeOnlineSettingApplicability) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationExchangeOnlineSettingApplicability
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationExchangeOnlineSettingApplicability: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationExchangeOnlineSettingApplicability: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationExchangeOnlineSettingApplicability"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationExchangeOnlineSettingApplicability: %+v", err)
	}

	return encoded, nil
}
