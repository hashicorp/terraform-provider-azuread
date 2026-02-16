package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrintOperation = PrinterCreateOperation{}

type PrinterCreateOperation struct {
	// The signed certificate created during the registration process. Read-only.
	Certificate nullable.Type[string] `json:"certificate,omitempty"`

	// The created printer entity. Read-only.
	Printer *Printer `json:"printer,omitempty"`

	// Fields inherited from PrintOperation

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

func (s PrinterCreateOperation) PrintOperation() BasePrintOperationImpl {
	return BasePrintOperationImpl{
		CreatedDateTime: s.CreatedDateTime,
		Status:          s.Status,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PrinterCreateOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrinterCreateOperation{}

func (s PrinterCreateOperation) MarshalJSON() ([]byte, error) {
	type wrapper PrinterCreateOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrinterCreateOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrinterCreateOperation: %+v", err)
	}

	delete(decoded, "certificate")
	delete(decoded, "printer")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printerCreateOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrinterCreateOperation: %+v", err)
	}

	return encoded, nil
}
