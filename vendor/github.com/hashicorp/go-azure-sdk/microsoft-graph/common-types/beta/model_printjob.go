package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrintJob{}

type PrintJob struct {
	AcknowledgedDateTime nullable.Type[string]  `json:"acknowledgedDateTime,omitempty"`
	CompletedDateTime    nullable.Type[string]  `json:"completedDateTime,omitempty"`
	Configuration        *PrintJobConfiguration `json:"configuration,omitempty"`
	CreatedBy            *UserIdentity          `json:"createdBy,omitempty"`

	// The DateTimeOffset when the job was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The name of the print job.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Documents *[]PrintDocument     `json:"documents,omitempty"`
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

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PrintJob{}

func (s *PrintJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AcknowledgedDateTime nullable.Type[string]  `json:"acknowledgedDateTime,omitempty"`
		CompletedDateTime    nullable.Type[string]  `json:"completedDateTime,omitempty"`
		Configuration        *PrintJobConfiguration `json:"configuration,omitempty"`
		CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]  `json:"displayName,omitempty"`
		Documents            *[]PrintDocument       `json:"documents,omitempty"`
		ErrorCode            nullable.Type[int64]   `json:"errorCode,omitempty"`
		IsFetchable          *bool                  `json:"isFetchable,omitempty"`
		RedirectedFrom       nullable.Type[string]  `json:"redirectedFrom,omitempty"`
		RedirectedTo         nullable.Type[string]  `json:"redirectedTo,omitempty"`
		Status               *PrintJobStatus        `json:"status,omitempty"`
		Tasks                *[]PrintTask           `json:"tasks,omitempty"`
		Id                   *string                `json:"id,omitempty"`
		ODataId              *string                `json:"@odata.id,omitempty"`
		ODataType            *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AcknowledgedDateTime = decoded.AcknowledgedDateTime
	s.CompletedDateTime = decoded.CompletedDateTime
	s.Configuration = decoded.Configuration
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Documents = decoded.Documents
	s.ErrorCode = decoded.ErrorCode
	s.IsFetchable = decoded.IsFetchable
	s.RedirectedFrom = decoded.RedirectedFrom
	s.RedirectedTo = decoded.RedirectedTo
	s.Status = decoded.Status
	s.Tasks = decoded.Tasks
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PrintJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PrintJob': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}
