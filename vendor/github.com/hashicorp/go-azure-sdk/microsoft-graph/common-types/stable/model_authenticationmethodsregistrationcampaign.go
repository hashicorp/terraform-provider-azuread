package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaign struct {
	// Users and groups of users that are excluded from being prompted to set up the authentication method.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// Users and groups of users that are prompted to set up the authentication method.
	IncludeTargets *[]AuthenticationMethodsRegistrationCampaignIncludeTarget `json:"includeTargets,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt.
	// Minimum: 0 days. Maximum: 14 days. If the value is '0', the user is prompted during every MFA attempt.
	SnoozeDurationInDays *int64 `json:"snoozeDurationInDays,omitempty"`

	State *AdvancedConfigState `json:"state,omitempty"`
}
