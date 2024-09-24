package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HardwareConfiguration{}

type HardwareConfiguration struct {
	// A list of the Entra user group ids that hardware configuration will be applied to. Only security groups and Office
	// 365 Groups are supported. Optional.
	Assignments *[]HardwareConfigurationAssignment `json:"assignments,omitempty"`

	// The file content contains custom hardware settings that will be applied to the assigned devices' BIOS. Max allowed
	// file size is 5KB. Represented as bytes. Required.
	ConfigurationFileContent string `json:"configurationFileContent"`

	// The date and time of when the BIOS configuration profile was created. The value cannot be modified and is
	// automatically populated when the device is enrolled. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-Only. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the hardware configuration. Use this to provide context, purpose, applications, etc of the BIOS
	// configuration profile for your organization's admins. Max length is 1000 characters. Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// List of run states for the hardware configuration across all devices. Read-Only.
	DeviceRunStates *[]HardwareConfigurationDeviceState `json:"deviceRunStates,omitempty"`

	// The name of the hardware BIOS configuration profile. It serves as user-friendly name to identify hardware BIOS
	// configuration profiles. Max length is 150 characters. Required. Read-Only.
	DisplayName string `json:"displayName"`

	// The file name for the BIOS configuration profile's ConfigurationFileContent. Max length is 150 characters. Required.
	FileName string `json:"fileName"`

	// Indicates the supported oems of hardware configuration
	HardwareConfigurationFormat *HardwareConfigurationFormat `json:"hardwareConfigurationFormat,omitempty"`

	// The date and time of when the BIOS configuration profile was last modified. The value cannot be modified and is
	// automatically populated when the device is enrolled. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-Only. Read-Only. This property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// When TRUE, indicates whether the policy-assigned devices' passwords are disabled. When FALSE, indicates they are
	// enabled. Default is FALSE. Required.
	PerDevicePasswordDisabled bool `json:"perDevicePasswordDisabled"`

	// A list of unique Scope Tag IDs associated with the hardware configuration. Optional.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// A summary of the results from an attempt to configure hardware settings. Read-Only.
	RunSummary *HardwareConfigurationRunSummary `json:"runSummary,omitempty"`

	// List of run states for the hardware configuration across all users. Read-Only.
	UserRunStates *[]HardwareConfigurationUserState `json:"userRunStates,omitempty"`

	// The version of the hardware configuration (E.g. 1, 2, 3 ...). This is incremented after a change to the BIOS
	// configuration profile's settings file name (FileName property), settings file content (ConfigurationFileContent
	// property), or the PerDevicePasswordDisabled property. Read-Only.
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

func (s HardwareConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HardwareConfiguration{}

func (s HardwareConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper HardwareConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HardwareConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HardwareConfiguration: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.hardwareConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HardwareConfiguration: %+v", err)
	}

	return encoded, nil
}
