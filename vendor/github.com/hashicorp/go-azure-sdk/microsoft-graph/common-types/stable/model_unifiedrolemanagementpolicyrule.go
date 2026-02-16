package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyRule interface {
	Entity
	UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl
}

var _ UnifiedRoleManagementPolicyRule = BaseUnifiedRoleManagementPolicyRuleImpl{}

type BaseUnifiedRoleManagementPolicyRuleImpl struct {
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

func (s BaseUnifiedRoleManagementPolicyRuleImpl) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return s
}

func (s BaseUnifiedRoleManagementPolicyRuleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ UnifiedRoleManagementPolicyRule = RawUnifiedRoleManagementPolicyRuleImpl{}

// RawUnifiedRoleManagementPolicyRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUnifiedRoleManagementPolicyRuleImpl struct {
	unifiedRoleManagementPolicyRule BaseUnifiedRoleManagementPolicyRuleImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawUnifiedRoleManagementPolicyRuleImpl) UnifiedRoleManagementPolicyRule() BaseUnifiedRoleManagementPolicyRuleImpl {
	return s.unifiedRoleManagementPolicyRule
}

func (s RawUnifiedRoleManagementPolicyRuleImpl) Entity() BaseEntityImpl {
	return s.unifiedRoleManagementPolicyRule.Entity()
}

var _ json.Marshaler = BaseUnifiedRoleManagementPolicyRuleImpl{}

func (s BaseUnifiedRoleManagementPolicyRuleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseUnifiedRoleManagementPolicyRuleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseUnifiedRoleManagementPolicyRuleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseUnifiedRoleManagementPolicyRuleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicyRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseUnifiedRoleManagementPolicyRuleImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalUnifiedRoleManagementPolicyRuleImplementation(input []byte) (UnifiedRoleManagementPolicyRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyApprovalRule") {
		var out UnifiedRoleManagementPolicyApprovalRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyApprovalRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyAuthenticationContextRule") {
		var out UnifiedRoleManagementPolicyAuthenticationContextRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyAuthenticationContextRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyEnablementRule") {
		var out UnifiedRoleManagementPolicyEnablementRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyEnablementRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyExpirationRule") {
		var out UnifiedRoleManagementPolicyExpirationRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyExpirationRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unifiedRoleManagementPolicyNotificationRule") {
		var out UnifiedRoleManagementPolicyNotificationRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnifiedRoleManagementPolicyNotificationRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseUnifiedRoleManagementPolicyRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUnifiedRoleManagementPolicyRuleImpl: %+v", err)
	}

	return RawUnifiedRoleManagementPolicyRuleImpl{
		unifiedRoleManagementPolicyRule: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
