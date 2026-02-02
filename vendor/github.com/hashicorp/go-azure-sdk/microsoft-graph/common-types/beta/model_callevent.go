package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEvent interface {
	Entity
	CallEvent() BaseCallEventImpl
}

var _ CallEvent = BaseCallEventImpl{}

type BaseCallEventImpl struct {
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

func (s BaseCallEventImpl) CallEvent() BaseCallEventImpl {
	return s
}

func (s BaseCallEventImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CallEvent = RawCallEventImpl{}

// RawCallEventImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCallEventImpl struct {
	callEvent BaseCallEventImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawCallEventImpl) CallEvent() BaseCallEventImpl {
	return s.callEvent
}

func (s RawCallEventImpl) Entity() BaseEntityImpl {
	return s.callEvent.Entity()
}

var _ json.Marshaler = BaseCallEventImpl{}

func (s BaseCallEventImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCallEventImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCallEventImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCallEventImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCallEventImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalCallEventImplementation(input []byte) (CallEvent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallEvent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.emergencyCallEvent") {
		var out EmergencyCallEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmergencyCallEvent: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallEventImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallEventImpl: %+v", err)
	}

	return RawCallEventImpl{
		callEvent: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
