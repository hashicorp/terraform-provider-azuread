package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BookingBusiness{}

type BookingBusiness struct {
	// The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of
	// a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the
	// addresses to the type others.
	Address *PhysicalAddress `json:"address,omitempty"`

	// All the appointments of this business. Read-only. Nullable.
	Appointments *[]BookingAppointment `json:"appointments,omitempty"`

	BookingPageSettings *BookingPageSettings `json:"bookingPageSettings,omitempty"`

	// The hours of operation for the business.
	BusinessHours *[]BookingWorkHours `json:"businessHours,omitempty"`

	// The type of business.
	BusinessType nullable.Type[string] `json:"businessType,omitempty"`

	// The set of appointments of this business in a specified date range. Read-only. Nullable.
	CalendarView *[]BookingAppointment `json:"calendarView,omitempty"`

	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// All the custom questions of this business. Read-only. Nullable.
	CustomQuestions *[]BookingCustomQuestion `json:"customQuestions,omitempty"`

	// All the customers of this business. Read-only. Nullable.
	Customers *[]BookingCustomerBase `json:"customers,omitempty"`

	// The code for the currency that the business operates in on Microsoft Bookings.
	DefaultCurrencyIso nullable.Type[string] `json:"defaultCurrencyIso,omitempty"`

	// The name of the business, which interfaces with customers. This name appears at the top of the business scheduling
	// page.
	DisplayName *string `json:"displayName,omitempty"`

	// The email address for the business.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this
	// property. Read-only.
	IsPublished nullable.Type[bool] `json:"isPublished,omitempty"`

	// The language of the self-service booking page.
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer
	// of a business scheduling page.
	Phone nullable.Type[string] `json:"phone,omitempty"`

	// The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
	PublicUrl nullable.Type[string] `json:"publicUrl,omitempty"`

	// Specifies how bookings can be created for this business.
	SchedulingPolicy *BookingSchedulingPolicy `json:"schedulingPolicy,omitempty"`

	// All the services offered by this business. Read-only. Nullable.
	Services *[]BookingService `json:"services,omitempty"`

	// All the staff members that provide services in this business. Read-only. Nullable.
	StaffMembers *[]BookingStaffMemberBase `json:"staffMembers,omitempty"`

	// The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a
	// business scheduling page.
	WebSiteUrl nullable.Type[string] `json:"webSiteUrl,omitempty"`

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

func (s BookingBusiness) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingBusiness{}

func (s BookingBusiness) MarshalJSON() ([]byte, error) {
	type wrapper BookingBusiness
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingBusiness: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingBusiness: %+v", err)
	}

	delete(decoded, "appointments")
	delete(decoded, "calendarView")
	delete(decoded, "customQuestions")
	delete(decoded, "customers")
	delete(decoded, "isPublished")
	delete(decoded, "publicUrl")
	delete(decoded, "services")
	delete(decoded, "staffMembers")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.bookingBusiness"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingBusiness: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BookingBusiness{}

func (s *BookingBusiness) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Address             *PhysicalAddress         `json:"address,omitempty"`
		Appointments        *[]BookingAppointment    `json:"appointments,omitempty"`
		BookingPageSettings *BookingPageSettings     `json:"bookingPageSettings,omitempty"`
		BusinessHours       *[]BookingWorkHours      `json:"businessHours,omitempty"`
		BusinessType        nullable.Type[string]    `json:"businessType,omitempty"`
		CalendarView        *[]BookingAppointment    `json:"calendarView,omitempty"`
		CreatedDateTime     nullable.Type[string]    `json:"createdDateTime,omitempty"`
		CustomQuestions     *[]BookingCustomQuestion `json:"customQuestions,omitempty"`
		DefaultCurrencyIso  nullable.Type[string]    `json:"defaultCurrencyIso,omitempty"`
		DisplayName         *string                  `json:"displayName,omitempty"`
		Email               nullable.Type[string]    `json:"email,omitempty"`
		IsPublished         nullable.Type[bool]      `json:"isPublished,omitempty"`
		LanguageTag         nullable.Type[string]    `json:"languageTag,omitempty"`
		LastUpdatedDateTime nullable.Type[string]    `json:"lastUpdatedDateTime,omitempty"`
		Phone               nullable.Type[string]    `json:"phone,omitempty"`
		PublicUrl           nullable.Type[string]    `json:"publicUrl,omitempty"`
		SchedulingPolicy    *BookingSchedulingPolicy `json:"schedulingPolicy,omitempty"`
		Services            *[]BookingService        `json:"services,omitempty"`
		WebSiteUrl          nullable.Type[string]    `json:"webSiteUrl,omitempty"`
		Id                  *string                  `json:"id,omitempty"`
		ODataId             *string                  `json:"@odata.id,omitempty"`
		ODataType           *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Address = decoded.Address
	s.Appointments = decoded.Appointments
	s.BookingPageSettings = decoded.BookingPageSettings
	s.BusinessHours = decoded.BusinessHours
	s.BusinessType = decoded.BusinessType
	s.CalendarView = decoded.CalendarView
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomQuestions = decoded.CustomQuestions
	s.DefaultCurrencyIso = decoded.DefaultCurrencyIso
	s.DisplayName = decoded.DisplayName
	s.Email = decoded.Email
	s.IsPublished = decoded.IsPublished
	s.LanguageTag = decoded.LanguageTag
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.Phone = decoded.Phone
	s.PublicUrl = decoded.PublicUrl
	s.SchedulingPolicy = decoded.SchedulingPolicy
	s.Services = decoded.Services
	s.WebSiteUrl = decoded.WebSiteUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BookingBusiness into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["customers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Customers into list []json.RawMessage: %+v", err)
		}

		output := make([]BookingCustomerBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalBookingCustomerBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Customers' for 'BookingBusiness': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Customers = &output
	}

	if v, ok := temp["staffMembers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling StaffMembers into list []json.RawMessage: %+v", err)
		}

		output := make([]BookingStaffMemberBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalBookingStaffMemberBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'StaffMembers' for 'BookingBusiness': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.StaffMembers = &output
	}

	return nil
}
