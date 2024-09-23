package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchSessionBase interface {
	Entity
	ExactMatchJobBase
	ExactMatchSessionBase() BaseExactMatchSessionBaseImpl
}

var _ ExactMatchSessionBase = BaseExactMatchSessionBaseImpl{}

type BaseExactMatchSessionBaseImpl struct {
	DataStoreId                  nullable.Type[string] `json:"dataStoreId,omitempty"`
	ProcessingCompletionDateTime nullable.Type[string] `json:"processingCompletionDateTime,omitempty"`
	RemainingBlockCount          nullable.Type[int64]  `json:"remainingBlockCount,omitempty"`
	RemainingJobCount            nullable.Type[int64]  `json:"remainingJobCount,omitempty"`
	State                        nullable.Type[string] `json:"state,omitempty"`
	TotalBlockCount              nullable.Type[int64]  `json:"totalBlockCount,omitempty"`
	TotalJobCount                nullable.Type[int64]  `json:"totalJobCount,omitempty"`
	UploadCompletionDateTime     nullable.Type[string] `json:"uploadCompletionDateTime,omitempty"`

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

func (s BaseExactMatchSessionBaseImpl) ExactMatchSessionBase() BaseExactMatchSessionBaseImpl {
	return s
}

func (s BaseExactMatchSessionBaseImpl) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
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

func (s BaseExactMatchSessionBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ExactMatchSessionBase = RawExactMatchSessionBaseImpl{}

// RawExactMatchSessionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExactMatchSessionBaseImpl struct {
	exactMatchSessionBase BaseExactMatchSessionBaseImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawExactMatchSessionBaseImpl) ExactMatchSessionBase() BaseExactMatchSessionBaseImpl {
	return s.exactMatchSessionBase
}

func (s RawExactMatchSessionBaseImpl) ExactMatchJobBase() BaseExactMatchJobBaseImpl {
	return s.exactMatchSessionBase.ExactMatchJobBase()
}

func (s RawExactMatchSessionBaseImpl) Entity() BaseEntityImpl {
	return s.exactMatchSessionBase.Entity()
}

var _ json.Marshaler = BaseExactMatchSessionBaseImpl{}

func (s BaseExactMatchSessionBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExactMatchSessionBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExactMatchSessionBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExactMatchSessionBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exactMatchSessionBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExactMatchSessionBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExactMatchSessionBaseImplementation(input []byte) (ExactMatchSessionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchSessionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchSession") {
		var out ExactMatchSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchSession: %+v", err)
		}
		return out, nil
	}

	var parent BaseExactMatchSessionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExactMatchSessionBaseImpl: %+v", err)
	}

	return RawExactMatchSessionBaseImpl{
		exactMatchSessionBase: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
