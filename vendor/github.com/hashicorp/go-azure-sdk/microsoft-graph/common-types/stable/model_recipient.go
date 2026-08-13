package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Recipient interface {
	Recipient() BaseRecipientImpl
}

var _ Recipient = BaseRecipientImpl{}

type BaseRecipientImpl struct {
	// The recipient's email address.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseRecipientImpl) Recipient() BaseRecipientImpl {
	return s
}

var _ Recipient = RawRecipientImpl{}

// RawRecipientImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawRecipientImpl struct {
	recipient BaseRecipientImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawRecipientImpl) Recipient() BaseRecipientImpl {
	return s.recipient
}

func (s RawRecipientImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalRecipientImplementation(input []byte) (Recipient, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Recipient into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.attendeeBase") {
		var out AttendeeBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AttendeeBase: %+v", err)
		}
		return out, nil
	}

	var parent BaseRecipientImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRecipientImpl: %+v", err)
	}

	return RawRecipientImpl{
		recipient: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
