package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyBase interface {
	Entity
	DirectoryObject
	PolicyBase() BasePolicyBaseImpl
}

var _ PolicyBase = BasePolicyBaseImpl{}

type BasePolicyBaseImpl struct {
	// Description for this policy. Required.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for this policy. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s BasePolicyBaseImpl) PolicyBase() BasePolicyBaseImpl {
	return s
}

func (s BasePolicyBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BasePolicyBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PolicyBase = RawPolicyBaseImpl{}

// RawPolicyBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPolicyBaseImpl struct {
	policyBase BasePolicyBaseImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawPolicyBaseImpl) PolicyBase() BasePolicyBaseImpl {
	return s.policyBase
}

func (s RawPolicyBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.policyBase.DirectoryObject()
}

func (s RawPolicyBaseImpl) Entity() BaseEntityImpl {
	return s.policyBase.Entity()
}

var _ json.Marshaler = BasePolicyBaseImpl{}

func (s BasePolicyBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePolicyBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePolicyBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePolicyBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policyBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePolicyBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPolicyBaseImplementation(input []byte) (PolicyBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appManagementPolicy") {
		var out AppManagementPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppManagementPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authorizationPolicy") {
		var out AuthorizationPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthorizationPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicy") {
		var out CrossTenantAccessPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identitySecurityDefaultsEnforcementPolicy") {
		var out IdentitySecurityDefaultsEnforcementPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentitySecurityDefaultsEnforcementPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.permissionGrantPolicy") {
		var out PermissionGrantPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionGrantPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.stsPolicy") {
		var out StsPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StsPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tenantAppManagementPolicy") {
		var out TenantAppManagementPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TenantAppManagementPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BasePolicyBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePolicyBaseImpl: %+v", err)
	}

	return RawPolicyBaseImpl{
		policyBase: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
