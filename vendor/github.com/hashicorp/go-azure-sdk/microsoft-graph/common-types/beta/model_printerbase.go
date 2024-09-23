package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterBase interface {
	Entity
	PrinterBase() BasePrinterBaseImpl
}

var _ PrinterBase = BasePrinterBaseImpl{}

type BasePrinterBaseImpl struct {
	Capabilities    *PrinterCapabilities  `json:"capabilities,omitempty"`
	Defaults        *PrinterDefaults      `json:"defaults,omitempty"`
	DisplayName     *string               `json:"displayName,omitempty"`
	IsAcceptingJobs nullable.Type[bool]   `json:"isAcceptingJobs,omitempty"`
	Jobs            *[]PrintJob           `json:"jobs,omitempty"`
	Location        *PrinterLocation      `json:"location,omitempty"`
	Manufacturer    nullable.Type[string] `json:"manufacturer,omitempty"`
	Model           nullable.Type[string] `json:"model,omitempty"`
	Name            nullable.Type[string] `json:"name,omitempty"`
	Status          *PrinterStatus        `json:"status,omitempty"`

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

func (s BasePrinterBaseImpl) PrinterBase() BasePrinterBaseImpl {
	return s
}

func (s BasePrinterBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrinterBase = RawPrinterBaseImpl{}

// RawPrinterBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrinterBaseImpl struct {
	printerBase BasePrinterBaseImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawPrinterBaseImpl) PrinterBase() BasePrinterBaseImpl {
	return s.printerBase
}

func (s RawPrinterBaseImpl) Entity() BaseEntityImpl {
	return s.printerBase.Entity()
}

var _ json.Marshaler = BasePrinterBaseImpl{}

func (s BasePrinterBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrinterBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrinterBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrinterBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printerBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrinterBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPrinterBaseImplementation(input []byte) (PrinterBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrinterBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.printer") {
		var out Printer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Printer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printerShare") {
		var out PrinterShare
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrinterShare: %+v", err)
		}
		return out, nil
	}

	var parent BasePrinterBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrinterBaseImpl: %+v", err)
	}

	return RawPrinterBaseImpl{
		printerBase: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
