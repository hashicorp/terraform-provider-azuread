package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CallTranscript{}

type CallTranscript struct {
	// The unique identifier for the call that is related to this transcript. Read-only.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// The content of the transcript. Read-only.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The unique identifier that links the transcript with its corresponding recording. Read-only.
	ContentCorrelationId nullable.Type[string] `json:"contentCorrelationId,omitempty"`

	// Date and time at which the transcript was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Date and time at which the transcription ends. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The unique identifier of the online meeting related to this transcript. Read-only.
	MeetingId nullable.Type[string] `json:"meetingId,omitempty"`

	// The identity information of the organizer of the onlineMeeting related to this transcript. Read-only.
	MeetingOrganizer *IdentitySet `json:"meetingOrganizer,omitempty"`

	// The time-aligned metadata of the utterances in the transcript. Read-only.
	MetadataContent nullable.Type[string] `json:"metadataContent,omitempty"`

	// The URL that can be used to access the content of the transcript. Read-only.
	TranscriptContentUrl nullable.Type[string] `json:"transcriptContentUrl,omitempty"`

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

func (s CallTranscript) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallTranscript{}

func (s CallTranscript) MarshalJSON() ([]byte, error) {
	type wrapper CallTranscript
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallTranscript: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallTranscript: %+v", err)
	}

	delete(decoded, "callId")
	delete(decoded, "content")
	delete(decoded, "contentCorrelationId")
	delete(decoded, "createdDateTime")
	delete(decoded, "endDateTime")
	delete(decoded, "meetingId")
	delete(decoded, "meetingOrganizer")
	delete(decoded, "metadataContent")
	delete(decoded, "transcriptContentUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callTranscript"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallTranscript: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallTranscript{}

func (s *CallTranscript) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CallId               nullable.Type[string] `json:"callId,omitempty"`
		Content              nullable.Type[string] `json:"content,omitempty"`
		ContentCorrelationId nullable.Type[string] `json:"contentCorrelationId,omitempty"`
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		EndDateTime          nullable.Type[string] `json:"endDateTime,omitempty"`
		MeetingId            nullable.Type[string] `json:"meetingId,omitempty"`
		MetadataContent      nullable.Type[string] `json:"metadataContent,omitempty"`
		TranscriptContentUrl nullable.Type[string] `json:"transcriptContentUrl,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CallId = decoded.CallId
	s.Content = decoded.Content
	s.ContentCorrelationId = decoded.ContentCorrelationId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.EndDateTime = decoded.EndDateTime
	s.MeetingId = decoded.MeetingId
	s.MetadataContent = decoded.MetadataContent
	s.TranscriptContentUrl = decoded.TranscriptContentUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallTranscript into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["meetingOrganizer"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MeetingOrganizer' for 'CallTranscript': %+v", err)
		}
		s.MeetingOrganizer = &impl
	}

	return nil
}
