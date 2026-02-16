package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionLifetimePolicy struct {
	// The human-readable details of the conditional access session management policy applied to the sign-in.
	Detail nullable.Type[string] `json:"detail,omitempty"`

	// If a conditional access session management policy required the user to authenticate in this sign-in event, this field
	// describes the policy type that required authentication. The possible values are:
	// rememberMultifactorAuthenticationOnTrustedDevices, tenantTokenLifetimePolicy, audienceTokenLifetimePolicy,
	// signInFrequencyPeriodicReauthentication, ngcMfa, signInFrequencyEveryTime, unknownFutureValue.
	ExpirationRequirement *ExpirationRequirement `json:"expirationRequirement,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
