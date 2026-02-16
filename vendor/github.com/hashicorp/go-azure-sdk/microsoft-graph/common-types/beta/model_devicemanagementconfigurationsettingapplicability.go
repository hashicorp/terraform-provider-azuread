package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicability interface {
	DeviceManagementConfigurationSettingApplicability() BaseDeviceManagementConfigurationSettingApplicabilityImpl
}

var _ DeviceManagementConfigurationSettingApplicability = BaseDeviceManagementConfigurationSettingApplicabilityImpl{}

type BaseDeviceManagementConfigurationSettingApplicabilityImpl struct {
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

func (s BaseDeviceManagementConfigurationSettingApplicabilityImpl) DeviceManagementConfigurationSettingApplicability() BaseDeviceManagementConfigurationSettingApplicabilityImpl {
	return s
}

var _ DeviceManagementConfigurationSettingApplicability = RawDeviceManagementConfigurationSettingApplicabilityImpl{}

// RawDeviceManagementConfigurationSettingApplicabilityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingApplicabilityImpl struct {
	deviceManagementConfigurationSettingApplicability BaseDeviceManagementConfigurationSettingApplicabilityImpl
	Type                                              string
	Values                                            map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingApplicabilityImpl) DeviceManagementConfigurationSettingApplicability() BaseDeviceManagementConfigurationSettingApplicabilityImpl {
	return s.deviceManagementConfigurationSettingApplicability
}

func UnmarshalDeviceManagementConfigurationSettingApplicabilityImplementation(input []byte) (DeviceManagementConfigurationSettingApplicability, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingApplicability into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationApplicationSettingApplicability") {
		var out DeviceManagementConfigurationApplicationSettingApplicability
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationApplicationSettingApplicability: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationExchangeOnlineSettingApplicability") {
		var out DeviceManagementConfigurationExchangeOnlineSettingApplicability
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationExchangeOnlineSettingApplicability: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationWindowsSettingApplicability") {
		var out DeviceManagementConfigurationWindowsSettingApplicability
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationWindowsSettingApplicability: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingApplicabilityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingApplicabilityImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingApplicabilityImpl{
		deviceManagementConfigurationSettingApplicability: parent,
		Type:   value,
		Values: temp,
	}, nil

}
