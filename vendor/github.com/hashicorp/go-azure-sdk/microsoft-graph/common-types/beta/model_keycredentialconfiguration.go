package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredentialConfiguration struct {
	// Collection of GUIDs that represent certificateBasedApplicationConfiguration that is allowed as root and intermediate
	// certificate authorities.
	CertificateBasedApplicationConfigurationIds *[]string `json:"certificateBasedApplicationConfigurationIds,omitempty"`

	// String value that indicates the maximum lifetime for key expiration, defined as an ISO 8601 duration. For example,
	// P4DT12H30M5S represents four days, 12 hours, 30 minutes, and five seconds. This property is required when
	// restrictionType is set to keyLifetime.
	MaxLifetime nullable.Type[string] `json:"maxLifetime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the date from which the policy restriction applies to newly created applications. For existing
	// applications, the enforcement date can be retroactively applied.
	RestrictForAppsCreatedAfterDateTime nullable.Type[string] `json:"restrictForAppsCreatedAfterDateTime,omitempty"`

	// The type of restriction being applied. Possible values are asymmetricKeyLifetime, and unknownFutureValue. Each value
	// of restrictionType can be used only once per policy.
	RestrictionType *AppKeyCredentialRestrictionType `json:"restrictionType,omitempty"`

	// String value that indicates if the restriction is evaluated. The possible values are: enabled, disabled, and
	// unknownFutureValue. If enabled, the restriction is evaluated. If disabled, the restriction isn't evaluated or
	// enforced.
	State *AppManagementRestrictionState `json:"state,omitempty"`
}
