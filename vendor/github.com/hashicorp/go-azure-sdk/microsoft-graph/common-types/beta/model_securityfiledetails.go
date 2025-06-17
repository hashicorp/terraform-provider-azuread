package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileDetails struct {
	// The name of the file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The file path (location) of the file instance.
	FilePath nullable.Type[string] `json:"filePath,omitempty"`

	// The publisher of the file.
	FilePublisher nullable.Type[string] `json:"filePublisher,omitempty"`

	// The size of the file in bytes.
	FileSize nullable.Type[int64] `json:"fileSize,omitempty"`

	// The certificate authority (CA) that issued the certificate.
	Issuer nullable.Type[string] `json:"issuer,omitempty"`

	// The Md5 cryptographic hash of the file content.
	Md5 nullable.Type[string] `json:"md5,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Sha1 cryptographic hash of the file content.
	Sha1 nullable.Type[string] `json:"sha1,omitempty"`

	// The Sha256 cryptographic hash of the file content.
	Sha256 nullable.Type[string] `json:"sha256,omitempty"`

	Sha256Ac nullable.Type[string] `json:"sha256Ac,omitempty"`

	// The signer of the signed file.
	Signer nullable.Type[string] `json:"signer,omitempty"`
}
