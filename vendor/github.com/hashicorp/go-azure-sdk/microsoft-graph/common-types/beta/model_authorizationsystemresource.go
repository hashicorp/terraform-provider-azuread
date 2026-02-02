package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemResource interface {
	Entity
	AuthorizationSystemResource() BaseAuthorizationSystemResourceImpl
}

var _ AuthorizationSystemResource = BaseAuthorizationSystemResourceImpl{}

type BaseAuthorizationSystemResourceImpl struct {
	// The authorization system that the resource exists in.
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`

	// The name of the resource. Read-only. Supports $filter (eq,contains).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The ID of the resource as defined by the authorization system provider. Read-only. Supports $filter (eq).
	ExternalId *string `json:"externalId,omitempty"`

	// The type of the resource. Read-only. Supports $filter (eq).
	ResourceType nullable.Type[string] `json:"resourceType,omitempty"`

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

func (s BaseAuthorizationSystemResourceImpl) AuthorizationSystemResource() BaseAuthorizationSystemResourceImpl {
	return s
}

func (s BaseAuthorizationSystemResourceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthorizationSystemResource = RawAuthorizationSystemResourceImpl{}

// RawAuthorizationSystemResourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthorizationSystemResourceImpl struct {
	authorizationSystemResource BaseAuthorizationSystemResourceImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawAuthorizationSystemResourceImpl) AuthorizationSystemResource() BaseAuthorizationSystemResourceImpl {
	return s.authorizationSystemResource
}

func (s RawAuthorizationSystemResourceImpl) Entity() BaseEntityImpl {
	return s.authorizationSystemResource.Entity()
}

var _ json.Marshaler = BaseAuthorizationSystemResourceImpl{}

func (s BaseAuthorizationSystemResourceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthorizationSystemResourceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthorizationSystemResourceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthorizationSystemResourceImpl: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "externalId")
	delete(decoded, "resourceType")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authorizationSystemResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthorizationSystemResourceImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAuthorizationSystemResourceImpl{}

func (s *BaseAuthorizationSystemResourceImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName  nullable.Type[string] `json:"displayName,omitempty"`
		ExternalId   *string               `json:"externalId,omitempty"`
		ResourceType nullable.Type[string] `json:"resourceType,omitempty"`
		Id           *string               `json:"id,omitempty"`
		ODataId      *string               `json:"@odata.id,omitempty"`
		ODataType    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.ResourceType = decoded.ResourceType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAuthorizationSystemResourceImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'BaseAuthorizationSystemResourceImpl': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	return nil
}

func UnmarshalAuthorizationSystemResourceImplementation(input []byte) (AuthorizationSystemResource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthorizationSystemResource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystemResource") {
		var out AwsAuthorizationSystemResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystemResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystemResource") {
		var out AzureAuthorizationSystemResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystemResource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystemResource") {
		var out GcpAuthorizationSystemResource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystemResource: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthorizationSystemResourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthorizationSystemResourceImpl: %+v", err)
	}

	return RawAuthorizationSystemResourceImpl{
		authorizationSystemResource: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
