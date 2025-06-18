package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingStaffMemberBase = BookingStaffMember{}

type BookingStaffMember struct {
	// True means that if the staff member is a Microsoft 365 user, the Bookings API would verify the staff member's
	// availability in their personal calendar in Microsoft 365, before making a booking.
	AvailabilityIsAffectedByPersonalCalendar *bool `json:"availabilityIsAffectedByPersonalCalendar,omitempty"`

	// The date, time, and time zone when the staff member was created. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the staff member, as displayed to customers. Required.
	DisplayName string `json:"displayName"`

	// The email address of the staff member. This email address can be in the same Microsoft 365 tenant as the business, or
	// in a different email domain. This email address can be used if the sendConfirmationsToOwner property is set to true
	// in the scheduling policy of the business. Required.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Indicates that a staff member is notified via email when a booking assigned to them is created or changed. The
	// default value is true.
	IsEmailNotificationEnabled *bool `json:"isEmailNotificationEnabled,omitempty"`

	// The date, time, and time zone when the staff member was last updated. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
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

func (s BookingStaffMember) BookingStaffMemberBase() BaseBookingStaffMemberBaseImpl {
	return BaseBookingStaffMemberBaseImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
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
