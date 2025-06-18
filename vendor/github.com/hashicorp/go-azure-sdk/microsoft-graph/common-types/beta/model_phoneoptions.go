package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneOptions struct {
	// A read-only, Microsoft-defined list of regions that already enable MFA. For more information, see the following list
	// of countries.
	DefaultRegions *[]int64 `json:"defaultRegions,omitempty"`

	// A numbers-only set representing the region telecom codes to prevent or disable the telephony service. Validates
	// against current International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must
	// be non-null.
	ExcludeRegions *[]int64 `json:"excludeRegions,omitempty"`

	// A numbers-only set representing the country codes that can be manually added to enable telephony service in those
	// regions, in addition to the list of countries that are already enabled. For more information about regions that
	// require opt in, see Regions that need to opt in for MFA telephony verification. Validates against current
	// International Subscriber Dialing (ISD) country codes where the maximum code length is 4. Values must be positive
	// integers and can't overlap with 'excludeRegions'.
	IncludeAdditionalRegions *[]int64 `json:"includeAdditionalRegions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
