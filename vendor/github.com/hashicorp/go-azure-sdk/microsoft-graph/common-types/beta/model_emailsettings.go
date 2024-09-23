package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailSettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the domain that should be used when sending email notifications. This domain must be verified in order to
	// be used. We recommend that you use a domain that has the appropriate DNS records to facilitate email validation, like
	// SPF, DKIM, DMARC, and MX, because this then complies with the RFC compliance for sending and receiving email. For
	// details, see Learn more about Exchange Online Email Routing.
	SenderDomain *string `json:"senderDomain,omitempty"`

	// Specifies if the organization’s banner logo should be included in email notifications. The banner logo will replace
	// the Microsoft logo at the top of the email notification. If true the banner logo will be taken from the tenant’s
	// branding settings. This value can only be set to true if the organizationalBranding bannerLogo property is set.
	UseCompanyBranding *bool `json:"useCompanyBranding,omitempty"`
}
