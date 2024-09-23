package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogCollectionDownloadDetails struct {
	AppLogDecryptionAlgorithm *AppLogDecryptionAlgorithm `json:"appLogDecryptionAlgorithm,omitempty"`

	// Decryption key that used to decrypt the log.
	DecryptionKey nullable.Type[string] `json:"decryptionKey,omitempty"`

	// Download SAS (Shared Access Signature) Url for completed app log request.
	DownloadUrl nullable.Type[string] `json:"downloadUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
