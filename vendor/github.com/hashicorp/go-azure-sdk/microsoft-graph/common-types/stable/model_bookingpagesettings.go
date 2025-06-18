package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingPageSettings struct {
	AccessControl *BookingPageAccessControl `json:"accessControl,omitempty"`

	// Custom color for the booking page. The value should be in Hex format. For example, #123456.
	BookingPageColorCode *string `json:"bookingPageColorCode,omitempty"`

	// The time zone of the customer. For a list of possible values, see dateTimeTimeZone.
	BusinessTimeZone *string `json:"businessTimeZone,omitempty"`

	// The personal data collection and usage consent message in the booking page.
	CustomerConsentMessage *string `json:"customerConsentMessage,omitempty"`

	// Determines whether the one-time password is required to create an appointment. The default value is false.
	EnforceOneTimePassword *bool `json:"enforceOneTimePassword,omitempty"`

	// Indicates whether the business logo is displayed on the booking page. The default value is false.
	IsBusinessLogoDisplayEnabled *bool `json:"isBusinessLogoDisplayEnabled,omitempty"`

	// Enables personal data collection and the usage consent toggle on the booking page. The default value is false.
	IsCustomerConsentEnabled *bool `json:"isCustomerConsentEnabled,omitempty"`

	// Indicates whether web crawlers index this page. The defaults value is false.
	IsSearchEngineIndexabilityDisabled *bool `json:"isSearchEngineIndexabilityDisabled,omitempty"`

	// Indicates whether the time zone of the time slot is set to the time zone of the business. The default value is false.
	IsTimeSlotTimeZoneSetToBusinessTimeZone *bool `json:"isTimeSlotTimeZoneSetToBusinessTimeZone,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// URL of a webpage that provides the terms and conditions of the business. If a privacy policy isn't included, the
	// following text appears on the booking page as default: 'The policies and practices of {bookingbusinessname} apply to
	// the use of your data.'
	PrivacyPolicyWebUrl *string `json:"privacyPolicyWebUrl,omitempty"`

	// URL of a webpage that provides the terms and conditions of the business.
	TermsAndConditionsWebUrl *string `json:"termsAndConditionsWebUrl,omitempty"`
}
