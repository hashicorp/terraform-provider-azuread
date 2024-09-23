package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TeamworkDeviceOperation{}

type TeamworkDeviceOperation struct {
	// Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// Identity of the user who created the device operation.
	CreatedBy IdentitySet `json:"createdBy"`

	// The UTC date and time when the device operation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Error details are available only in case of a failed status.
	Error *OperationError `json:"error,omitempty"`

	// Identity of the user who last modified the device operation.
	LastActionBy IdentitySet `json:"lastActionBy"`

	// The UTC date and time when the device operation was last modified.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	OperationType *TeamworkDeviceOperationType `json:"operationType,omitempty"`

	// Time at which the operation was started.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	// The current status of the async operation, for example, Queued, Scheduled, InProgress, Successful, Cancelled, and
	// Failed.
	Status *string `json:"status,omitempty"`

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

func (s TeamworkDeviceOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamworkDeviceOperation{}

func (s TeamworkDeviceOperation) MarshalJSON() ([]byte, error) {
	type wrapper TeamworkDeviceOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamworkDeviceOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkDeviceOperation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkDeviceOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamworkDeviceOperation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TeamworkDeviceOperation{}

func (s *TeamworkDeviceOperation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CompletedDateTime  nullable.Type[string]        `json:"completedDateTime,omitempty"`
		CreatedDateTime    nullable.Type[string]        `json:"createdDateTime,omitempty"`
		Error              *OperationError              `json:"error,omitempty"`
		LastActionDateTime nullable.Type[string]        `json:"lastActionDateTime,omitempty"`
		OperationType      *TeamworkDeviceOperationType `json:"operationType,omitempty"`
		StartedDateTime    nullable.Type[string]        `json:"startedDateTime,omitempty"`
		Status             *string                      `json:"status,omitempty"`
		Id                 *string                      `json:"id,omitempty"`
		ODataId            *string                      `json:"@odata.id,omitempty"`
		ODataType          *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.LastActionDateTime = decoded.LastActionDateTime
	s.OperationType = decoded.OperationType
	s.StartedDateTime = decoded.StartedDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TeamworkDeviceOperation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'TeamworkDeviceOperation': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastActionBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastActionBy' for 'TeamworkDeviceOperation': %+v", err)
		}
		s.LastActionBy = impl
	}

	return nil
}
