package stable

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

	// The dateTimeOffset when the job was completed, canceled, or aborted. Read-only.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// The number of copies that were printed. Read-only.
	CopiesPrinted *int64 `json:"copiesPrinted,omitempty"`

	// The user who created the print job. Read-only.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// The dateTimeOffset when the job was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The archived print job's GUID. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The printer ID that the job was queued for. Read-only.
	PrinterId nullable.Type[string] `json:"printerId,omitempty"`

	// The printer name that the job was queued for. Read-only.
	PrinterName nullable.Type[string] `json:"printerName,omitempty"`

	ProcessingState *PrintJobProcessingState `json:"processingState,omitempty"`
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
	delete(decoded, "completionDateTime")
	delete(decoded, "copiesPrinted")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "id")
	delete(decoded, "printerId")
	delete(decoded, "printerName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ArchivedPrintJob: %+v", err)
	}

	return encoded, nil
}
