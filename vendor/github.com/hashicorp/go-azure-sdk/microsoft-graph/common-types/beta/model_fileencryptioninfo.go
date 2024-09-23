package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileEncryptionInfo struct {
	// The key used to encrypt the file content.
	EncryptionKey nullable.Type[string] `json:"encryptionKey,omitempty"`

	// The file digest prior to encryption. ProfileVersion1 requires a non-null FileDigest.
	FileDigest nullable.Type[string] `json:"fileDigest,omitempty"`

	// The file digest algorithm. ProfileVersion1 currently only supports SHA256 for the FileDigestAlgorithm.
	FileDigestAlgorithm nullable.Type[string] `json:"fileDigestAlgorithm,omitempty"`

	// The initialization vector (IV) used for the encryption algorithm. Must be 16 bytes.
	InitializationVector nullable.Type[string] `json:"initializationVector,omitempty"`

	// The hash of the concatenation of the IV and encrypted file content. Must be 32 bytes.
	Mac nullable.Type[string] `json:"mac,omitempty"`

	// The key used to compute the message authentication code of the concatenation of the IV and encrypted file content.
	// Must be 32 bytes.
	MacKey nullable.Type[string] `json:"macKey,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The profile identifier. Maps to the strategy used to encrypt the file. Currently, only ProfileVersion1 is supported.
	ProfileIdentifier nullable.Type[string] `json:"profileIdentifier,omitempty"`
}
