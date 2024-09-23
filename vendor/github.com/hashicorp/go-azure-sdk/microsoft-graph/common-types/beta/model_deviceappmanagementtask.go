package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTask interface {
	Entity
	DeviceAppManagementTask() BaseDeviceAppManagementTaskImpl
}

var _ DeviceAppManagementTask = BaseDeviceAppManagementTaskImpl{}

type BaseDeviceAppManagementTaskImpl struct {
	// The name or email of the admin this task is assigned to.
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// Device app management task category.
	Category *DeviceAppManagementTaskCategory `json:"category,omitempty"`

	// The created date.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The email address of the creator.
	Creator nullable.Type[string] `json:"creator,omitempty"`

	// Notes from the creator.
	CreatorNotes nullable.Type[string] `json:"creatorNotes,omitempty"`

	// The description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The due date.
	DueDateTime *string `json:"dueDateTime,omitempty"`

	// Device app management task priority.
	Priority *DeviceAppManagementTaskPriority `json:"priority,omitempty"`

	// Device app management task status.
	Status *DeviceAppManagementTaskStatus `json:"status,omitempty"`

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

func (s BaseDeviceAppManagementTaskImpl) DeviceAppManagementTask() BaseDeviceAppManagementTaskImpl {
	return s
}

func (s BaseDeviceAppManagementTaskImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceAppManagementTask = RawDeviceAppManagementTaskImpl{}

// RawDeviceAppManagementTaskImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceAppManagementTaskImpl struct {
	deviceAppManagementTask BaseDeviceAppManagementTaskImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawDeviceAppManagementTaskImpl) DeviceAppManagementTask() BaseDeviceAppManagementTaskImpl {
	return s.deviceAppManagementTask
}

func (s RawDeviceAppManagementTaskImpl) Entity() BaseEntityImpl {
	return s.deviceAppManagementTask.Entity()
}

var _ json.Marshaler = BaseDeviceAppManagementTaskImpl{}

func (s BaseDeviceAppManagementTaskImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceAppManagementTaskImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceAppManagementTaskImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceAppManagementTaskImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceAppManagementTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceAppManagementTaskImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceAppManagementTaskImplementation(input []byte) (DeviceAppManagementTask, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAppManagementTask into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appVulnerabilityTask") {
		var out AppVulnerabilityTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppVulnerabilityTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityConfigurationTask") {
		var out SecurityConfigurationTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityConfigurationTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unmanagedDeviceDiscoveryTask") {
		var out UnmanagedDeviceDiscoveryTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnmanagedDeviceDiscoveryTask: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceAppManagementTaskImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceAppManagementTaskImpl: %+v", err)
	}

	return RawDeviceAppManagementTaskImpl{
		deviceAppManagementTask: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
