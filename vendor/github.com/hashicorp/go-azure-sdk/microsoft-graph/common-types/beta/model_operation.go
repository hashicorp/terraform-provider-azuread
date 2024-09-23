package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Operation interface {
	Entity
	Operation() BaseOperationImpl
}

var _ Operation = BaseOperationImpl{}

type BaseOperationImpl struct {
	// The start time of the operation.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The time of the last action of the operation.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// Possible values are: notStarted, running, completed, failed. Read-only.
	Status *OperationStatus `json:"status,omitempty"`

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

func (s BaseOperationImpl) Operation() BaseOperationImpl {
	return s
}

func (s BaseOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Operation = RawOperationImpl{}

// RawOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOperationImpl struct {
	operation BaseOperationImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawOperationImpl) Operation() BaseOperationImpl {
	return s.operation
}

func (s RawOperationImpl) Entity() BaseEntityImpl {
	return s.operation.Entity()
}

var _ json.Marshaler = BaseOperationImpl{}

func (s BaseOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOperationImpl: %+v", err)
	}

	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.operation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOperationImplementation(input []byte) (Operation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Operation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onenoteOperation") {
		var out OnenoteOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnenoteOperation: %+v", err)
		}
		return out, nil
	}

	var parent BaseOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOperationImpl: %+v", err)
	}

	return RawOperationImpl{
		operation: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
