package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingPerson = BookingStaffMember{}

type BookingStaffMember struct {
	// True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's
	// availability in their personal calendar in Microsoft 365, before making a booking.
	AvailabilityIsAffectedByPersonalCalendar *bool `json:"availabilityIsAffectedByPersonalCalendar,omitempty"`

	// Identifies a color to represent the staff member. The color corresponds to the color palette in the Staff details
	// page in the Bookings app.
	ColorIndex nullable.Type[int64] `json:"colorIndex,omitempty"`

	// The date, time and timezone when the staff member was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates that a staff members are notified via email when a booking assigned to them is created or changed. The
	// default value is true
	IsEmailNotificationEnabled *bool `json:"isEmailNotificationEnabled,omitempty"`

	// The date, time and timezone when the staff member was last updated.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	MembershipStatus *BookingStaffMembershipStatus `json:"membershipStatus,omitempty"`
	Role             *BookingStaffRole             `json:"role,omitempty"`

	// The time zone of the staff member. For a list of possible values, see dateTimeTimeZone.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`

	// True means the staff member's availability is as specified in the businessHours property of the business. False means
	// the availability is determined by the staff member's workingHours property setting.
	UseBusinessHours *bool `json:"useBusinessHours,omitempty"`

	// The range of hours each day of the week that the staff member is available for booking. By default, they're
	// initialized to be the same as the businessHours property of the business.
	WorkingHours *[]BookingWorkHours `json:"workingHours,omitempty"`

	// Fields inherited from BookingPerson

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

func (s BookingStaffMember) BookingPerson() BaseBookingPersonImpl {
	return BaseBookingPersonImpl{
		EmailAddress: s.EmailAddress,
		DisplayName:  s.DisplayName,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s BookingStaffMember) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return BaseBookingNamedEntityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s BookingStaffMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingStaffMember{}

func (s BookingStaffMember) MarshalJSON() ([]byte, error) {
	type wrapper BookingStaffMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingStaffMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingStaffMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingStaffMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingStaffMember: %+v", err)
	}

	return encoded, nil
}
