package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPageSettings struct {
	AccessControl *BookingPageAccessControl `json:"accessControl,omitempty"`

	// Custom color for the bookings page. The value should be in Hex format. Example: #123456.
	BookingPageColorCode *string `json:"bookingPageColorCode,omitempty"`

	// The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
	BusinessTimeZone *string `json:"businessTimeZone,omitempty"`

	// The personal data collection and usage consent message in the bookings page.
	CustomerConsentMessage *string `json:"customerConsentMessage,omitempty"`

	// Determines if the one-time password is required to create an appointment. The default value is false.
	EnforceOneTimePassword *bool `json:"enforceOneTimePassword,omitempty"`

	// Indicates if the business logo is displayed on the bookings page. The default value is false.
	IsBusinessLogoDisplayEnabled *bool `json:"isBusinessLogoDisplayEnabled,omitempty"`

	// Enables personal data collection and the usage consent toggle on the bookings page. The default value is false.
	IsCustomerConsentEnabled *bool `json:"isCustomerConsentEnabled,omitempty"`

	// Ensures that the web crawlers don't index this page. The defaults value is false.
	IsSearchEngineIndexabilityDisabled *bool `json:"isSearchEngineIndexabilityDisabled,omitempty"`

	// Displays the booking time slots in the business time zone. The default value is false.
	IsTimeSlotTimeZoneSetToBusinessTimeZone *bool `json:"isTimeSlotTimeZoneSetToBusinessTimeZone,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// RL of a webpage that provides the terms and conditions of the business. If a privacy policy isn't included, the
	// following text appears on the bookings page as default: 'The policies and practices of <booking business's name>
	// apply to the use of your data.
	PrivacyPolicyWebUrl *string `json:"privacyPolicyWebUrl,omitempty"`

	// URL of a webpage that provides the terms and conditions of the business.
	TermsAndConditionsWebUrl *string `json:"termsAndConditionsWebUrl,omitempty"`
}
