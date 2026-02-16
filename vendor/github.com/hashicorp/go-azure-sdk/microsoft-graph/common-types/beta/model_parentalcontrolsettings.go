package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParentalControlSettings struct {
	// Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries
	// specified in this list.
	CountriesBlockedForMinors *[]string `json:"countriesBlockedForMinors,omitempty"`

	// Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values:
	// ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the
	// European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with
	// COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country/region minor
	// rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country/region minor
	// rules.BlockMinorsBlocks minors from using the app.
	LegalAgeGroupRule nullable.Type[string] `json:"legalAgeGroupRule,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
