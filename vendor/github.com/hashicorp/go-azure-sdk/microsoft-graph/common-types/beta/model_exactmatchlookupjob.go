package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExactMatchJobBase = ExactMatchLookupJob{}

type ExactMatchLookupJob struct {
	MatchingRows *[]LookupResultRow    `json:"matchingRows,omitempty"`
	State        nullable.Type[string] `json:"state,omitempty"`

	// Fields inherited from ExactMatchJobBase

	CompletionDateTime  nullable.Type[string] `json:"completionDateTime,omitempty"`
	CreationDateTime    nullable.Type[string] `json:"creationDateTime,omitempty"`
	Error               *ClassificationError  `json:"error,omitempty"`
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
	StartDateTime       nullable.Type[string] `json:"startDateTime,omitempty"`

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

func (s ExactMatchLookupJob) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return BaseExactMatchJobBaseImpl{
		CompletionDateTime:  s.CompletionDateTime,
		CreationDateTime:    s.CreationDateTime,
		Error:               s.Error,
		LastUpdatedDateTime: s.LastUpdatedDateTime,
		StartDateTime:       s.StartDateTime,
		Id:                  s.Id,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

func (s ExactMatchLookupJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExactMatchLookupJob{}

func (s ExactMatchLookupJob) MarshalJSON() ([]byte, error) {
	type wrapper ExactMatchLookupJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExactMatchLookupJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchLookupJob: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exactMatchLookupJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExactMatchLookupJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExactMatchLookupJob{}

func (s *ExactMatchLookupJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		MatchingRows        *[]LookupResultRow    `json:"matchingRows,omitempty"`
		State               nullable.Type[string] `json:"state,omitempty"`
		CompletionDateTime  nullable.Type[string] `json:"completionDateTime,omitempty"`
		CreationDateTime    nullable.Type[string] `json:"creationDateTime,omitempty"`
		LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
		StartDateTime       nullable.Type[string] `json:"startDateTime,omitempty"`
		Id                  *string               `json:"id,omitempty"`
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.MatchingRows = decoded.MatchingRows
	s.State = decoded.State
	s.CompletionDateTime = decoded.CompletionDateTime
	s.CreationDateTime = decoded.CreationDateTime
	s.Id = decoded.Id
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.StartDateTime = decoded.StartDateTime

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExactMatchLookupJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["error"]; ok {
		impl, err := UnmarshalClassificationErrorImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Error' for 'ExactMatchLookupJob': %+v", err)
		}
		s.Error = &impl
	}

	return nil
}
