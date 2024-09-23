package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementPolicyRule = UnifiedRoleManagementPolicyAuthenticationContextRule{}

type UnifiedRoleManagementPolicyAuthenticationContextRule struct {
	// The value of the authentication context claim.
	ClaimValue nullable.Type[string] `json:"claimValue,omitempty"`

	// Determines whether this rule is enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// Fields inherited from UnifiedRoleManagementPolicyRule

	// Defines details of scope that's targeted by role management policy rule. The details can include the principal type,
	// the role assignment type, and actions affecting a role. Supports $filter (eq, ne).
	Target *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`

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

func (s UnifiedRoleManagementPolicyAuthenticationContextRule) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return BaseUnifiedRoleManagementPolicyRuleImpl{
		Target:    s.Target,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s UnifiedRoleManagementPolicyAuthenticationContextRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicyAuthenticationContextRule{}

func (s UnifiedRoleManagementPolicyAuthenticationContextRule) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicyAuthenticationContextRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicyAuthenticationContextRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyAuthenticationContextRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyAuthenticationContextRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicyAuthenticationContextRule: %+v", err)
	}

	return encoded, nil
}
