package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatus struct {
	// A human-readable description of the printer's current processing state. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The list of details describing why the printer is in the current state. Valid values are described in the following
	// table. Read-only.
	Details *[]PrinterProcessingStateDetail `json:"details,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ProcessingState            *PrinterProcessingState         `json:"processingState,omitempty"`
	ProcessingStateDescription nullable.Type[string]           `json:"processingStateDescription,omitempty"`
	ProcessingStateReasons     *[]PrinterProcessingStateReason `json:"processingStateReasons,omitempty"`
	State                      *PrinterProcessingState         `json:"state,omitempty"`
}

var _ json.Marshaler = PrinterStatus{}

func (s PrinterStatus) MarshalJSON() ([]byte, error) {
	type wrapper PrinterStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrinterStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrinterStatus: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "details")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrinterStatus: %+v", err)
	}

	return encoded, nil
}
