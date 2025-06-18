package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingSchedulingPolicy struct {
	// True to allow customers to choose a specific person for the booking.
	AllowStaffSelection *bool `json:"allowStaffSelection,omitempty"`

	// Custom availability of the service in a given time frame.
	CustomAvailabilities *[]BookingsAvailabilityWindow `json:"customAvailabilities,omitempty"`

	// General availability of the service defined by the scheduling policy.
	GeneralAvailability BookingsAvailability `json:"generalAvailability"`

	// Indicates whether the meeting invite is sent to the customers. The default value is false.
	IsMeetingInviteToCustomersEnabled nullable.Type[bool] `json:"isMeetingInviteToCustomersEnabled,omitempty"`

	// Maximum number of days in advance that a booking can be made. It follows the ISO 8601 format.
	MaximumAdvance *string `json:"maximumAdvance,omitempty"`

	// The minimum amount of time before which bookings and cancellations must be made. It follows the ISO 8601 format.
	MinimumLeadTime *string `json:"minimumLeadTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// True to notify the business via email when a booking is created or changed. Use the email address specified in the
	// email property of the bookingBusiness entity for the business.
	SendConfirmationsToOwner *bool `json:"sendConfirmationsToOwner,omitempty"`

	// Duration of each time slot, denoted in ISO 8601 format.
	TimeSlotInterval *string `json:"timeSlotInterval,omitempty"`
}

var _ json.Unmarshaler = &BookingSchedulingPolicy{}

func (s *BookingSchedulingPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowStaffSelection               *bool                         `json:"allowStaffSelection,omitempty"`
		CustomAvailabilities              *[]BookingsAvailabilityWindow `json:"customAvailabilities,omitempty"`
		IsMeetingInviteToCustomersEnabled nullable.Type[bool]           `json:"isMeetingInviteToCustomersEnabled,omitempty"`
		MaximumAdvance                    *string                       `json:"maximumAdvance,omitempty"`
		MinimumLeadTime                   *string                       `json:"minimumLeadTime,omitempty"`
		ODataId                           *string                       `json:"@odata.id,omitempty"`
		ODataType                         *string                       `json:"@odata.type,omitempty"`
		SendConfirmationsToOwner          *bool                         `json:"sendConfirmationsToOwner,omitempty"`
		TimeSlotInterval                  *string                       `json:"timeSlotInterval,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowStaffSelection = decoded.AllowStaffSelection
	s.CustomAvailabilities = decoded.CustomAvailabilities
	s.IsMeetingInviteToCustomersEnabled = decoded.IsMeetingInviteToCustomersEnabled
	s.MaximumAdvance = decoded.MaximumAdvance
	s.MinimumLeadTime = decoded.MinimumLeadTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SendConfirmationsToOwner = decoded.SendConfirmationsToOwner
	s.TimeSlotInterval = decoded.TimeSlotInterval

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BookingSchedulingPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["generalAvailability"]; ok {
		impl, err := UnmarshalBookingsAvailabilityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GeneralAvailability' for 'BookingSchedulingPolicy': %+v", err)
		}
		s.GeneralAvailability = impl
	}

	return nil
}
