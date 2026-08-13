package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallOptions interface {
	CallOptions() BaseCallOptionsImpl
}

var _ CallOptions = BaseCallOptionsImpl{}

type BaseCallOptionsImpl struct {
	// Indicates whether to hide the app after the call is escalated.
	HideBotAfterEscalation nullable.Type[bool] `json:"hideBotAfterEscalation,omitempty"`

	// Indicates whether content sharing notifications should be enabled for the call.
	IsContentSharingNotificationEnabled nullable.Type[bool] `json:"isContentSharingNotificationEnabled,omitempty"`

	// Indicates whether delta roster is enabled for the call.
	IsDeltaRosterEnabled nullable.Type[bool] `json:"isDeltaRosterEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCallOptionsImpl) CallOptions() BaseCallOptionsImpl {
	return s
}

var _ CallOptions = RawCallOptionsImpl{}

// RawCallOptionsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCallOptionsImpl struct {
	callOptions BaseCallOptionsImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawCallOptionsImpl) CallOptions() BaseCallOptionsImpl {
	return s.callOptions
}

func (s RawCallOptionsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCallOptionsImplementation(input []byte) (CallOptions, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallOptions into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.incomingCallOptions") {
		var out IncomingCallOptions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IncomingCallOptions: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.outgoingCallOptions") {
		var out OutgoingCallOptions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OutgoingCallOptions: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallOptionsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallOptionsImpl: %+v", err)
	}

	return RawCallOptionsImpl{
		callOptions: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
