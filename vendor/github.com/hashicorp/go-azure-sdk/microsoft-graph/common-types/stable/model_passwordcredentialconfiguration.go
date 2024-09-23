package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredentialConfiguration struct {
	// Value that can be used as the maximum number for setting password expiration time in days, hours, minutes or seconds.
	// Defined in ISO 8601 format for Durations. For example, 'P4DT12H30M5S' represents a duration of four days, twelve
	// hours, thirty minutes, and five seconds. This property is required when restriction type is set to passwordLifetime.
	MaxLifetime nullable.Type[string] `json:"maxLifetime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Enforces the policy for an app created on or after the enforcement date. For existing applications, the enforcement
	// date would be back dated. To apply to all applications, enforcement datetime would be null.
	RestrictForAppsCreatedAfterDateTime nullable.Type[string] `json:"restrictForAppsCreatedAfterDateTime,omitempty"`

	// The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime,
	// symmetricKeyAddition, symmetricKeyLifetime,customPasswordAddition, unknownFutureValue. Each value of restrictionType
	// can be used only once per policy.
	RestrictionType *AppCredentialRestrictionType `json:"restrictionType,omitempty"`
}
