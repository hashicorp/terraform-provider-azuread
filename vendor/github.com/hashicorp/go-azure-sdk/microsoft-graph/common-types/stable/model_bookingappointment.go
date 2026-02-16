package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BookingAppointment{}

type BookingAppointment struct {
	// Additional information that is sent to the customer when an appointment is confirmed.
	AdditionalInformation nullable.Type[string] `json:"additionalInformation,omitempty"`

	// The URL of the meeting to join anonymously.
	AnonymousJoinWebUrl nullable.Type[string] `json:"anonymousJoinWebUrl,omitempty"`

	// The custom label that can be stamped on this appointment by users.
	AppointmentLabel nullable.Type[string] `json:"appointmentLabel,omitempty"`

	// The date, time, and time zone when the appointment was created. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The SMTP address of the bookingCustomer who books the appointment.
	CustomerEmailAddress nullable.Type[string] `json:"customerEmailAddress,omitempty"`

	// The customer's name.
	CustomerName nullable.Type[string] `json:"customerName,omitempty"`

	// Notes from the customer associated with this appointment. You can get the value only when you read this
	// bookingAppointment by its ID. You can set this property only when you initially create an appointment with a new
	// customer.
	CustomerNotes nullable.Type[string] `json:"customerNotes,omitempty"`

	// The customer's phone number.
	CustomerPhone nullable.Type[string] `json:"customerPhone,omitempty"`

	// The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
	CustomerTimeZone nullable.Type[string] `json:"customerTimeZone,omitempty"`

	// A collection of customer properties for an appointment. An appointment contains a list of customer information and
	// each unit will indicate the properties of a customer who is part of that appointment. Optional.
	Customers *[]BookingCustomerInformationBase `json:"customers,omitempty"`

	// The length of the appointment, denoted in ISO8601 format.
	Duration *string `json:"duration,omitempty"`

	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`

	// The current number of customers in the appointment.
	FilledAttendeesCount *int64 `json:"filledAttendeesCount,omitempty"`

	// Indicates that the customer can manage bookings created by the staff. The default value is false.
	IsCustomerAllowedToManageBooking nullable.Type[bool] `json:"isCustomerAllowedToManageBooking,omitempty"`

	// Indicates that the appointment is held online. The default value is false.
	IsLocationOnline *bool `json:"isLocationOnline,omitempty"`

	// The URL of the online meeting for the appointment.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

	// The date, time, and time zone when the booking business was last updated. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The maximum number of customers allowed in an appointment. If maximumAttendeesCount of the service is greater than 1,
	// pass valid customer IDs while creating or updating an appointment. To create a customer, use the Create
	// bookingCustomer operation.
	MaximumAttendeesCount *int64 `json:"maximumAttendeesCount,omitempty"`

	// If true indicates that the bookingCustomer for this appointment doesn't wish to receive a confirmation for this
	// appointment.
	OptOutOfCustomerEmail *bool `json:"optOutOfCustomerEmail,omitempty"`

	// The amount of time to reserve after the appointment ends, for cleaning up, as an example. The value is expressed in
	// ISO8601 format.
	PostBuffer *string `json:"postBuffer,omitempty"`

	// The amount of time to reserve before the appointment begins, for preparation, as an example. The value is expressed
	// in ISO8601 format.
	PreBuffer *string `json:"preBuffer,omitempty"`

	// Represents the type of pricing of a booking service.
	PriceType *BookingPriceType `json:"priceType,omitempty"`

	// The collection of customer reminders sent for this appointment. The value of this property is available only when
	// reading this bookingAppointment by its ID.
	Reminders *[]BookingReminder `json:"reminders,omitempty"`

	// Another tracking ID for the appointment, if the appointment was created directly by the customer on the scheduling
	// page, as opposed to by a staff member on behalf of the customer.
	SelfServiceAppointmentId nullable.Type[string] `json:"selfServiceAppointmentId,omitempty"`

	// The ID of the bookingService associated with this appointment.
	ServiceId nullable.Type[string] `json:"serviceId,omitempty"`

	// The location where the service is delivered.
	ServiceLocation Location `json:"serviceLocation"`

	// The name of the bookingService associated with this appointment.This property is optional when creating a new
	// appointment. If not specified, it's computed from the service associated with the appointment by the serviceId
	// property.
	ServiceName *string `json:"serviceName,omitempty"`

	// Notes from a bookingStaffMember. The value of this property is available only when reading this bookingAppointment by
	// its ID.
	ServiceNotes nullable.Type[string] `json:"serviceNotes,omitempty"`

	// If true, indicates SMS notifications will be sent to the customers for the appointment. Default value is false.
	SmsNotificationsEnabled *bool `json:"smsNotificationsEnabled,omitempty"`

	// The ID of each bookingStaffMember who is scheduled in this appointment.
	StaffMemberIds *[]string `json:"staffMemberIds,omitempty"`

	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

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

func (s BookingAppointment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingAppointment{}

func (s BookingAppointment) MarshalJSON() ([]byte, error) {
	type wrapper BookingAppointment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingAppointment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingAppointment: %+v", err)
	}

	delete(decoded, "duration")
	delete(decoded, "filledAttendeesCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingAppointment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingAppointment: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BookingAppointment{}

func (s *BookingAppointment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdditionalInformation            nullable.Type[string] `json:"additionalInformation,omitempty"`
		AnonymousJoinWebUrl              nullable.Type[string] `json:"anonymousJoinWebUrl,omitempty"`
		AppointmentLabel                 nullable.Type[string] `json:"appointmentLabel,omitempty"`
		CreatedDateTime                  nullable.Type[string] `json:"createdDateTime,omitempty"`
		CustomerEmailAddress             nullable.Type[string] `json:"customerEmailAddress,omitempty"`
		CustomerName                     nullable.Type[string] `json:"customerName,omitempty"`
		CustomerNotes                    nullable.Type[string] `json:"customerNotes,omitempty"`
		CustomerPhone                    nullable.Type[string] `json:"customerPhone,omitempty"`
		CustomerTimeZone                 nullable.Type[string] `json:"customerTimeZone,omitempty"`
		Duration                         *string               `json:"duration,omitempty"`
		EndDateTime                      *DateTimeTimeZone     `json:"endDateTime,omitempty"`
		FilledAttendeesCount             *int64                `json:"filledAttendeesCount,omitempty"`
		IsCustomerAllowedToManageBooking nullable.Type[bool]   `json:"isCustomerAllowedToManageBooking,omitempty"`
		IsLocationOnline                 *bool                 `json:"isLocationOnline,omitempty"`
		JoinWebUrl                       nullable.Type[string] `json:"joinWebUrl,omitempty"`
		LastUpdatedDateTime              nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`
		MaximumAttendeesCount            *int64                `json:"maximumAttendeesCount,omitempty"`
		OptOutOfCustomerEmail            *bool                 `json:"optOutOfCustomerEmail,omitempty"`
		PostBuffer                       *string               `json:"postBuffer,omitempty"`
		PreBuffer                        *string               `json:"preBuffer,omitempty"`
		PriceType                        *BookingPriceType     `json:"priceType,omitempty"`
		Reminders                        *[]BookingReminder    `json:"reminders,omitempty"`
		SelfServiceAppointmentId         nullable.Type[string] `json:"selfServiceAppointmentId,omitempty"`
		ServiceId                        nullable.Type[string] `json:"serviceId,omitempty"`
		ServiceName                      *string               `json:"serviceName,omitempty"`
		ServiceNotes                     nullable.Type[string] `json:"serviceNotes,omitempty"`
		SmsNotificationsEnabled          *bool                 `json:"smsNotificationsEnabled,omitempty"`
		StaffMemberIds                   *[]string             `json:"staffMemberIds,omitempty"`
		StartDateTime                    *DateTimeTimeZone     `json:"startDateTime,omitempty"`
		Id                               *string               `json:"id,omitempty"`
		ODataId                          *string               `json:"@odata.id,omitempty"`
		ODataType                        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdditionalInformation = decoded.AdditionalInformation
	s.AnonymousJoinWebUrl = decoded.AnonymousJoinWebUrl
	s.AppointmentLabel = decoded.AppointmentLabel
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomerEmailAddress = decoded.CustomerEmailAddress
	s.CustomerName = decoded.CustomerName
	s.CustomerNotes = decoded.CustomerNotes
	s.CustomerPhone = decoded.CustomerPhone
	s.CustomerTimeZone = decoded.CustomerTimeZone
	s.Duration = decoded.Duration
	s.EndDateTime = decoded.EndDateTime
	s.FilledAttendeesCount = decoded.FilledAttendeesCount
	s.IsCustomerAllowedToManageBooking = decoded.IsCustomerAllowedToManageBooking
	s.IsLocationOnline = decoded.IsLocationOnline
	s.JoinWebUrl = decoded.JoinWebUrl
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.MaximumAttendeesCount = decoded.MaximumAttendeesCount
	s.OptOutOfCustomerEmail = decoded.OptOutOfCustomerEmail
	s.PostBuffer = decoded.PostBuffer
	s.PreBuffer = decoded.PreBuffer
	s.PriceType = decoded.PriceType
	s.Reminders = decoded.Reminders
	s.SelfServiceAppointmentId = decoded.SelfServiceAppointmentId
	s.ServiceId = decoded.ServiceId
	s.ServiceName = decoded.ServiceName
	s.ServiceNotes = decoded.ServiceNotes
	s.SmsNotificationsEnabled = decoded.SmsNotificationsEnabled
	s.StaffMemberIds = decoded.StaffMemberIds
	s.StartDateTime = decoded.StartDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BookingAppointment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Customers into list []json.RawMessage: %+v", err)
		}

		output := make([]BookingCustomerInformationBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalBookingCustomerInformationBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Customers' for 'BookingAppointment': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Customers = &output
	}

	if v, ok := temp["serviceLocation"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ServiceLocation' for 'BookingAppointment': %+v", err)
		}
		s.ServiceLocation = impl
	}

	return nil
}
