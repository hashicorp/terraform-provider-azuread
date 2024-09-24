package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = CallEndedEventMessageDetail{}

type CallEndedEventMessageDetail struct {
	// Duration of the call.
	CallDuration nullable.Type[string] `json:"callDuration,omitempty"`

	// Represents the call event type. Possible values are: call, meeting, screenShare, unknownFutureValue.
	CallEventType *TeamworkCallEventType `json:"callEventType,omitempty"`

	// Unique identifier of the call.
	CallId nullable.Type[string] `json:"callId,omitempty"`

	// List of call participants.
	CallParticipants *[]CallParticipantInfo `json:"callParticipants,omitempty"`

	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallEndedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallEndedEventMessageDetail{}

func (s CallEndedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper CallEndedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallEndedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallEndedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callEndedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallEndedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallEndedEventMessageDetail{}

func (s *CallEndedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CallDuration     nullable.Type[string]  `json:"callDuration,omitempty"`
		CallEventType    *TeamworkCallEventType `json:"callEventType,omitempty"`
		CallId           nullable.Type[string]  `json:"callId,omitempty"`
		CallParticipants *[]CallParticipantInfo `json:"callParticipants,omitempty"`
		ODataId          *string                `json:"@odata.id,omitempty"`
		ODataType        *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CallDuration = decoded.CallDuration
	s.CallEventType = decoded.CallEventType
	s.CallId = decoded.CallId
	s.CallParticipants = decoded.CallParticipants
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallEndedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'CallEndedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
