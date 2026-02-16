package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperation interface {
	Entity
	PrintOperation() BasePrintOperationImpl
}

var _ PrintOperation = BasePrintOperationImpl{}

type BasePrintOperationImpl struct {
	// The DateTimeOffset when the operation was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	Status *PrintOperationStatus `json:"status,omitempty"`

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

func (s BasePrintOperationImpl) PrintOperation() BasePrintOperationImpl {
	return s
}

func (s BasePrintOperationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrintOperation = RawPrintOperationImpl{}

// RawPrintOperationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrintOperationImpl struct {
	printOperation BasePrintOperationImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawPrintOperationImpl) PrintOperation() BasePrintOperationImpl {
	return s.printOperation
}

func (s RawPrintOperationImpl) Entity() BaseEntityImpl {
	return s.printOperation.Entity()
}

var _ json.Marshaler = BasePrintOperationImpl{}

func (s BasePrintOperationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrintOperationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrintOperationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrintOperationImpl: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrintOperationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPrintOperationImplementation(input []byte) (PrintOperation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintOperation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.printerCreateOperation") {
		var out PrinterCreateOperation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterCreateOperation: %+v", err)
		}
		return out, nil
	}

	var parent BasePrintOperationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrintOperationImpl: %+v", err)
	}

	return RawPrintOperationImpl{
		printOperation: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
