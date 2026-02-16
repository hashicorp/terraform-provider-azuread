package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RoleDefinition = DeviceAndAppManagementRoleDefinition{}

type DeviceAndAppManagementRoleDefinition struct {

	// Fields inherited from RoleDefinition

	// Description of the Role definition.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name of the Role definition.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`

	// List of Role assignments for this role definition.
	RoleAssignments *[]RoleAssignment `json:"roleAssignments,omitempty"`

	// List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of
	// the rolePermission.
	RolePermissions *[]RolePermission `json:"rolePermissions,omitempty"`

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

func (s DeviceAndAppManagementRoleDefinition) RoleDefinition() BaseRoleDefinitionImpl {
	return BaseRoleDefinitionImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		IsBuiltIn:       s.IsBuiltIn,
		RoleAssignments: s.RoleAssignments,
		RolePermissions: s.RolePermissions,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s DeviceAndAppManagementRoleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceAndAppManagementRoleDefinition{}

func (s DeviceAndAppManagementRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementRoleDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceAndAppManagementRoleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementRoleDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceAndAppManagementRoleDefinition{}

func (s *DeviceAndAppManagementRoleDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description     nullable.Type[string] `json:"description,omitempty"`
		DisplayName     nullable.Type[string] `json:"displayName,omitempty"`
		IsBuiltIn       *bool                 `json:"isBuiltIn,omitempty"`
		RolePermissions *[]RolePermission     `json:"rolePermissions,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IsBuiltIn = decoded.IsBuiltIn
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RolePermissions = decoded.RolePermissions

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceAndAppManagementRoleDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleAssignments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RoleAssignments into list []json.RawMessage: %+v", err)
		}

		output := make([]RoleAssignment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRoleAssignmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RoleAssignments' for 'DeviceAndAppManagementRoleDefinition': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RoleAssignments = &output
	}

	return nil
}
