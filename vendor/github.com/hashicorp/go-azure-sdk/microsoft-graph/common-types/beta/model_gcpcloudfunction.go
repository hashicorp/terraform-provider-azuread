package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GcpIdentity = GcpCloudFunction{}

type GcpCloudFunction struct {
	// Represents the resources in an authorization system..
	Resource *GcpAuthorizationSystemResource `json:"resource,omitempty"`

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

func (s GcpCloudFunction) GcpIdentity() BaseGcpIdentityImpl {
	return BaseGcpIdentityImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		Source:              s.Source,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s GcpCloudFunction) AuthorizationSystemIdentity() BaseAuthorizationSystemIdentityImpl {
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

func (s GcpCloudFunction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpCloudFunction{}

func (s GcpCloudFunction) MarshalJSON() ([]byte, error) {
	type wrapper GcpCloudFunction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpCloudFunction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpCloudFunction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpCloudFunction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpCloudFunction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GcpCloudFunction{}

func (s *GcpCloudFunction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Resource    *GcpAuthorizationSystemResource `json:"resource,omitempty"`
		DisplayName nullable.Type[string]           `json:"displayName,omitempty"`
		ExternalId  *string                         `json:"externalId,omitempty"`
		Id          *string                         `json:"id,omitempty"`
		ODataId     *string                         `json:"@odata.id,omitempty"`
		ODataType   *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Resource = decoded.Resource
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GcpCloudFunction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'GcpCloudFunction': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentitySourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'GcpCloudFunction': %+v", err)
		}
		s.Source = impl
	}

	return nil
}
