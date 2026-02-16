package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintUsage interface {
	Entity
	PrintUsage() BasePrintUsageImpl
}

var _ PrintUsage = BasePrintUsageImpl{}

type BasePrintUsageImpl struct {
	BlackAndWhitePageCount         nullable.Type[int64] `json:"blackAndWhitePageCount,omitempty"`
	ColorPageCount                 nullable.Type[int64] `json:"colorPageCount,omitempty"`
	CompletedBlackAndWhiteJobCount *int64               `json:"completedBlackAndWhiteJobCount,omitempty"`
	CompletedColorJobCount         *int64               `json:"completedColorJobCount,omitempty"`
	CompletedJobCount              nullable.Type[int64] `json:"completedJobCount,omitempty"`
	DoubleSidedSheetCount          nullable.Type[int64] `json:"doubleSidedSheetCount,omitempty"`
	IncompleteJobCount             *int64               `json:"incompleteJobCount,omitempty"`
	MediaSheetCount                nullable.Type[int64] `json:"mediaSheetCount,omitempty"`
	PageCount                      nullable.Type[int64] `json:"pageCount,omitempty"`
	SingleSidedSheetCount          nullable.Type[int64] `json:"singleSidedSheetCount,omitempty"`
	UsageDate                      *string              `json:"usageDate,omitempty"`

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

func (s BasePrintUsageImpl) PrintUsage() BasePrintUsageImpl {
	return s
}

func (s BasePrintUsageImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrintUsage = RawPrintUsageImpl{}

// RawPrintUsageImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrintUsageImpl struct {
	printUsage BasePrintUsageImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawPrintUsageImpl) PrintUsage() BasePrintUsageImpl {
	return s.printUsage
}

func (s RawPrintUsageImpl) Entity() BaseEntityImpl {
	return s.printUsage.Entity()
}

var _ json.Marshaler = BasePrintUsageImpl{}

func (s BasePrintUsageImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrintUsageImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrintUsageImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrintUsageImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printUsage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrintUsageImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPrintUsageImplementation(input []byte) (PrintUsage, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintUsage into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.printUsageByPrinter") {
		var out PrintUsageByPrinter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsageByPrinter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.printUsageByUser") {
		var out PrintUsageByUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrintUsageByUser: %+v", err)
		}
		return out, nil
	}

	var parent BasePrintUsageImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrintUsageImpl: %+v", err)
	}

	return RawPrintUsageImpl{
		printUsage: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
