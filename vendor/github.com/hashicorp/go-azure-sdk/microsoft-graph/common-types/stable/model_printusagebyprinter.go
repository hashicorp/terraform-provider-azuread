package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrintUsage = PrintUsageByPrinter{}

type PrintUsageByPrinter struct {
	// The ID of the printer represented by these statistics.
	PrinterId *string `json:"printerId,omitempty"`

	// The name of the printer represented by these statistics.
	PrinterName nullable.Type[string] `json:"printerName,omitempty"`

	// Fields inherited from PrintUsage

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

func (s PrintUsageByPrinter) PrintUsage() BasePrintUsageImpl {
	return BasePrintUsageImpl{
		BlackAndWhitePageCount:         s.BlackAndWhitePageCount,
		ColorPageCount:                 s.ColorPageCount,
		CompletedBlackAndWhiteJobCount: s.CompletedBlackAndWhiteJobCount,
		CompletedColorJobCount:         s.CompletedColorJobCount,
		CompletedJobCount:              s.CompletedJobCount,
		DoubleSidedSheetCount:          s.DoubleSidedSheetCount,
		IncompleteJobCount:             s.IncompleteJobCount,
		MediaSheetCount:                s.MediaSheetCount,
		PageCount:                      s.PageCount,
		SingleSidedSheetCount:          s.SingleSidedSheetCount,
		UsageDate:                      s.UsageDate,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s PrintUsageByPrinter) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrintUsageByPrinter{}

func (s PrintUsageByPrinter) MarshalJSON() ([]byte, error) {
	type wrapper PrintUsageByPrinter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintUsageByPrinter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintUsageByPrinter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printUsageByPrinter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintUsageByPrinter: %+v", err)
	}

	return encoded, nil
}
