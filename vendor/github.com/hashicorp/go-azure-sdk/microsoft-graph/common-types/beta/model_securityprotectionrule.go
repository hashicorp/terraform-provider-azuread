package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProtectionRule interface {
	Entity
	SecurityProtectionRule() BaseSecurityProtectionRuleImpl
}

var _ SecurityProtectionRule = BaseSecurityProtectionRuleImpl{}

type BaseSecurityProtectionRuleImpl struct {
	// Name of the user or application that created the rule.
	CreatedBy *string `json:"createdBy,omitempty"`

	// Timestamp of rule creation.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Name of the rule.
	DisplayName *string `json:"displayName,omitempty"`

	// Whether rule is turned on for the tenant.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Name of the user or application who last updated the rule.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// Timestamp of when the rule was last updated.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseSecurityProtectionRuleImpl) SecurityProtectionRule() BaseSecurityProtectionRuleImpl {
	return s
}

func (s BaseSecurityProtectionRuleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityProtectionRule = RawSecurityProtectionRuleImpl{}

// RawSecurityProtectionRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityProtectionRuleImpl struct {
	securityProtectionRule BaseSecurityProtectionRuleImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawSecurityProtectionRuleImpl) SecurityProtectionRule() BaseSecurityProtectionRuleImpl {
	return s.securityProtectionRule
}

func (s RawSecurityProtectionRuleImpl) Entity() BaseEntityImpl {
	return s.securityProtectionRule.Entity()
}

var _ json.Marshaler = BaseSecurityProtectionRuleImpl{}

func (s BaseSecurityProtectionRuleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityProtectionRuleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityProtectionRuleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityProtectionRuleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.protectionRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityProtectionRuleImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityProtectionRuleImplementation(input []byte) (SecurityProtectionRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityProtectionRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.detectionRule") {
		var out SecurityDetectionRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDetectionRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityProtectionRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityProtectionRuleImpl: %+v", err)
	}

	return RawSecurityProtectionRuleImpl{
		securityProtectionRule: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
