package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElevationRequestApplicationDetail struct {
	// The path of the file in the request for elevation, for example, %programfiles%/git/cmd
	FileDescription nullable.Type[string] `json:"fileDescription,omitempty"`

	// The SHA256 hash of the file in the request for elevation, for example,
	// '18ee24150dcb1d96752a4d6dd0f20dfd8ba8c38527e40aa8509b7adecf78f9c6'
	FileHash nullable.Type[string] `json:"fileHash,omitempty"`

	// The name of the file in the request for elevation, for example, git.exe
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The path of the file in the request for elevation, for example, %programfiles%/git/cmd
	FilePath nullable.Type[string] `json:"filePath,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The internal name of the application for which elevation request has been made. For example, 'git'
	ProductInternalName nullable.Type[string] `json:"productInternalName,omitempty"`

	// The product name of the application for which elevation request has been made. For example, 'Git'
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The product version of the application for which elevation request has been made. For example, '2.40.1.0'
	ProductVersion nullable.Type[string] `json:"productVersion,omitempty"`

	// The list of base64 encoded certificate for each signer, for example, string[encodedleafcert1, encodedleafcert2....]
	PublisherCert nullable.Type[string] `json:"publisherCert,omitempty"`

	// The certificate issuer name of the certificate used to sign the application, for example, 'Sectigo Public Code
	// Signing CA R36'
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`
}
