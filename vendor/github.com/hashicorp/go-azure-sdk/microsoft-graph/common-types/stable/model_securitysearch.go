package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySearch interface {
	Entity
	SecuritySearch() BaseSecuritySearchImpl
}

var _ SecuritySearch = BaseSecuritySearchImpl{}

type BaseSecuritySearchImpl struct {
	ContentQuery         nullable.Type[string] `json:"contentQuery,omitempty"`
	CreatedBy            IdentitySet           `json:"createdBy"`
	CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
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

func (s BaseSecuritySearchImpl) SecuritySearch() BaseSecuritySearchImpl {
	return s
}

func (s BaseSecuritySearchImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecuritySearch = RawSecuritySearchImpl{}

// RawSecuritySearchImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecuritySearchImpl struct {
	securitySearch BaseSecuritySearchImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawSecuritySearchImpl) SecuritySearch() BaseSecuritySearchImpl {
	return s.securitySearch
}

func (s RawSecuritySearchImpl) Entity() BaseEntityImpl {
	return s.securitySearch.Entity()
}

var _ json.Marshaler = BaseSecuritySearchImpl{}

func (s BaseSecuritySearchImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecuritySearchImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecuritySearchImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecuritySearchImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.search"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecuritySearchImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecuritySearchImpl{}

func (s *BaseSecuritySearchImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ContentQuery         nullable.Type[string] `json:"contentQuery,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
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

	s.ContentQuery = decoded.ContentQuery
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseSecuritySearchImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecuritySearchImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseSecuritySearchImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalSecuritySearchImplementation(input []byte) (SecuritySearch, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySearch into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryReviewSetQuery") {
		var out SecurityEdiscoveryReviewSetQuery
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryReviewSetQuery: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoverySearch") {
		var out SecurityEdiscoverySearch
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoverySearch: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecuritySearchImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecuritySearchImpl: %+v", err)
	}

	return RawSecuritySearchImpl{
		securitySearch: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
