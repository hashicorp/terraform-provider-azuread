package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailAddress interface {
	EmailAddress() BaseEmailAddressImpl
}

var _ EmailAddress = BaseEmailAddressImpl{}

type BaseEmailAddressImpl struct {
	// The email address of an entity instance.
	Address nullable.Type[string] `json:"address,omitempty"`

	// The display name of an entity instance.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEmailAddressImpl) EmailAddress() BaseEmailAddressImpl {
	return s
}

var _ EmailAddress = RawEmailAddressImpl{}

// RawEmailAddressImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEmailAddressImpl struct {
	emailAddress BaseEmailAddressImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawEmailAddressImpl) EmailAddress() BaseEmailAddressImpl {
	return s.emailAddress
}

func (s RawEmailAddressImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEmailAddressImplementation(input []byte) (EmailAddress, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EmailAddress into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.typedEmailAddress") {
		var out TypedEmailAddress
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TypedEmailAddress: %+v", err)
		}
		return out, nil
	}

	var parent BaseEmailAddressImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEmailAddressImpl: %+v", err)
	}

	return RawEmailAddressImpl{
		emailAddress: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
