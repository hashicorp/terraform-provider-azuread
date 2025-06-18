package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ServiceApp{}

type ServiceApp struct {
	// The Entra ID application ID.
	Application Identity `json:"application"`

	// Timestamp of the effective activation of the service app.
	EffectiveDateTime nullable.Type[string] `json:"effectiveDateTime,omitempty"`

	// Identity of the person who last modified the entity.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification of the entity.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Timestamp of the creation of the service app entity.
	RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`

	// The status of the service app. This value indicates whether or not the application can be used to control the backup
	// service. The possible values are: inactive, active, pendingActive, pendingInactive, unknownFutureValue.
	Status *ServiceAppStatus `json:"status,omitempty"`

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

func (s ServiceApp) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceApp{}

func (s ServiceApp) MarshalJSON() ([]byte, error) {
	type wrapper ServiceApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceApp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ServiceApp{}

func (s *ServiceApp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EffectiveDateTime    nullable.Type[string] `json:"effectiveDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`
		Status               *ServiceAppStatus     `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EffectiveDateTime = decoded.EffectiveDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RegistrationDateTime = decoded.RegistrationDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceApp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'ServiceApp': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ServiceApp': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
