package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatus struct {
	AcquiredByPrinter nullable.Type[bool] `json:"acquiredByPrinter,omitempty"`

	// A human-readable description of the print job's current processing state. Read-only.
	Description *string `json:"description,omitempty"`

	// Additional details for print job state. Valid values are described in the following table. Read-only.
	Details *[]PrintJobStateDetail `json:"details,omitempty"`

	// True if the job was acknowledged by a printer; false otherwise. Read-only.
	IsAcquiredByPrinter *bool `json:"isAcquiredByPrinter,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ProcessingState            *PrintJobProcessingState `json:"processingState,omitempty"`
	ProcessingStateDescription nullable.Type[string]    `json:"processingStateDescription,omitempty"`
	State                      *PrintJobProcessingState `json:"state,omitempty"`
}

var _ json.Marshaler = PrintJobStatus{}

func (s PrintJobStatus) MarshalJSON() ([]byte, error) {
	type wrapper PrintJobStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintJobStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintJobStatus: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "details")
	delete(decoded, "isAcquiredByPrinter")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintJobStatus: %+v", err)
	}

	return encoded, nil
}
