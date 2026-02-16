package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoveryReviewSet{}

type EdiscoveryReviewSet struct {
	// The user who created the review set. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The datetime when the review set was created. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The review set name. The name is unique with a maximum limit of 64 characters.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Queries *[]EdiscoveryReviewSetQuery `json:"queries,omitempty"`

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

func (s EdiscoveryReviewSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryReviewSet{}

func (s EdiscoveryReviewSet) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryReviewSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryReviewSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryReviewSet: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.reviewSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryReviewSet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryReviewSet{}

func (s *EdiscoveryReviewSet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string]       `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]       `json:"displayName,omitempty"`
		Queries         *[]EdiscoveryReviewSetQuery `json:"queries,omitempty"`
		Id              *string                     `json:"id,omitempty"`
		ODataId         *string                     `json:"@odata.id,omitempty"`
		ODataType       *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Queries = decoded.Queries
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryReviewSet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoveryReviewSet': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}
