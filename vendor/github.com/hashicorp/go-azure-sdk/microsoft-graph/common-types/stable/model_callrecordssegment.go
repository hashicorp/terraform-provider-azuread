package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CallRecordsSegment{}

type CallRecordsSegment struct {
	// Endpoint that answered this segment.
	Callee CallRecordsEndpoint `json:"callee"`

	// Endpoint that initiated this segment.
	Caller CallRecordsEndpoint `json:"caller"`

	// UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Failure information associated with the segment if it failed.
	FailureInfo *CallRecordsFailureInfo `json:"failureInfo,omitempty"`

	// Media associated with this segment.
	Media *[]CallRecordsMedia `json:"media,omitempty"`

	// UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime *string `json:"startDateTime,omitempty"`

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

func (s CallRecordsSegment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsSegment{}

func (s CallRecordsSegment) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsSegment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsSegment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsSegment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.segment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsSegment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordsSegment{}

func (s *CallRecordsSegment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndDateTime   *string                 `json:"endDateTime,omitempty"`
		FailureInfo   *CallRecordsFailureInfo `json:"failureInfo,omitempty"`
		Media         *[]CallRecordsMedia     `json:"media,omitempty"`
		StartDateTime *string                 `json:"startDateTime,omitempty"`
		Id            *string                 `json:"id,omitempty"`
		ODataId       *string                 `json:"@odata.id,omitempty"`
		ODataType     *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EndDateTime = decoded.EndDateTime
	s.FailureInfo = decoded.FailureInfo
	s.Media = decoded.Media
	s.StartDateTime = decoded.StartDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRecordsSegment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["callee"]; ok {
		impl, err := UnmarshalCallRecordsEndpointImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Callee' for 'CallRecordsSegment': %+v", err)
		}
		s.Callee = impl
	}

	if v, ok := temp["caller"]; ok {
		impl, err := UnmarshalCallRecordsEndpointImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Caller' for 'CallRecordsSegment': %+v", err)
		}
		s.Caller = impl
	}

	return nil
}
