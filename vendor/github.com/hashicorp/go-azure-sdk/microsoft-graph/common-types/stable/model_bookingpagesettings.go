package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPageSettings struct {
	AccessControl *BookingPageAccessControl `json:"accessControl,omitempty"`

	// Custom color for bookings page. Value should be in Hex format. Example: `#123456`.
	BookingPageColorCode *string `json:"bookingPageColorCode,omitempty"`

	// The time zone of the customer. For a list of possible values, see
	// [dateTimeTimeZone](https://learn.microsoft.com/en-us/graph/api/resources/datetimetimezone?view=graph-rest-beta).
	BusinessTimeZone *string `json:"businessTimeZone,omitempty"`

	// Customer consent message that is displayed in the Booking page.
	CustomerConsentMessage *string `json:"customerConsentMessage,omitempty"`

	// Enforcing One Time Password (OTP) during appointment creation.
	EnforceOneTimePassword *bool `json:"enforceOneTimePassword,omitempty"`

	// Enable display of business logo display on the Bookings page.
	IsBusinessLogoDisplayEnabled *bool `json:"isBusinessLogoDisplayEnabled,omitempty"`

	// Enforces customer consent on the customer consent message before appointment is booked.
	IsCustomerConsentEnabled *bool `json:"isCustomerConsentEnabled,omitempty"`

	// Disable booking page to be indexed by search engines. False by default.
	IsSearchEngineIndexabilityDisabled *bool `json:"isSearchEngineIndexabilityDisabled,omitempty"`

	// If business time zone the default value for the time slots that we show in the bookings page. False by default.
	IsTimeSlotTimeZoneSetToBusinessTimeZone *bool `json:"isTimeSlotTimeZoneSetToBusinessTimeZone,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The URL of the business' Privacy Policy.
	PrivacyPolicyWebUrl *string `json:"privacyPolicyWebUrl,omitempty"`

	// The URL of the business' Terms and Conditions.
	TermsAndConditionsWebUrl *string `json:"termsAndConditionsWebUrl,omitempty"`
}
