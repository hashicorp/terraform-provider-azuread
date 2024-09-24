package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystemResource = GcpAuthorizationSystemResource{}

type GcpAuthorizationSystemResource struct {
	// The service associated with the resource in a GCP authorization system. This object is autoexpanded.
	Service *AuthorizationSystemTypeService `json:"service,omitempty"`

	// Fields inherited from AuthorizationSystemResource

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

func (s GcpAuthorizationSystemResource) AuthorizationSystemResource() BaseAuthorizationSystemResourceImpl {
	return BaseAuthorizationSystemResourceImpl{
		AuthorizationSystem: s.AuthorizationSystem,
		DisplayName:         s.DisplayName,
		ExternalId:          s.ExternalId,
		ResourceType:        s.ResourceType,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s GcpAuthorizationSystemResource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpAuthorizationSystemResource{}

func (s GcpAuthorizationSystemResource) MarshalJSON() ([]byte, error) {
	type wrapper GcpAuthorizationSystemResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpAuthorizationSystemResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpAuthorizationSystemResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpAuthorizationSystemResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpAuthorizationSystemResource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GcpAuthorizationSystemResource{}

func (s *GcpAuthorizationSystemResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Service      *AuthorizationSystemTypeService `json:"service,omitempty"`
		DisplayName  nullable.Type[string]           `json:"displayName,omitempty"`
		ExternalId   *string                         `json:"externalId,omitempty"`
		ResourceType nullable.Type[string]           `json:"resourceType,omitempty"`
		Id           *string                         `json:"id,omitempty"`
		ODataId      *string                         `json:"@odata.id,omitempty"`
		ODataType    *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Service = decoded.Service
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceType = decoded.ResourceType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GcpAuthorizationSystemResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'GcpAuthorizationSystemResource': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	return nil
}
