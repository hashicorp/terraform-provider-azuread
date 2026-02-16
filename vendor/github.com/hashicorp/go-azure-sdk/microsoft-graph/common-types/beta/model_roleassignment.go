package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignment interface {
	Entity
	RoleAssignment() BaseRoleAssignmentImpl
}

var _ RoleAssignment = BaseRoleAssignmentImpl{}

type BaseRoleAssignmentImpl struct {
	// Indicates the description of the role assignment. For example: 'All administrators, employees and scope tags
	// associated with the Houston office.' Max length is 1024 characters.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates the display name of the role assignment. For example: 'Houston administrators and users'. Max length is 128
	// characters.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates the list of resource scope security group Entra IDs. For example: {dec942f4-6777-4998-96b4-522e383b08e2}.
	ResourceScopes *[]string `json:"resourceScopes,omitempty"`

	// Indicates the role definition for this role assignment.
	RoleDefinition *RoleDefinition `json:"roleDefinition,omitempty"`

	// Indicates the list of role scope member security groups Entra IDs. For example,
	// {dec942f4-6777-4998-96b4-522e383b08e2}.
	ScopeMembers *[]string `json:"scopeMembers,omitempty"`

	// Specifies the type of scope for a Role Assignment.
	ScopeType *RoleAssignmentScopeType `json:"scopeType,omitempty"`

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

func (s BaseRoleAssignmentImpl) RoleAssignment() BaseRoleAssignmentImpl {
	return s
}

func (s BaseRoleAssignmentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RoleAssignment = RawRoleAssignmentImpl{}

// RawRoleAssignmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRoleAssignmentImpl struct {
	roleAssignment BaseRoleAssignmentImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawRoleAssignmentImpl) RoleAssignment() BaseRoleAssignmentImpl {
	return s.roleAssignment
}

func (s RawRoleAssignmentImpl) Entity() BaseEntityImpl {
	return s.roleAssignment.Entity()
}

var _ json.Marshaler = BaseRoleAssignmentImpl{}

func (s BaseRoleAssignmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRoleAssignmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRoleAssignmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRoleAssignmentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.roleAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRoleAssignmentImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseRoleAssignmentImpl{}

func (s *BaseRoleAssignmentImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description    nullable.Type[string]    `json:"description,omitempty"`
		DisplayName    nullable.Type[string]    `json:"displayName,omitempty"`
		ResourceScopes *[]string                `json:"resourceScopes,omitempty"`
		ScopeMembers   *[]string                `json:"scopeMembers,omitempty"`
		ScopeType      *RoleAssignmentScopeType `json:"scopeType,omitempty"`
		Id             *string                  `json:"id,omitempty"`
		ODataId        *string                  `json:"@odata.id,omitempty"`
		ODataType      *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ResourceScopes = decoded.ResourceScopes
	s.ScopeMembers = decoded.ScopeMembers
	s.ScopeType = decoded.ScopeType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRoleAssignmentImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleDefinition"]; ok {
		impl, err := UnmarshalRoleDefinitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RoleDefinition' for 'BaseRoleAssignmentImpl': %+v", err)
		}
		s.RoleDefinition = &impl
	}

	return nil
}

func UnmarshalRoleAssignmentImplementation(input []byte) (RoleAssignment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleAssignment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceAndAppManagementRoleAssignment") {
		var out DeviceAndAppManagementRoleAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceAndAppManagementRoleAssignment: %+v", err)
		}
		return out, nil
	}

	var parent BaseRoleAssignmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRoleAssignmentImpl: %+v", err)
	}

	return RawRoleAssignmentImpl{
		roleAssignment: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
