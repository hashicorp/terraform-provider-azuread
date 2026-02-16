package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnAuthenticatorAttestationResponse struct {
	// Base64URL encoded, cryptographically protected attestation about the WebAuthn Authenticator.
	AttestationObject nullable.Type[string] `json:"attestationObject,omitempty"`

	ClientDataJSON nullable.Type[string] `json:"clientDataJSON,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
