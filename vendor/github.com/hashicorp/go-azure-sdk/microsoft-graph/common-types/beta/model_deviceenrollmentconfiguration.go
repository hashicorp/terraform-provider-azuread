package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentConfiguration interface {
	Entity
	DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl
}

var _ DeviceEnrollmentConfiguration = BaseDeviceEnrollmentConfigurationImpl{}

type BaseDeviceEnrollmentConfigurationImpl struct {
	// The list of group assignments for the device configuration profile
	Assignments *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`

	// Created date time in UTC of the device enrollment configuration
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the device enrollment configuration
	Description nullable.Type[string] `json:"description,omitempty"`

	// Describes the TemplateFamily for the Template entity
	DeviceEnrollmentConfigurationType *DeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`

	// The display name of the device enrollment configuration
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Last modified date time in UTC of the device enrollment configuration
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Priority is used when a user exists in multiple groups that are assigned enrollment configuration. Users are subject
	// only to the configuration with the lowest priority value.
	Priority *int64 `json:"priority,omitempty"`

	// Optional role scope tags for the enrollment restrictions.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// The version of the device enrollment configuration
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

func (s BaseDeviceEnrollmentConfigurationImpl) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return s
}

func (s BaseDeviceEnrollmentConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceEnrollmentConfiguration = RawDeviceEnrollmentConfigurationImpl{}

// RawDeviceEnrollmentConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceEnrollmentConfigurationImpl struct {
	deviceEnrollmentConfiguration BaseDeviceEnrollmentConfigurationImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawDeviceEnrollmentConfigurationImpl) DeviceEnrollmentConfiguration() BaseDeviceEnrollmentConfigurationImpl {
	return s.deviceEnrollmentConfiguration
}

func (s RawDeviceEnrollmentConfigurationImpl) Entity() BaseEntityImpl {
	return s.deviceEnrollmentConfiguration.Entity()
}

var _ json.Marshaler = BaseDeviceEnrollmentConfigurationImpl{}

func (s BaseDeviceEnrollmentConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceEnrollmentConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceEnrollmentConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceEnrollmentConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceEnrollmentConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceEnrollmentConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceEnrollmentConfigurationImplementation(input []byte) (DeviceEnrollmentConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceEnrollmentConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComanagementAuthorityConfiguration") {
		var out DeviceComanagementAuthorityConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComanagementAuthorityConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentLimitConfiguration") {
		var out DeviceEnrollmentLimitConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentLimitConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentNotificationConfiguration") {
		var out DeviceEnrollmentNotificationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentNotificationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentPlatformRestrictionConfiguration") {
		var out DeviceEnrollmentPlatformRestrictionConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentPlatformRestrictionConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentPlatformRestrictionsConfiguration") {
		var out DeviceEnrollmentPlatformRestrictionsConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentPlatformRestrictionsConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceEnrollmentWindowsHelloForBusinessConfiguration") {
		var out DeviceEnrollmentWindowsHelloForBusinessConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceEnrollmentWindowsHelloForBusinessConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration") {
		var out Windows10EnrollmentCompletionPageConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Windows10EnrollmentCompletionPageConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsRestoreDeviceEnrollmentConfiguration") {
		var out WindowsRestoreDeviceEnrollmentConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsRestoreDeviceEnrollmentConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceEnrollmentConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceEnrollmentConfigurationImpl: %+v", err)
	}

	return RawDeviceEnrollmentConfigurationImpl{
		deviceEnrollmentConfiguration: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
