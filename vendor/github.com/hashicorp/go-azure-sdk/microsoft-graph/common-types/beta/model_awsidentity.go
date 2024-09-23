package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsIdentity interface {
	Entity
	AuthorizationSystemIdentity
	AwsIdentity() BaseAwsIdentityImpl
}

var _ AwsIdentity = BaseAwsIdentityImpl{}

type BaseAwsIdentityImpl struct {

	// Fields inherited from AuthorizationSystemIdentity

	// Navigation to the authorizationSystem object
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`

	// The name of the identity. Read-only. Supports $filter and (eq,contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Unique ID of the identity within the external system. Read-only.
	ExternalId *string `json:"externalId,omitempty"`

	// Represents details of the source of the identity.
	Source AuthorizationSystemIdentitySource `json:"source"`

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

func (s BaseAwsIdentityImpl) AwsIdentity() BaseAwsIdentityImpl {
	return s
}

func (s BaseAwsIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return BaseAuthorizationSystemIdentityImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		Source:              s.Source,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s BaseAwsIdentityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AwsIdentity = RawAwsIdentityImpl{}

// RawAwsIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAwsIdentityImpl struct {
	awsIdentity BaseAwsIdentityImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawAwsIdentityImpl) AwsIdentity() BaseAwsIdentityImpl {
	return s.awsIdentity
}

func (s RawAwsIdentityImpl) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
	return s.awsIdentity.AuthorizationSystemIdentity()
}

func (s RawAwsIdentityImpl) Entity() BaseEntityImpl {
	return s.awsIdentity.Entity()
}

var _ json.Marshaler = BaseAwsIdentityImpl{}

func (s BaseAwsIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAwsIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAwsIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAwsIdentityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAwsIdentityImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAwsIdentityImpl{}

func (s *BaseAwsIdentityImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		ExternalId  *string               `json:"externalId,omitempty"`
		Id          *string               `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAwsIdentityImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'BaseAwsIdentityImpl': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'BaseAwsIdentityImpl': %+v", err)
		}
		s.Source = impl
	}

	return nil
}

func UnmarshalAwsIdentityImplementation(input []byte) (AwsIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAccessKey") {
		var out AwsAccessKey
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAccessKey: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsEc2Instance") {
		var out AwsEc2Instance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsEc2Instance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsGroup") {
		var out AwsGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsLambda") {
		var out AwsLambda
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsLambda: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsRole") {
		var out AwsRole
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsRole: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsUser") {
		var out AwsUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsUser: %+v", err)
		}
		return out, nil
	}

	var parent BaseAwsIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAwsIdentityImpl: %+v", err)
	}

	return RawAwsIdentityImpl{
		awsIdentity: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
