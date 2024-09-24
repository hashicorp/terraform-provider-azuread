package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedMobileApp{}

type ManagedMobileApp struct {
	// The identifier for an app with it's operating system type.
	MobileAppIdentifier MobileAppIdentifier `json:"mobileAppIdentifier"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s ManagedMobileApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedMobileApp{}

func (s ManagedMobileApp) MarshalJSON() ([]byte, error) {
	type wrapper ManagedMobileApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedMobileApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedMobileApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedMobileApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedMobileApp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ManagedMobileApp{}

func (s *ManagedMobileApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Version   nullable.Type[string] `json:"version,omitempty"`
		Id        *string               `json:"id,omitempty"`
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ManagedMobileApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["mobileAppIdentifier"]; ok {
		impl, err := UnmarshalMobileAppIdentifierImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MobileAppIdentifier' for 'ManagedMobileApp': %+v", err)
		}
		s.MobileAppIdentifier = impl
	}

	return nil
}
