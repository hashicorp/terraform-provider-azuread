package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingNamedEntity interface {
	Entity
	BookingNamedEntity() BaseBookingNamedEntityImpl
}

var _ BookingNamedEntity = BaseBookingNamedEntityImpl{}

type BaseBookingNamedEntityImpl struct {
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

func (s BaseBookingNamedEntityImpl) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return s
}

func (s BaseBookingNamedEntityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ BookingNamedEntity = RawBookingNamedEntityImpl{}

// RawBookingNamedEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBookingNamedEntityImpl struct {
	bookingNamedEntity BaseBookingNamedEntityImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawBookingNamedEntityImpl) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return s.bookingNamedEntity
}

func (s RawBookingNamedEntityImpl) Entity() BaseEntityImpl {
	return s.bookingNamedEntity.Entity()
}

var _ json.Marshaler = BaseBookingNamedEntityImpl{}

func (s BaseBookingNamedEntityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseBookingNamedEntityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseBookingNamedEntityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseBookingNamedEntityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingNamedEntity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseBookingNamedEntityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalBookingNamedEntityImplementation(input []byte) (BookingNamedEntity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingNamedEntity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingBusiness") {
		var out BookingBusiness
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingBusiness: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingPerson") {
		var out BookingPerson
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingPerson: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingService") {
		var out BookingService
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingService: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingNamedEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingNamedEntityImpl: %+v", err)
	}

	return RawBookingNamedEntityImpl{
		bookingNamedEntity: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
