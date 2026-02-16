package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementPolicyRule = UnifiedRoleManagementPolicyEnablementRule{}

type UnifiedRoleManagementPolicyEnablementRule struct {
	// The collection of rules that are enabled for this policy rule. For example, MultiFactorAuthentication, Ticketing, and
	// Justification.
	EnabledRules *[]string `json:"enabledRules,omitempty"`

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

func (s UnifiedRoleManagementPolicyEnablementRule) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return BaseUnifiedRoleManagementPolicyRuleImpl{
		Target:    s.Target,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s UnifiedRoleManagementPolicyEnablementRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicyEnablementRule{}

func (s UnifiedRoleManagementPolicyEnablementRule) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicyEnablementRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicyEnablementRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyEnablementRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyEnablementRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicyEnablementRule: %+v", err)
	}

	return encoded, nil
}
