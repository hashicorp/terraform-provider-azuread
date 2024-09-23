package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredentialConfiguration struct {
	// Value that can be used as the maximum duration in days, hours, minutes, or seconds from the date of key creation, for
	// which the key is valid. Defined in ISO 8601 format for Durations. For example, P4DT12H30M5S represents a duration of
	// four days, twelve hours, thirty minutes, and five seconds. This property is required when restrictionType is set to
	// keyLifetime.
	MaxLifetime nullable.Type[string] `json:"maxLifetime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Timestamp when the policy is enforced for all apps created on or after the specified date. For existing applications,
	// the enforcement date would be back dated. To apply to all applications regardless of their creation date, this
	// property would be null. Nullable.
	RestrictForAppsCreatedAfterDateTime nullable.Type[string] `json:"restrictForAppsCreatedAfterDateTime,omitempty"`

	// The type of restriction being applied. Possible values are asymmetricKeyLifetime, unknownFutureValue. Each value of
	// restrictionType can be used only once per policy.
	RestrictionType *AppKeyCredentialRestrictionType `json:"restrictionType,omitempty"`
}
