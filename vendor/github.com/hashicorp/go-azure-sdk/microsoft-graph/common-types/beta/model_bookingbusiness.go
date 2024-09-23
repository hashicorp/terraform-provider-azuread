package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingNamedEntity = BookingBusiness{}

type BookingBusiness struct {
	// The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of
	// a business scheduling page.
	Address *PhysicalAddress `json:"address,omitempty"`

	// All the appointments of this business. Read-only. Nullable.
	Appointments *[]BookingAppointment `json:"appointments,omitempty"`

	// Settings for the published booking page.
	BookingPageSettings *BookingPageSettings `json:"bookingPageSettings,omitempty"`

	// The hours of operation for the business.
	BusinessHours *[]BookingWorkHours `json:"businessHours,omitempty"`

	// The type of business.
	BusinessType nullable.Type[string] `json:"businessType,omitempty"`

	// The set of appointments of this business in a specified date range. Read-only. Nullable.
	CalendarView *[]BookingAppointment `json:"calendarView,omitempty"`

	// The date, time and timezone when the booking business was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// All custom questions of this business.
	CustomQuestions *[]BookingCustomQuestion `json:"customQuestions,omitempty"`

	// All the customers of this business. Read-only. Nullable.
	Customers *[]BookingCustomer `json:"customers,omitempty"`

	// The code for the currency that the business operates in on Microsoft Bookings.
	DefaultCurrencyIso nullable.Type[string] `json:"defaultCurrencyIso,omitempty"`

	// The email address for the business.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this
	// property. Read-only.
	IsPublished nullable.Type[bool] `json:"isPublished,omitempty"`

	// The language of the self service booking page
	LanguageTag nullable.Type[string] `json:"languageTag,omitempty"`

	// The date, time and timezone when the booking business was last updated.
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
	StaffMembers *[]BookingStaffMember `json:"staffMembers,omitempty"`

	// The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a
	// business scheduling page.
	WebSiteUrl nullable.Type[string] `json:"webSiteUrl,omitempty"`

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

func (s BookingBusiness) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return BaseBookingNamedEntityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
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
