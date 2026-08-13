package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingsAvailability interface {
	BookingsAvailability() BaseBookingsAvailabilityImpl
}

var _ BookingsAvailability = BaseBookingsAvailabilityImpl{}

type BaseBookingsAvailabilityImpl struct {
	AvailabilityType *BookingsServiceAvailabilityType `json:"availabilityType,omitempty"`

	// The hours of operation in a week. The business hours value is set to null if the availability type isn't
	// customWeeklyHours.
	BusinessHours *[]BookingWorkHours `json:"businessHours,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseBookingsAvailabilityImpl) BookingsAvailability() BaseBookingsAvailabilityImpl {
	return s
}

var _ BookingsAvailability = RawBookingsAvailabilityImpl{}

// RawBookingsAvailabilityImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawBookingsAvailabilityImpl struct {
	bookingsAvailability BaseBookingsAvailabilityImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawBookingsAvailabilityImpl) BookingsAvailability() BaseBookingsAvailabilityImpl {
	return s.bookingsAvailability
}

func (s RawBookingsAvailabilityImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalBookingsAvailabilityImplementation(input []byte) (BookingsAvailability, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingsAvailability into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingsAvailabilityWindow") {
		var out BookingsAvailabilityWindow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingsAvailabilityWindow: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingsAvailabilityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingsAvailabilityImpl: %+v", err)
	}

	return RawBookingsAvailabilityImpl{
		bookingsAvailability: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
