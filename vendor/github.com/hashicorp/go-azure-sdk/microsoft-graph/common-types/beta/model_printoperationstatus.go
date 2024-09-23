package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperationStatus struct {
	// A human-readable description of the printOperation's current processing state. Read-only.
	Description *string `json:"description,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *PrintOperationProcessingState `json:"state,omitempty"`
}

var _ json.Marshaler = PrintOperationStatus{}

func (s PrintOperationStatus) MarshalJSON() ([]byte, error) {
	type wrapper PrintOperationStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintOperationStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintOperationStatus: %+v", err)
	}

	delete(decoded, "description")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintOperationStatus: %+v", err)
	}

	return encoded, nil
}
