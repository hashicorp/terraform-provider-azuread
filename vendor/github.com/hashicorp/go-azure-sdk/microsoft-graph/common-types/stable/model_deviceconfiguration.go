package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfiguration interface {
	Entity
	DeviceConfiguration() BaseDeviceConfigurationImpl
}

var _ DeviceConfiguration = BaseDeviceConfigurationImpl{}

type BaseDeviceConfigurationImpl struct {
	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s
}

func (s BaseDeviceConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceConfiguration = RawDeviceConfigurationImpl{}

// RawDeviceConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceConfigurationImpl struct {
	deviceConfiguration BaseDeviceConfigurationImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawDeviceConfigurationImpl) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return s.deviceConfiguration
}

func (s RawDeviceConfigurationImpl) Entity() BaseEntityImpl {
	return s.deviceConfiguration.Entity()
}

var _ json.Marshaler = BaseDeviceConfigurationImpl{}

func (s BaseDeviceConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceConfigurationImplementation(input []byte) (DeviceConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.androidCustomConfiguration") {
		var out AndroidCustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidCustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidGeneralDeviceConfiguration") {
		var out AndroidGeneralDeviceConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidGeneralDeviceConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileCustomConfiguration") {
		var out AndroidWorkProfileCustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileCustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.androidWorkProfileGeneralDeviceConfiguration") {
		var out AndroidWorkProfileGeneralDeviceConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AndroidWorkProfileGeneralDeviceConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appleDeviceFeaturesConfigurationBase") {
		var out AppleDeviceFeaturesConfigurationBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleDeviceFeaturesConfigurationBase: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.editionUpgradeConfiguration") {
		var out EditionUpgradeConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EditionUpgradeConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCertificateProfile") {
		var out IosCertificateProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCertificateProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosCustomConfiguration") {
		var out IosCustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosCustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosGeneralDeviceConfiguration") {
		var out IosGeneralDeviceConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosGeneralDeviceConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.iosUpdateConfiguration") {
		var out IosUpdateConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IosUpdateConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSCustomConfiguration") {
		var out MacOSCustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSCustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.macOSGeneralDeviceConfiguration") {
		var out MacOSGeneralDeviceConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MacOSGeneralDeviceConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharedPCConfiguration") {
		var out SharedPCConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedPCConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10CustomConfiguration") {
		var out Windows10CustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10CustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EndpointProtectionConfiguration") {
		var out Windows10EndpointProtectionConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EndpointProtectionConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EnterpriseModernAppManagementConfiguration") {
		var out Windows10EnterpriseModernAppManagementConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnterpriseModernAppManagementConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10GeneralConfiguration") {
		var out Windows10GeneralConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10GeneralConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10SecureAssessmentConfiguration") {
		var out Windows10SecureAssessmentConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10SecureAssessmentConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10TeamGeneralConfiguration") {
		var out Windows10TeamGeneralConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10TeamGeneralConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows81GeneralConfiguration") {
		var out Windows81GeneralConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows81GeneralConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration") {
		var out WindowsDefenderAdvancedThreatProtectionConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsDefenderAdvancedThreatProtectionConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81CustomConfiguration") {
		var out WindowsPhone81CustomConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81CustomConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsPhone81GeneralConfiguration") {
		var out WindowsPhone81GeneralConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsPhone81GeneralConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdateForBusinessConfiguration") {
		var out WindowsUpdateForBusinessConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdateForBusinessConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceConfigurationImpl: %+v", err)
	}

	return RawDeviceConfigurationImpl{
		deviceConfiguration: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
