package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleManagementPolicyAssignment{}

type UnifiedRoleManagementPolicyAssignment struct {
	// The policy that's associated with a policy assignment. Supports $expand and a nested $expand of the rules and
	// effectiveRules relationships for the policy.
	Policy *UnifiedRoleManagementPolicy `json:"policy,omitempty"`

	// The id of the policy. Inherited from entity.
	PolicyId *string `json:"policyId,omitempty"`

	// For Microsoft Entra roles policy, it's the identifier of the role definition object where the policy applies. For PIM
	// for groups membership and ownership, it's either member or owner. Supports $filter (eq).
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The identifier of the scope where the policy is assigned. Can be / for the tenant or a group ID. Required.
	ScopeId string `json:"scopeId"`

	// The type of the scope where the policy is assigned. One of Directory, DirectoryRole, Group. Required.
	ScopeType string `json:"scopeType"`

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

func (s UnifiedRoleManagementPolicyAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicyAssignment{}

func (s UnifiedRoleManagementPolicyAssignment) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicyAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicyAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicyAssignment: %+v", err)
	}

	return encoded, nil
}
