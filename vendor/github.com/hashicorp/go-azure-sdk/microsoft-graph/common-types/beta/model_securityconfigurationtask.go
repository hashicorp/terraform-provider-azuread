package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceAppManagementTask = SecurityConfigurationTask{}

type SecurityConfigurationTask struct {
	// The endpoint security configuration applicable platform.
	ApplicablePlatform *EndpointSecurityConfigurationApplicablePlatform `json:"applicablePlatform,omitempty"`

	// The endpoint security policy type.
	EndpointSecurityPolicy *EndpointSecurityConfigurationType `json:"endpointSecurityPolicy,omitempty"`

	// The endpoint security policy profile type.
	EndpointSecurityPolicyProfile *EndpointSecurityConfigurationProfileType `json:"endpointSecurityPolicyProfile,omitempty"`

	// Information about the mitigation.
	Insights nullable.Type[string] `json:"insights,omitempty"`

	// The intended settings and their values.
	IntendedSettings *[]KeyValuePair `json:"intendedSettings,omitempty"`

	// The number of vulnerable devices. Valid values 0 to 65536
	ManagedDeviceCount *int64 `json:"managedDeviceCount,omitempty"`

	// The vulnerable managed devices.
	ManagedDevices *[]VulnerableManagedDevice `json:"managedDevices,omitempty"`

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

func (s SecurityConfigurationTask) DeviceAppManagementTask() BaseDeviceAppManagementTaskImpl {
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

func (s SecurityConfigurationTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityConfigurationTask{}

func (s SecurityConfigurationTask) MarshalJSON() ([]byte, error) {
	type wrapper SecurityConfigurationTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityConfigurationTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityConfigurationTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityConfigurationTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityConfigurationTask: %+v", err)
	}

	return encoded, nil
}
