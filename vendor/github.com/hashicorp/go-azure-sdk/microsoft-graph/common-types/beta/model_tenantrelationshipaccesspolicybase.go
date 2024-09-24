package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantRelationshipAccessPolicyBase interface {
	Entity
	DirectoryObject
	PolicyBase
	TenantRelationshipAccessPolicyBase() BaseTenantRelationshipAccessPolicyBaseImpl
}

var _ TenantRelationshipAccessPolicyBase = BaseTenantRelationshipAccessPolicyBaseImpl{}

type BaseTenantRelationshipAccessPolicyBaseImpl struct {
	// The raw JSON definition of the cross-tenant access policy. Deprecated. Do not use.
	Definition *[]string `json:"definition,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description string `json:"description"`

	// Display name for this policy. Required.
	DisplayName string `json:"displayName"`

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

func (s BaseTenantRelationshipAccessPolicyBaseImpl) TenantRelationshipAccessPolicyBase() BaseTenantRelationshipAccessPolicyBaseImpl {
	return s
}

func (s BaseTenantRelationshipAccessPolicyBaseImpl) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseTenantRelationshipAccessPolicyBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseTenantRelationshipAccessPolicyBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TenantRelationshipAccessPolicyBase = RawTenantRelationshipAccessPolicyBaseImpl{}

// RawTenantRelationshipAccessPolicyBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTenantRelationshipAccessPolicyBaseImpl struct {
	tenantRelationshipAccessPolicyBase BaseTenantRelationshipAccessPolicyBaseImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawTenantRelationshipAccessPolicyBaseImpl) TenantRelationshipAccessPolicyBase() BaseTenantRelationshipAccessPolicyBaseImpl {
	return s.tenantRelationshipAccessPolicyBase
}

func (s RawTenantRelationshipAccessPolicyBaseImpl) PolicyBase() BasePolicyBaseImpl {
	return s.tenantRelationshipAccessPolicyBase.PolicyBase()
}

func (s RawTenantRelationshipAccessPolicyBaseImpl) DirectoryObject() BaseDirectoryObjectImpl {
	return s.tenantRelationshipAccessPolicyBase.DirectoryObject()
}

func (s RawTenantRelationshipAccessPolicyBaseImpl) Entity() BaseEntityImpl {
	return s.tenantRelationshipAccessPolicyBase.Entity()
}

var _ json.Marshaler = BaseTenantRelationshipAccessPolicyBaseImpl{}

func (s BaseTenantRelationshipAccessPolicyBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTenantRelationshipAccessPolicyBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTenantRelationshipAccessPolicyBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTenantRelationshipAccessPolicyBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tenantRelationshipAccessPolicyBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTenantRelationshipAccessPolicyBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTenantRelationshipAccessPolicyBaseImplementation(input []byte) (TenantRelationshipAccessPolicyBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TenantRelationshipAccessPolicyBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.crossTenantAccessPolicy") {
		var out CrossTenantAccessPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossTenantAccessPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseTenantRelationshipAccessPolicyBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTenantRelationshipAccessPolicyBaseImpl: %+v", err)
	}

	return RawTenantRelationshipAccessPolicyBaseImpl{
		tenantRelationshipAccessPolicyBase: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
