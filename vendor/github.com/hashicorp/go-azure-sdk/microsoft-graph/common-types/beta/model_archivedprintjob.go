package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJob struct {
	// True if the job was acquired by a printer; false otherwise. Read-only.
	AcquiredByPrinter *bool `json:"acquiredByPrinter,omitempty"`

	// The dateTimeOffset when the job was acquired by the printer, if any. Read-only.
	AcquiredDateTime nullable.Type[string] `json:"acquiredDateTime,omitempty"`

	// The number of black and white pages that were printed. Read-only.
	BlackAndWhitePageCount *int64 `json:"blackAndWhitePageCount,omitempty"`

	// The number of color pages that were printed. Read-only.
	ColorPageCount *int64 `json:"colorPageCount,omitempty"`

	// The dateTimeOffset when the job was completed, canceled, or aborted. Read-only.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// The number of copies that were printed. Read-only.
	CopiesPrinted *int64 `json:"copiesPrinted,omitempty"`

	// The user who created the print job. Read-only.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// The dateTimeOffset when the job was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The number of duplex (double-sided) pages that were printed. Read-only.
	DuplexPageCount *int64 `json:"duplexPageCount,omitempty"`

	// The archived print job's GUID. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of pages that were printed. Read-only.
	PageCount *int64 `json:"pageCount,omitempty"`

	// The printer ID that the job was queued for. Read-only.
	PrinterId nullable.Type[string] `json:"printerId,omitempty"`

	// The printer name that the job was queued for. Read-only.
	PrinterName nullable.Type[string] `json:"printerName,omitempty"`

	ProcessingState *PrintJobProcessingState `json:"processingState,omitempty"`

	// The number of simplex (single-sided) pages that were printed. Read-only.
	SimplexPageCount *int64 `json:"simplexPageCount,omitempty"`
}

var _ json.Marshaler = ArchivedPrintJob{}

func (s ArchivedPrintJob) MarshalJSON() ([]byte, error) {
	type wrapper ArchivedPrintJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ArchivedPrintJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ArchivedPrintJob: %+v", err)
	}

	delete(decoded, "acquiredByPrinter")
	delete(decoded, "acquiredDateTime")
	delete(decoded, "blackAndWhitePageCount")
	delete(decoded, "colorPageCount")
	delete(decoded, "completionDateTime")
	delete(decoded, "copiesPrinted")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "duplexPageCount")
	delete(decoded, "id")
	delete(decoded, "pageCount")
	delete(decoded, "printerId")
	delete(decoded, "printerName")
	delete(decoded, "simplexPageCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ArchivedPrintJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ArchivedPrintJob{}

func (s *ArchivedPrintJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AcquiredByPrinter      *bool                    `json:"acquiredByPrinter,omitempty"`
		AcquiredDateTime       nullable.Type[string]    `json:"acquiredDateTime,omitempty"`
		BlackAndWhitePageCount *int64                   `json:"blackAndWhitePageCount,omitempty"`
		ColorPageCount         *int64                   `json:"colorPageCount,omitempty"`
		CompletionDateTime     nullable.Type[string]    `json:"completionDateTime,omitempty"`
		CopiesPrinted          *int64                   `json:"copiesPrinted,omitempty"`
		CreatedDateTime        *string                  `json:"createdDateTime,omitempty"`
		DuplexPageCount        *int64                   `json:"duplexPageCount,omitempty"`
		Id                     *string                  `json:"id,omitempty"`
		ODataId                *string                  `json:"@odata.id,omitempty"`
		ODataType              *string                  `json:"@odata.type,omitempty"`
		PageCount              *int64                   `json:"pageCount,omitempty"`
		PrinterId              nullable.Type[string]    `json:"printerId,omitempty"`
		PrinterName            nullable.Type[string]    `json:"printerName,omitempty"`
		ProcessingState        *PrintJobProcessingState `json:"processingState,omitempty"`
		SimplexPageCount       *int64                   `json:"simplexPageCount,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AcquiredByPrinter = decoded.AcquiredByPrinter
	s.AcquiredDateTime = decoded.AcquiredDateTime
	s.BlackAndWhitePageCount = decoded.BlackAndWhitePageCount
	s.ColorPageCount = decoded.ColorPageCount
	s.CompletionDateTime = decoded.CompletionDateTime
	s.CopiesPrinted = decoded.CopiesPrinted
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DuplexPageCount = decoded.DuplexPageCount
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PageCount = decoded.PageCount
	s.PrinterId = decoded.PrinterId
	s.PrinterName = decoded.PrinterName
	s.ProcessingState = decoded.ProcessingState
	s.SimplexPageCount = decoded.SimplexPageCount

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ArchivedPrintJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'ArchivedPrintJob': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}
