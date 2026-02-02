package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RoleAssignment = DeviceAndAppManagementRoleAssignment{}

type DeviceAndAppManagementRoleAssignment struct {
	// Indicates the list of role member security group Entra IDs. For example: {dec942f4-6777-4998-96b4-522e383b08e2}.
	Members *[]string `json:"members,omitempty"`

	// Indicates the set of role scope tag IDs for the role assignment. These scope tags will limit the visibility of any
	// Intune resources to those that match any of the scope tags in this collection.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates the set of scope tags for the role assignment. These scope tags will limit the visibility of any Intune
	// resources to those that match any of the scope tags in this collection.
	RoleScopeTags *[]RoleScopeTag `json:"roleScopeTags,omitempty"`

	// Fields inherited from RoleAssignment

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

func (s DeviceAndAppManagementRoleAssignment) RoleAssignment() BaseRoleAssignmentImpl {
	return BaseRoleAssignmentImpl{
		Description:    s.Description,
		DisplayName:    s.DisplayName,
		ResourceScopes: s.ResourceScopes,
		RoleDefinition: s.RoleDefinition,
		ScopeMembers:   s.ScopeMembers,
		ScopeType:      s.ScopeType,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s DeviceAndAppManagementRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceAndAppManagementRoleAssignment{}

func (s DeviceAndAppManagementRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceAndAppManagementRoleAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementRoleAssignment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceAndAppManagementRoleAssignment{}

func (s *DeviceAndAppManagementRoleAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Members         *[]string                `json:"members,omitempty"`
		RoleScopeTagIds *[]string                `json:"roleScopeTagIds,omitempty"`
		RoleScopeTags   *[]RoleScopeTag          `json:"roleScopeTags,omitempty"`
		Description     nullable.Type[string]    `json:"description,omitempty"`
		DisplayName     nullable.Type[string]    `json:"displayName,omitempty"`
		ResourceScopes  *[]string                `json:"resourceScopes,omitempty"`
		ScopeMembers    *[]string                `json:"scopeMembers,omitempty"`
		ScopeType       *RoleAssignmentScopeType `json:"scopeType,omitempty"`
		Id              *string                  `json:"id,omitempty"`
		ODataId         *string                  `json:"@odata.id,omitempty"`
		ODataType       *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Members = decoded.Members
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.RoleScopeTags = decoded.RoleScopeTags
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceScopes = decoded.ResourceScopes
	s.ScopeMembers = decoded.ScopeMembers
	s.ScopeType = decoded.ScopeType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceAndAppManagementRoleAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleDefinition"]; ok {
		impl, err := UnmarshalRoleDefinitionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RoleDefinition' for 'DeviceAndAppManagementRoleAssignment': %+v", err)
		}
		s.RoleDefinition = &impl
	}

	return nil
}
