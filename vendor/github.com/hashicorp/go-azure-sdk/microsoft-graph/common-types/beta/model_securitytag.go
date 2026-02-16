package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTag interface {
	Entity
	SecurityTag() BaseSecurityTagImpl
}

var _ SecurityTag = BaseSecurityTagImpl{}

type BaseSecurityTagImpl struct {
	CreatedBy            IdentitySet           `json:"createdBy"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseSecurityTagImpl) SecurityTag() BaseSecurityTagImpl {
	return s
}

func (s BaseSecurityTagImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityTag = RawSecurityTagImpl{}

// RawSecurityTagImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityTagImpl struct {
	securityTag BaseSecurityTagImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawSecurityTagImpl) SecurityTag() BaseSecurityTagImpl {
	return s.securityTag
}

func (s RawSecurityTagImpl) Entity() BaseEntityImpl {
	return s.securityTag.Entity()
}

var _ json.Marshaler = BaseSecurityTagImpl{}

func (s BaseSecurityTagImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityTagImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityTagImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityTagImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.tag"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityTagImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityTagImpl{}

func (s *BaseSecurityTagImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description          nullable.Type[string] `json:"description,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecurityTagImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecurityTagImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalSecurityTagImplementation(input []byte) (SecurityTag, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityTag into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewTag") {
		var out SecurityEdiscoveryReviewTag
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewTag: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityTagImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityTagImpl: %+v", err)
	}

	return RawSecurityTagImpl{
		securityTag: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
