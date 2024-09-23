package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebauthnPublicKeyCredentialDescriptor struct {
	// The unique identifier of the credential.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The transport method used by the WebAuthn authenticator to communicate with the client. For example, usb, nfc, ble.
	Transports *[]string `json:"transports,omitempty"`

	// Type of public key credential. The only supported value is public-key.
	Type nullable.Type[string] `json:"type,omitempty"`
}
