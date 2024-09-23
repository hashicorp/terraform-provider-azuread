package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CallRecordsSession{}

type CallRecordsSession struct {
	// Endpoint that answered the session.
	Callee CallRecordsEndpoint `json:"callee"`

	// Endpoint that initiated the session.
	Caller CallRecordsEndpoint `json:"caller"`

	// UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Failure information associated with the session if the session failed.
	FailureInfo *CallRecordsFailureInfo `json:"failureInfo,omitempty"`

	// Specifies whether the session is a test.
	IsTest nullable.Type[bool] `json:"isTest,omitempty"`

	// List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data,
	// screenSharing, unknownFutureValue.
	Modalities *[]CallRecordsModality `json:"modalities,omitempty"`

	// The list of segments involved in the session. Read-only. Nullable.
	Segments *[]CallRecordsSegment `json:"segments,omitempty"`

	// UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
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

func (s CallRecordsSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsSession{}

func (s CallRecordsSession) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsSession: %+v", err)
	}

	delete(decoded, "segments")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.session"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsSession: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordsSession{}

func (s *CallRecordsSession) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndDateTime   *string                 `json:"endDateTime,omitempty"`
		FailureInfo   *CallRecordsFailureInfo `json:"failureInfo,omitempty"`
		IsTest        nullable.Type[bool]     `json:"isTest,omitempty"`
		Modalities    *[]CallRecordsModality  `json:"modalities,omitempty"`
		Segments      *[]CallRecordsSegment   `json:"segments,omitempty"`
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
	s.IsTest = decoded.IsTest
	s.Modalities = decoded.Modalities
	s.Segments = decoded.Segments
	s.StartDateTime = decoded.StartDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRecordsSession into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["callee"]; ok {
		impl, err := UnmarshalCallRecordsEndpointImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Callee' for 'CallRecordsSession': %+v", err)
		}
		s.Callee = impl
	}

	if v, ok := temp["caller"]; ok {
		impl, err := UnmarshalCallRecordsEndpointImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Caller' for 'CallRecordsSession': %+v", err)
		}
		s.Caller = impl
	}

	return nil
}
