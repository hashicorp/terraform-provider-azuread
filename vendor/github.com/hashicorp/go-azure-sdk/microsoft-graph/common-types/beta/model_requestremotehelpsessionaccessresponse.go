package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestRemoteHelpSessionAccessResponse struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// AES encryption Initialization Vector for encrypting client messages sent to PubSub
	PubSubEncryption *string `json:"pubSubEncryption,omitempty"`

	// The unique identifier for encrypting client messages sent to PubSub
	PubSubEncryptionKey *string `json:"pubSubEncryptionKey,omitempty"`

	// The unique identifier for a session
	SessionKey *string `json:"sessionKey,omitempty"`
}
