package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrintTask{}

type PrintTask struct {
	Definition *PrintTaskDefinition `json:"definition,omitempty"`

	// The URL for the print entity that triggered this task. For example,
	// https://graph.microsoft.com/v1.0/print/printers/{printerId}/jobs/{jobId}. Read-only.
	ParentUrl *string `json:"parentUrl,omitempty"`

	Status  *PrintTaskStatus  `json:"status,omitempty"`
	Trigger *PrintTaskTrigger `json:"trigger,omitempty"`

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

func (s PrintTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrintTask{}

func (s PrintTask) MarshalJSON() ([]byte, error) {
	type wrapper PrintTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintTask: %+v", err)
	}

	delete(decoded, "parentUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintTask: %+v", err)
	}

	return encoded, nil
}
