package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchDataStoreBase interface {
	Entity
	ExactMatchDataStoreBase() BaseExactMatchDataStoreBaseImpl
}

var _ ExactMatchDataStoreBase = BaseExactMatchDataStoreBaseImpl{}

type BaseExactMatchDataStoreBaseImpl struct {
	Columns                 *[]ExactDataMatchStoreColumn `json:"columns,omitempty"`
	DataLastUpdatedDateTime nullable.Type[string]        `json:"dataLastUpdatedDateTime,omitempty"`
	Description             nullable.Type[string]        `json:"description,omitempty"`
	DisplayName             nullable.Type[string]        `json:"displayName,omitempty"`

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

func (s BaseExactMatchDataStoreBaseImpl) ExactMatchDataStoreBase() BaseExactMatchDataStoreBaseImpl {
	return s
}

func (s BaseExactMatchDataStoreBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ExactMatchDataStoreBase = RawExactMatchDataStoreBaseImpl{}

// RawExactMatchDataStoreBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExactMatchDataStoreBaseImpl struct {
	exactMatchDataStoreBase BaseExactMatchDataStoreBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawExactMatchDataStoreBaseImpl) ExactMatchDataStoreBase() BaseExactMatchDataStoreBaseImpl {
	return s.exactMatchDataStoreBase
}

func (s RawExactMatchDataStoreBaseImpl) Entity() BaseEntityImpl {
	return s.exactMatchDataStoreBase.Entity()
}

var _ json.Marshaler = BaseExactMatchDataStoreBaseImpl{}

func (s BaseExactMatchDataStoreBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExactMatchDataStoreBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExactMatchDataStoreBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExactMatchDataStoreBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exactMatchDataStoreBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExactMatchDataStoreBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExactMatchDataStoreBaseImplementation(input []byte) (ExactMatchDataStoreBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchDataStoreBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchDataStore") {
		var out ExactMatchDataStore
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchDataStore: %+v", err)
		}
		return out, nil
	}

	var parent BaseExactMatchDataStoreBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExactMatchDataStoreBaseImpl: %+v", err)
	}

	return RawExactMatchDataStoreBaseImpl{
		exactMatchDataStoreBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
