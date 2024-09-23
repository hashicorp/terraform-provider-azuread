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
