package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRbacApplication{}

type UnifiedRbacApplication struct {
	// Workload-specific scope object that represents the resources for which the principal has been granted access.
	CustomAppScopes *[]CustomAppScope `json:"customAppScopes,omitempty"`

	// Resource that represents a collection of related actions.
	ResourceNamespaces *[]UnifiedRbacResourceNamespace `json:"resourceNamespaces,omitempty"`

	// Resource to grant access to users or groups.
	RoleAssignments *[]UnifiedRoleAssignment `json:"roleAssignments,omitempty"`

	// The roles allowed by RBAC providers and the permissions assigned to the roles.
	RoleDefinitions *[]UnifiedRoleDefinition `json:"roleDefinitions,omitempty"`

	// Resource to grant access to users or groups that are transitive.
	TransitiveRoleAssignments *[]UnifiedRoleAssignment `json:"transitiveRoleAssignments,omitempty"`

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

func (s UnifiedRbacApplication) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRbacApplication{}

func (s UnifiedRbacApplication) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRbacApplication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRbacApplication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRbacApplication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRbacApplication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRbacApplication: %+v", err)
	}

	return encoded, nil
}
