package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceAppManagementTask = UnmanagedDeviceDiscoveryTask{}

type UnmanagedDeviceDiscoveryTask struct {
	// Unmanaged devices discovered in the network.
	UnmanagedDevices *[]UnmanagedDevice `json:"unmanagedDevices,omitempty"`

	// Fields inherited from DeviceAppManagementTask

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

func (s UnmanagedDeviceDiscoveryTask) DeviceAppManagementTask() BaseDeviceAppManagementTaskImpl {
	return BaseDeviceAppManagementTaskImpl{
		AssignedTo:      s.AssignedTo,
		Category:        s.Category,
		CreatedDateTime: s.CreatedDateTime,
		Creator:         s.Creator,
		CreatorNotes:    s.CreatorNotes,
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DueDateTime:     s.DueDateTime,
		Priority:        s.Priority,
		Status:          s.Status,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s UnmanagedDeviceDiscoveryTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnmanagedDeviceDiscoveryTask{}

func (s UnmanagedDeviceDiscoveryTask) MarshalJSON() ([]byte, error) {
	type wrapper UnmanagedDeviceDiscoveryTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnmanagedDeviceDiscoveryTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnmanagedDeviceDiscoveryTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unmanagedDeviceDiscoveryTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnmanagedDeviceDiscoveryTask: %+v", err)
	}

	return encoded, nil
}
