package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredentialConfiguration struct {
	// String value that indicates the maximum lifetime for password expiration, defined as an ISO 8601 duration. For
	// example, P4DT12H30M5S represents four days, 12 hours, 30 minutes, and five seconds. This property is required when
	// restrictionType is set to passwordLifetime.
	MaxLifetime nullable.Type[string] `json:"maxLifetime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the date from which the policy restriction applies to newly created applications. For existing
	// applications, the enforcement date can be retroactively applied.
	RestrictForAppsCreatedAfterDateTime nullable.Type[string] `json:"restrictForAppsCreatedAfterDateTime,omitempty"`

	// The type of restriction being applied. The possible values are: passwordAddition, passwordLifetime,
	// symmetricKeyAddition, symmetricKeyLifetime, customPasswordAddition, and unknownFutureValue. Each value of
	// restrictionType can be used only once per policy.
	RestrictionType *AppCredentialRestrictionType `json:"restrictionType,omitempty"`

	State *AppManagementRestrictionState `json:"state,omitempty"`
}
