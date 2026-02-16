package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = CallTranscriptEventMessageDetail{}

type CallTranscriptEventMessageDetail struct {
	// Unique identifier of the call.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// Unique identifier for a call transcript.
	CallTranscriptICalUid nullable.Type[string] `json:"callTranscriptICalUid,omitempty"`

	// The organizer of the meeting.
	MeetingOrganizer IdentitySet `json:"meetingOrganizer"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallTranscriptEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallTranscriptEventMessageDetail{}

func (s CallTranscriptEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper CallTranscriptEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallTranscriptEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallTranscriptEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callTranscriptEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallTranscriptEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallTranscriptEventMessageDetail{}

func (s *CallTranscriptEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CallId                nullable.Type[string] `json:"callId,omitempty"`
		CallTranscriptICalUid nullable.Type[string] `json:"callTranscriptICalUid,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CallId = decoded.CallId
	s.CallTranscriptICalUid = decoded.CallTranscriptICalUid
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallTranscriptEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["meetingOrganizer"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MeetingOrganizer' for 'CallTranscriptEventMessageDetail': %+v", err)
		}
		s.MeetingOrganizer = impl
	}

	return nil
}
