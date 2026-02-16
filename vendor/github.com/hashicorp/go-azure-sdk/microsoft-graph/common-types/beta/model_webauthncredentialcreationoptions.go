package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnCredentialCreationOptions struct {
	// Defines when the challenge in the creation options is no longer valid. Expired challenges are rejected when you
	// attempt to create a new fido2AuthenticationMethod.
	ChallengeTimeoutDateTime nullable.Type[string] `json:"challengeTimeoutDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines public key options for the creation of a new WebAuthn public key credential.
	PublicKey *WebauthnPublicKeyCredentialCreationOptions `json:"publicKey,omitempty"`
}
