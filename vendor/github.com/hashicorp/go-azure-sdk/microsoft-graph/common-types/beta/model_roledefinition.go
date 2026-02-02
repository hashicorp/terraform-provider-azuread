package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleDefinition interface {
	Entity
	RoleDefinition() BaseRoleDefinitionImpl
}

var _ RoleDefinition = BaseRoleDefinitionImpl{}

type BaseRoleDefinitionImpl struct {
	// Description of the Role definition.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display Name of the Role definition.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`

	// Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
	IsBuiltInRoleDefinition *bool `json:"isBuiltInRoleDefinition,omitempty"`

	// List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of
	// the rolePermission.
	Permissions *[]RolePermission `json:"permissions,omitempty"`

	// List of Role assignments for this role definition.
	RoleAssignments *[]RoleAssignment `json:"roleAssignments,omitempty"`

	// List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of
	// the rolePermission.
	RolePermissions *[]RolePermission `json:"rolePermissions,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

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

func (s BaseRoleDefinitionImpl) RoleDefinition() BaseRoleDefinitionImpl {
	return s
}

func (s BaseRoleDefinitionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RoleDefinition = RawRoleDefinitionImpl{}

// RawRoleDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRoleDefinitionImpl struct {
	roleDefinition BaseRoleDefinitionImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawRoleDefinitionImpl) RoleDefinition() BaseRoleDefinitionImpl {
	return s.roleDefinition
}

func (s RawRoleDefinitionImpl) Entity() BaseEntityImpl {
	return s.roleDefinition.Entity()
}

var _ json.Marshaler = BaseRoleDefinitionImpl{}

func (s BaseRoleDefinitionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRoleDefinitionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRoleDefinitionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRoleDefinitionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.roleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRoleDefinitionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseRoleDefinitionImpl{}

func (s *BaseRoleDefinitionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description             nullable.Type[string] `json:"description,omitempty"`
		DisplayName             nullable.Type[string] `json:"displayName,omitempty"`
		IsBuiltIn               *bool                 `json:"isBuiltIn,omitempty"`
		IsBuiltInRoleDefinition *bool                 `json:"isBuiltInRoleDefinition,omitempty"`
		Permissions             *[]RolePermission     `json:"permissions,omitempty"`
		RolePermissions         *[]RolePermission     `json:"rolePermissions,omitempty"`
		RoleScopeTagIds         *[]string             `json:"roleScopeTagIds,omitempty"`
		Id                      *string               `json:"id,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsBuiltIn = decoded.IsBuiltIn
	s.IsBuiltInRoleDefinition = decoded.IsBuiltInRoleDefinition
	s.Permissions = decoded.Permissions
	s.RolePermissions = decoded.RolePermissions
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRoleDefinitionImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'RoleAssignments' for 'BaseRoleDefinitionImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RoleAssignments = &output
	}

	return nil
}

func UnmarshalRoleDefinitionImplementation(input []byte) (RoleDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementRoleDefinition") {
		var out DeviceAndAppManagementRoleDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BaseRoleDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRoleDefinitionImpl: %+v", err)
	}

	return RawRoleDefinitionImpl{
		roleDefinition: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
