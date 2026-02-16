package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallEvent = EmergencyCallEvent{}

type EmergencyCallEvent struct {
	// The information of the emergency caller.
	CallerInfo *EmergencyCallerInfo `json:"callerInfo,omitempty"`

	// The emergency number dialed.
	EmergencyNumberDialed nullable.Type[string] `json:"emergencyNumberDialed,omitempty"`

	// The policy name for emergency call event.
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// Fields inherited from CallEvent

	// The event type of the call. Possible values are: callStarted, callEnded, unknownFutureValue, rosterUpdated. You must
	// use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum:
	// rosterUpdated.
	CallEventType *CallEventType `json:"callEventType,omitempty"`

	// The time when event occurred.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// Participants collection for the call event.
	Participants *[]Participant `json:"participants,omitempty"`

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

func (s EmergencyCallEvent) CallEvent() BaseCallEventImpl {
	return BaseCallEventImpl{
		CallEventType: s.CallEventType,
		EventDateTime: s.EventDateTime,
		Participants:  s.Participants,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s EmergencyCallEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EmergencyCallEvent{}

func (s EmergencyCallEvent) MarshalJSON() ([]byte, error) {
	type wrapper EmergencyCallEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmergencyCallEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmergencyCallEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.emergencyCallEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmergencyCallEvent: %+v", err)
	}

	return encoded, nil
}
