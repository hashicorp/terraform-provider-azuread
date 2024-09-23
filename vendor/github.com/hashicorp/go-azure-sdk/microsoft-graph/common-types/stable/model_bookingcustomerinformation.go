package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingCustomerInformationBase = BookingCustomerInformation{}

type BookingCustomerInformation struct {
	// It consists of the list of custom questions and answers given by the customer as part of the appointment
	CustomQuestionAnswers *[]BookingQuestionAnswer `json:"customQuestionAnswers,omitempty"`

	// The ID of the bookingCustomer for this appointment. If no ID is specified when an appointment is created, then a new
	// bookingCustomer object is created. Once set, you should consider the customerId immutable.
	CustomerId nullable.Type[string] `json:"customerId,omitempty"`

	// The SMTP address of the bookingCustomer who is booking the appointment
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Represents location information for the bookingCustomer who is booking the appointment.
	Location Location `json:"location"`

	// The customer's name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Notes from the customer associated with this appointment. You can get the value only when reading this
	// bookingAppointment by its ID. You can set this property only when initially creating an appointment with a new
	// customer. After that point, the value is computed from the customer represented by the customerId.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The customer's phone number.
	Phone nullable.Type[string] `json:"phone,omitempty"`

	// The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`

	// Fields inherited from BookingCustomerInformationBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BookingCustomerInformation) BookingCustomerInformationBase() BaseBookingCustomerInformationBaseImpl {
	return BaseBookingCustomerInformationBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingCustomerInformation{}

func (s BookingCustomerInformation) MarshalJSON() ([]byte, error) {
	type wrapper BookingCustomerInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingCustomerInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingCustomerInformation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingCustomerInformation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingCustomerInformation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BookingCustomerInformation{}

func (s *BookingCustomerInformation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CustomQuestionAnswers *[]BookingQuestionAnswer `json:"customQuestionAnswers,omitempty"`
		CustomerId            nullable.Type[string]    `json:"customerId,omitempty"`
		EmailAddress          nullable.Type[string]    `json:"emailAddress,omitempty"`
		Name                  nullable.Type[string]    `json:"name,omitempty"`
		Notes                 nullable.Type[string]    `json:"notes,omitempty"`
		Phone                 nullable.Type[string]    `json:"phone,omitempty"`
		TimeZone              nullable.Type[string]    `json:"timeZone,omitempty"`
		ODataId               *string                  `json:"@odata.id,omitempty"`
		ODataType             *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CustomQuestionAnswers = decoded.CustomQuestionAnswers
	s.CustomerId = decoded.CustomerId
	s.EmailAddress = decoded.EmailAddress
	s.Name = decoded.Name
	s.Notes = decoded.Notes
	s.Phone = decoded.Phone
	s.TimeZone = decoded.TimeZone
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BookingCustomerInformation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["location"]; ok {
		impl, err := UnmarshalLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Location' for 'BookingCustomerInformation': %+v", err)
		}
		s.Location = impl
	}

	return nil
}
