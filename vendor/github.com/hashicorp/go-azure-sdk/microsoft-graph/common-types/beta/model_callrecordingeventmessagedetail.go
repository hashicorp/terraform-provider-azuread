package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = CallRecordingEventMessageDetail{}

type CallRecordingEventMessageDetail struct {
	// Unique identifier of the call.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// Display name for the call recording.
	CallRecordingDisplayName nullable.Type[string] `json:"callRecordingDisplayName,omitempty"`

	// Duration of the call recording.
	CallRecordingDuration nullable.Type[string] `json:"callRecordingDuration,omitempty"`

	// Status of the call recording. Possible values are: success, failure, initial, chunkFinished, unknownFutureValue.
	CallRecordingStatus *CallRecordingStatus `json:"callRecordingStatus,omitempty"`

	// Call recording URL.
	CallRecordingUrl nullable.Type[string] `json:"callRecordingUrl,omitempty"`

	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Organizer of the meeting.
	MeetingOrganizer IdentitySet `json:"meetingOrganizer"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallRecordingEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordingEventMessageDetail{}

func (s CallRecordingEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordingEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordingEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordingEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecordingEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordingEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordingEventMessageDetail{}

func (s *CallRecordingEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CallId                   nullable.Type[string] `json:"callId,omitempty"`
		CallRecordingDisplayName nullable.Type[string] `json:"callRecordingDisplayName,omitempty"`
		CallRecordingDuration    nullable.Type[string] `json:"callRecordingDuration,omitempty"`
		CallRecordingStatus      *CallRecordingStatus  `json:"callRecordingStatus,omitempty"`
		CallRecordingUrl         nullable.Type[string] `json:"callRecordingUrl,omitempty"`
		ODataId                  *string               `json:"@odata.id,omitempty"`
		ODataType                *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CallId = decoded.CallId
	s.CallRecordingDisplayName = decoded.CallRecordingDisplayName
	s.CallRecordingDuration = decoded.CallRecordingDuration
	s.CallRecordingStatus = decoded.CallRecordingStatus
	s.CallRecordingUrl = decoded.CallRecordingUrl
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRecordingEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'CallRecordingEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	if v, ok := temp["meetingOrganizer"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MeetingOrganizer' for 'CallRecordingEventMessageDetail': %+v", err)
		}
		s.MeetingOrganizer = impl
	}

	return nil
}
