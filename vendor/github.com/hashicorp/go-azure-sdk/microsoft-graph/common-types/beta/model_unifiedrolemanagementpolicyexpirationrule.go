package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementPolicyRule = UnifiedRoleManagementPolicyExpirationRule{}

type UnifiedRoleManagementPolicyExpirationRule struct {
	// Indicates whether expiration is required or if it's a permanently active assignment or eligibility.
	IsExpirationRequired nullable.Type[bool] `json:"isExpirationRequired,omitempty"`

	// The maximum duration allowed for eligibility or assignment that isn't permanent. Required when isExpirationRequired
	// is true.
	MaximumDuration nullable.Type[string] `json:"maximumDuration,omitempty"`

	// Fields inherited from UnifiedRoleManagementPolicyRule

	// Not implemented. Defines details of scope that's targeted by role management policy rule. The details can include the
	// principal type, the role assignment type, and actions affecting a role. Supports $filter (eq, ne).
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

func (s UnifiedRoleManagementPolicyExpirationRule) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return BaseUnifiedRoleManagementPolicyRuleImpl{
		Target:    s.Target,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s UnifiedRoleManagementPolicyExpirationRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicyExpirationRule{}

func (s UnifiedRoleManagementPolicyExpirationRule) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicyExpirationRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicyExpirationRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyExpirationRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyExpirationRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicyExpirationRule: %+v", err)
	}

	return encoded, nil
}
