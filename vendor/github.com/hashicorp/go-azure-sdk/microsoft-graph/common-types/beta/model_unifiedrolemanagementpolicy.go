package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRoleManagementPolicy{}

type UnifiedRoleManagementPolicy struct {
	// Description for the policy.
	Description *string `json:"description,omitempty"`

	// Display name for the policy.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of effective rules like approval rules and expiration rules evaluated based on inherited referenced rules.
	// For example, if there is a tenant-wide policy to enforce enabling an approval rule, the effective rule will be to
	// enable approval even if the policy has a rule to disable approval. Supports $expand.
	EffectiveRules *[]UnifiedRoleManagementPolicyRule `json:"effectiveRules,omitempty"`

	// This can only be set to true for a single tenant-wide policy which will apply to all scopes and roles. Set the
	// scopeId to / and scopeType to Directory. Supports $filter (eq, ne).
	IsOrganizationDefault nullable.Type[bool] `json:"isOrganizationDefault,omitempty"`

	// The identity who last modified the role setting.
	LastModifiedBy Identity `json:"lastModifiedBy"`

	// The time when the role setting was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The collection of rules like approval rules and expiration rules. Supports $expand.
	Rules *[]UnifiedRoleManagementPolicyRule `json:"rules,omitempty"`

	// The identifier of the scope where the policy is created. Can be / for the tenant or a group ID. Required.
	ScopeId string `json:"scopeId"`

	// The type of the scope where the policy is created. One of Directory, DirectoryRole, Group. Required.
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

func (s UnifiedRoleManagementPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRoleManagementPolicy{}

func (s UnifiedRoleManagementPolicy) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRoleManagementPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRoleManagementPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRoleManagementPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRoleManagementPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRoleManagementPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UnifiedRoleManagementPolicy{}

func (s *UnifiedRoleManagementPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description           *string               `json:"description,omitempty"`
		DisplayName           *string               `json:"displayName,omitempty"`
		IsOrganizationDefault nullable.Type[bool]   `json:"isOrganizationDefault,omitempty"`
		LastModifiedDateTime  nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ScopeId               string                `json:"scopeId"`
		ScopeType             string                `json:"scopeType"`
		Id                    *string               `json:"id,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.IsOrganizationDefault = decoded.IsOrganizationDefault
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ScopeId = decoded.ScopeId
	s.ScopeType = decoded.ScopeType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleManagementPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["effectiveRules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EffectiveRules into list []json.RawMessage: %+v", err)
		}

		output := make([]UnifiedRoleManagementPolicyRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUnifiedRoleManagementPolicyRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EffectiveRules' for 'UnifiedRoleManagementPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EffectiveRules = &output
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'UnifiedRoleManagementPolicy': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["rules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Rules into list []json.RawMessage: %+v", err)
		}

		output := make([]UnifiedRoleManagementPolicyRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUnifiedRoleManagementPolicyRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Rules' for 'UnifiedRoleManagementPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Rules = &output
	}

	return nil
}
