package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10CustomConfiguration{}

type Windows10CustomConfiguration struct {
	// OMA settings. This collection can contain a maximum of 1000 elements.
	OmaSettings *[]OmaSetting `json:"omaSettings,omitempty"`

	// Fields inherited from DeviceConfiguration

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

func (s Windows10CustomConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s Windows10CustomConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10CustomConfiguration{}

func (s Windows10CustomConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10CustomConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10CustomConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10CustomConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10CustomConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10CustomConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Windows10CustomConfiguration{}

func (s *Windows10CustomConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Assignments                 *[]DeviceConfigurationAssignment   `json:"assignments,omitempty"`
		CreatedDateTime             *string                            `json:"createdDateTime,omitempty"`
		Description                 nullable.Type[string]              `json:"description,omitempty"`
		DeviceSettingStateSummaries *[]SettingStateDeviceSummary       `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview        *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses              *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
		DisplayName                 *string                            `json:"displayName,omitempty"`
		LastModifiedDateTime        *string                            `json:"lastModifiedDateTime,omitempty"`
		UserStatusOverview          *DeviceConfigurationUserOverview   `json:"userStatusOverview,omitempty"`
		UserStatuses                *[]DeviceConfigurationUserStatus   `json:"userStatuses,omitempty"`
		Version                     *int64                             `json:"version,omitempty"`
		Id                          *string                            `json:"id,omitempty"`
		ODataId                     *string                            `json:"@odata.id,omitempty"`
		ODataType                   *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Windows10CustomConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["omaSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OmaSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]OmaSetting, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalOmaSettingImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OmaSettings' for 'Windows10CustomConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OmaSettings = &output
	}

	return nil
}
