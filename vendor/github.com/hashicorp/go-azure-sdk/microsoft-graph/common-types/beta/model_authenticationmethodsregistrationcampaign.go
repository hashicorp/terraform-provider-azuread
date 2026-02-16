package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaign struct {
	// Specifies whether a user is required to perform registration after snoozing 3 times. If true, the user is required to
	// register after 3 snoozes. If false, the user can snooze indefinitely. The default value is true.
	EnforceRegistrationAfterAllowedSnoozes nullable.Type[bool] `json:"enforceRegistrationAfterAllowedSnoozes,omitempty"`

	// Users and groups of users that are excluded from being prompted to set up the authentication method.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// Users and groups of users that are prompted to set up the authentication method.
	IncludeTargets *[]AuthenticationMethodsRegistrationCampaignIncludeTarget `json:"includeTargets,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the number of days that the user sees a prompt again if they select 'Not now' and snoozes the prompt.
	// Minimum 0 days. Maximum: 14 days. If the value is 0 – The user is prompted during every MFA attempt.
	SnoozeDurationInDays *int64 `json:"snoozeDurationInDays,omitempty"`

	State *AdvancedConfigState `json:"state,omitempty"`
}
