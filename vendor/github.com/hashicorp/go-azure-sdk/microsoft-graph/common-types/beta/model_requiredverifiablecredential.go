package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequiredVerifiableCredential struct {
	ClaimBindings *[]VerifiableCredentialClaimBinding `json:"claimBindings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TrustedIssuer            nullable.Type[string] `json:"trustedIssuer,omitempty"`
	VerifiableCredentialType nullable.Type[string] `json:"verifiableCredentialType,omitempty"`
}
