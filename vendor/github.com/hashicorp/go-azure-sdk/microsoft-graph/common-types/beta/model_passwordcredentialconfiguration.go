package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredentialConfiguration struct {
	// Collection of custom security attribute exemptions. If an actor user or service principal has the custom security
	// attribute defined in this section, they're exempted from the restriction. This means that calls the user or service
	// principal makes to create or update apps are exempt from this policy enforcement.
	ExcludeActors *AppManagementPolicyActorExemptions `json:"excludeActors,omitempty"`

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

	RestrictionType *AppCredentialRestrictionType `json:"restrictionType,omitempty"`

	// Indicates whether the restriction is evaluated. The possible values are: enabled, disabled, unknownFutureValue. If
	// enabled, the restriction is evaluated. If disabled, the restriction isn't evaluated or enforced.
	State *AppManagementRestrictionState `json:"state,omitempty"`
}
