package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingStaffMemberBase interface {
	Entity
	BookingStaffMemberBase() BaseBookingStaffMemberBaseImpl
}

var _ BookingStaffMemberBase = BaseBookingStaffMemberBaseImpl{}

type BaseBookingStaffMemberBaseImpl struct {

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

func (s BaseBookingStaffMemberBaseImpl) BookingStaffMemberBase() BaseBookingStaffMemberBaseImpl {
	return s
}

func (s BaseBookingStaffMemberBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ BookingStaffMemberBase = RawBookingStaffMemberBaseImpl{}

// RawBookingStaffMemberBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBookingStaffMemberBaseImpl struct {
	bookingStaffMemberBase BaseBookingStaffMemberBaseImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawBookingStaffMemberBaseImpl) BookingStaffMemberBase() BaseBookingStaffMemberBaseImpl {
	return s.bookingStaffMemberBase
}

func (s RawBookingStaffMemberBaseImpl) Entity() BaseEntityImpl {
	return s.bookingStaffMemberBase.Entity()
}

var _ json.Marshaler = BaseBookingStaffMemberBaseImpl{}

func (s BaseBookingStaffMemberBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseBookingStaffMemberBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseBookingStaffMemberBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseBookingStaffMemberBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingStaffMemberBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseBookingStaffMemberBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalBookingStaffMemberBaseImplementation(input []byte) (BookingStaffMemberBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingStaffMemberBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingStaffMember") {
		var out BookingStaffMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingStaffMember: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingStaffMemberBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingStaffMemberBaseImpl: %+v", err)
	}

	return RawBookingStaffMemberBaseImpl{
		bookingStaffMemberBase: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
