package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleAssignmentMultiple{}

type UnifiedRoleAssignmentMultiple struct {
	// Ids of the app specific scopes when the assignment scopes are app specific. The scopes of an assignment determine the
	// set of resources for which the principal has access. Directory scopes are shared scopes stored in the directory that
	// are understood by multiple applications. Use / for tenant-wide scope. App scopes are scopes that are defined and
	// understood by this application only.
	AppScopeIds *[]string `json:"appScopeIds,omitempty"`

	// Read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment
	// entity. Read-only.
	AppScopes *[]AppScope `json:"appScopes,omitempty"`

	Condition nullable.Type[string] `json:"condition,omitempty"`

	// Description of the role assignment.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Ids of the directory objects that represent the scopes of the assignment. The scopes of an assignment determine the
	// set of resources for which the principals have been granted access. Directory scopes are shared scopes stored in the
	// directory that are understood by multiple applications. App scopes are scopes that are defined and understood by this
	// application only.
	DirectoryScopeIds *[]string `json:"directoryScopeIds,omitempty"`

	// Read-only collection that references the directory objects that are scope of the assignment. Provided so that callers
	// can get the directory objects using $expand at the same time as getting the role assignment. Read-only. Supports
	// $expand.
	DirectoryScopes *[]DirectoryObject `json:"directoryScopes,omitempty"`

	// List of OData IDs for `DirectoryScopes` to bind to this entity
	DirectoryScopes_ODataBind *[]string `json:"directoryScopes@odata.bind,omitempty"`

	// Name of the role assignment. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identifiers of the principals to which the assignment is granted. Supports $filter (any operator only).
	PrincipalIds *[]string `json:"principalIds,omitempty"`

	// Read-only collection that references the assigned principals. Provided so that callers can get the principals using
	// $expand at the same time as getting the role assignment. Read-only. Supports $expand.
	Principals *[]DirectoryObject `json:"principals,omitempty"`

	// List of OData IDs for `Principals` to bind to this entity
	Principals_ODataBind *[]string `json:"principals@odata.bind,omitempty"`

	// Specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using
	// $expand at the same time as getting the role assignment. Supports $filter (eq operator on id, isBuiltIn, and
	// displayName, and startsWith operator on displayName) and $expand.
	RoleDefinition *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`

	// Identifier of the unifiedRoleDefinition the assignment is for.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

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

func (s UnifiedRoleAssignmentMultiple) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleAssignmentMultiple{}

func (s UnifiedRoleAssignmentMultiple) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleAssignmentMultiple
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleAssignmentMultiple: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleAssignmentMultiple: %+v", err)
	}

	delete(decoded, "appScopes")
	delete(decoded, "directoryScopes")
	delete(decoded, "principals")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleAssignmentMultiple"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleAssignmentMultiple: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleAssignmentMultiple{}

func (s *UnifiedRoleAssignmentMultiple) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppScopeIds               *[]string              `json:"appScopeIds,omitempty"`
		Condition                 nullable.Type[string]  `json:"condition,omitempty"`
		Description               nullable.Type[string]  `json:"description,omitempty"`
		DirectoryScopeIds         *[]string              `json:"directoryScopeIds,omitempty"`
		DirectoryScopes_ODataBind *[]string              `json:"directoryScopes@odata.bind,omitempty"`
		DisplayName               nullable.Type[string]  `json:"displayName,omitempty"`
		PrincipalIds              *[]string              `json:"principalIds,omitempty"`
		Principals_ODataBind      *[]string              `json:"principals@odata.bind,omitempty"`
		RoleDefinition            *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
		RoleDefinitionId          *string                `json:"roleDefinitionId,omitempty"`
		Id                        *string                `json:"id,omitempty"`
		ODataId                   *string                `json:"@odata.id,omitempty"`
		ODataType                 *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppScopeIds = decoded.AppScopeIds
	s.Condition = decoded.Condition
	s.Description = decoded.Description
	s.DirectoryScopeIds = decoded.DirectoryScopeIds
	s.DirectoryScopes_ODataBind = decoded.DirectoryScopes_ODataBind
	s.DisplayName = decoded.DisplayName
	s.PrincipalIds = decoded.PrincipalIds
	s.Principals_ODataBind = decoded.Principals_ODataBind
	s.RoleDefinition = decoded.RoleDefinition
	s.RoleDefinitionId = decoded.RoleDefinitionId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleAssignmentMultiple into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appScopes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppScopes into list []json.RawMessage: %+v", err)
		}

		output := make([]AppScope, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppScopeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppScopes' for 'UnifiedRoleAssignmentMultiple': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppScopes = &output
	}

	if v, ok := temp["directoryScopes"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DirectoryScopes into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DirectoryScopes' for 'UnifiedRoleAssignmentMultiple': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DirectoryScopes = &output
	}

	if v, ok := temp["principals"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Principals into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Principals' for 'UnifiedRoleAssignmentMultiple': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Principals = &output
	}

	return nil
}
