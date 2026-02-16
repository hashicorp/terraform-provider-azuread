package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnPublicKeyCredential struct {
	// The untyped results from the execution of extensions requested by the client when creating a new public key
	// credential.
	ClientExtensionResults *WebauthnAuthenticationExtensionsClientOutputs `json:"clientExtensionResults,omitempty"`

	// The credential ID created by the WebAuthn Authenticator.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Response data returned from a WebAuthn authenticator after it creates a new public key credential.
	Response *WebauthnAuthenticatorAttestationResponse `json:"response,omitempty"`
}
