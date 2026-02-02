package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityDataSet = SecurityEdiscoveryReviewSet{}

type SecurityEdiscoveryReviewSet struct {
	// Represents queries within the review set.
	Queries *[]SecurityEdiscoveryReviewSetQuery `json:"queries,omitempty"`

	// Fields inherited from SecurityDataSet

	// The user who created the data set. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The date and time when the review set was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of the data set.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the data set. The name is unique with a maximum limit of 64 characters.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

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

func (s SecurityEdiscoveryReviewSet) SecurityDataSet() BaseSecurityDataSetImpl {
	return BaseSecurityDataSetImpl{
		CreatedBy:       s.CreatedBy,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecurityEdiscoveryReviewSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryReviewSet{}

func (s SecurityEdiscoveryReviewSet) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryReviewSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryReviewSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryReviewSet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryReviewSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryReviewSet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryReviewSet{}

func (s *SecurityEdiscoveryReviewSet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Queries         *[]SecurityEdiscoveryReviewSetQuery `json:"queries,omitempty"`
		CreatedDateTime nullable.Type[string]               `json:"createdDateTime,omitempty"`
		Description     nullable.Type[string]               `json:"description,omitempty"`
		DisplayName     nullable.Type[string]               `json:"displayName,omitempty"`
		Id              *string                             `json:"id,omitempty"`
		ODataId         *string                             `json:"@odata.id,omitempty"`
		ODataType       *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Queries = decoded.Queries
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryReviewSet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoveryReviewSet': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}
