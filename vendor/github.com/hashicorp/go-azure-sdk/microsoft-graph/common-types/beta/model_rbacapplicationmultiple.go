package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RbacApplicationMultiple{}

type RbacApplicationMultiple struct {
	ResourceNamespaces *[]UnifiedRbacResourceNamespace  `json:"resourceNamespaces,omitempty"`
	RoleAssignments    *[]UnifiedRoleAssignmentMultiple `json:"roleAssignments,omitempty"`
	RoleDefinitions    *[]UnifiedRoleDefinition         `json:"roleDefinitions,omitempty"`

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

func (s RbacApplicationMultiple) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RbacApplicationMultiple{}

func (s RbacApplicationMultiple) MarshalJSON() ([]byte, error) {
	type wrapper RbacApplicationMultiple
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RbacApplicationMultiple: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RbacApplicationMultiple: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.rbacApplicationMultiple"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RbacApplicationMultiple: %+v", err)
	}

	return encoded, nil
}
