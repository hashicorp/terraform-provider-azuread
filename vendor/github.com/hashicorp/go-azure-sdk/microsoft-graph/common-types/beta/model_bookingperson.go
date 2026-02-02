package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPerson interface {
	Entity
	BookingNamedEntity
	BookingPerson() BaseBookingPersonImpl
}

var _ BookingPerson = BaseBookingPersonImpl{}

type BaseBookingPersonImpl struct {
	// The email address of the person.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Fields inherited from BookingNamedEntity

	// A name for the derived entity, which interfaces with customers.
	DisplayName *string `json:"displayName,omitempty"`

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

func (s BaseBookingPersonImpl) BookingPerson() BaseBookingPersonImpl {
	return s
}

func (s BaseBookingPersonImpl) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return BaseBookingNamedEntityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s BaseBookingPersonImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ BookingPerson = RawBookingPersonImpl{}

// RawBookingPersonImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBookingPersonImpl struct {
	bookingPerson BaseBookingPersonImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawBookingPersonImpl) BookingPerson() BaseBookingPersonImpl {
	return s.bookingPerson
}

func (s RawBookingPersonImpl) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return s.bookingPerson.BookingNamedEntity()
}

func (s RawBookingPersonImpl) Entity() BaseEntityImpl {
	return s.bookingPerson.Entity()
}

var _ json.Marshaler = BaseBookingPersonImpl{}

func (s BaseBookingPersonImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseBookingPersonImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseBookingPersonImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseBookingPersonImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingPerson"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseBookingPersonImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalBookingPersonImplementation(input []byte) (BookingPerson, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingPerson into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomer") {
		var out BookingCustomer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingStaffMember") {
		var out BookingStaffMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingStaffMember: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingPersonImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingPersonImpl: %+v", err)
	}

	return RawBookingPersonImpl{
		bookingPerson: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
