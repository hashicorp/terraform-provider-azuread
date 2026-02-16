package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSecurityState struct {
	// Complex type containing file hashes (cryptographic and location-sensitive).
	FileHash *FileHash `json:"fileHash,omitempty"`

	// File name (without path).
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Full file path of the file/imageFile.
	Path nullable.Type[string] `json:"path,omitempty"`

	// Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a
	// percentage.
	RiskScore nullable.Type[string] `json:"riskScore,omitempty"`
}
