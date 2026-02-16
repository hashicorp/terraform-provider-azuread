package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecuritySearch = SecurityEdiscoveryReviewSetQuery{}

type SecurityEdiscoveryReviewSetQuery struct {

	// Fields inherited from SecuritySearch

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

func (s SecurityEdiscoveryReviewSetQuery) SecuritySearch() BaseSecuritySearchImpl {
	return BaseSecuritySearchImpl{
		ContentQuery:         s.ContentQuery,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryReviewSetQuery) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryReviewSetQuery{}

func (s SecurityEdiscoveryReviewSetQuery) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryReviewSetQuery
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryReviewSetQuery: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryReviewSetQuery: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryReviewSetQuery"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryReviewSetQuery: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryReviewSetQuery{}

func (s *SecurityEdiscoveryReviewSetQuery) UnmarshalJSON(bytes []byte) error {
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
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryReviewSetQuery into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryReviewSetQuery': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityEdiscoveryReviewSetQuery': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
