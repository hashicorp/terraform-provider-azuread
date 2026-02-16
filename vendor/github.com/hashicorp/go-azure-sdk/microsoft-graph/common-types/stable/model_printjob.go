package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrintJob{}

type PrintJob struct {
	// The dateTimeOffset when the job was acknowledged. Read-only.
	AcknowledgedDateTime nullable.Type[string] `json:"acknowledgedDateTime,omitempty"`

	Configuration *PrintJobConfiguration `json:"configuration,omitempty"`
	CreatedBy     *UserIdentity          `json:"createdBy,omitempty"`

	// The DateTimeOffset when the job was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	Documents *[]PrintDocument `json:"documents,omitempty"`

	// The error code of the print job. Read-only.
	ErrorCode nullable.Type[int64] `json:"errorCode,omitempty"`

	// If true, document can be fetched by printer.
	IsFetchable *bool `json:"isFetchable,omitempty"`

	// Contains the source job URL, if the job has been redirected from another printer.
	RedirectedFrom nullable.Type[string] `json:"redirectedFrom,omitempty"`

	// Contains the destination job URL, if the job has been redirected to another printer.
	RedirectedTo nullable.Type[string] `json:"redirectedTo,omitempty"`

	Status *PrintJobStatus `json:"status,omitempty"`

	// A list of printTasks that were triggered by this print job.
	Tasks *[]PrintTask `json:"tasks,omitempty"`

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

func (s PrintJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrintJob{}

func (s PrintJob) MarshalJSON() ([]byte, error) {
	type wrapper PrintJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintJob: %+v", err)
	}

	delete(decoded, "acknowledgedDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "errorCode")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintJob: %+v", err)
	}

	return encoded, nil
}
