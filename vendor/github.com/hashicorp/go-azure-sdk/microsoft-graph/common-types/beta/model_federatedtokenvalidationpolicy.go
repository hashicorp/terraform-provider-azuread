package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = FederatedTokenValidationPolicy{}

type FederatedTokenValidationPolicy struct {
	ValidatingDomains ValidatingDomains `json:"validatingDomains"`

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

func (s FederatedTokenValidationPolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s FederatedTokenValidationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FederatedTokenValidationPolicy{}

func (s FederatedTokenValidationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper FederatedTokenValidationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FederatedTokenValidationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FederatedTokenValidationPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.federatedTokenValidationPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FederatedTokenValidationPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &FederatedTokenValidationPolicy{}

func (s *FederatedTokenValidationPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FederatedTokenValidationPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["validatingDomains"]; ok {
		impl, err := UnmarshalValidatingDomainsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ValidatingDomains' for 'FederatedTokenValidationPolicy': %+v", err)
		}
		s.ValidatingDomains = impl
	}

	return nil
}
