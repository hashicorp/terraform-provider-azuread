package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerSettings struct {
	// Indicates whether versioning is enabled for items in the container. Optional. Read-write.
	IsItemVersioningEnabled nullable.Type[bool] `json:"isItemVersioningEnabled,omitempty"`

	// Indicates whether Optical Character Recognition (OCR) is enabled for the container. The default value is false. When
	// set to true, OCR extraction is performed for new and updated documents of supported document types, and the extracted
	// fields in the metadata of the document enable end-user search and search-driven solutions. When set to false,
	// existing OCR metadata is not impacted. Optional. Read-write.
	IsOcrEnabled nullable.Type[bool] `json:"isOcrEnabled,omitempty"`

	// The maximum major versions allowed for items in the container. Optional. Read-write.
	ItemMajorVersionLimit nullable.Type[int64] `json:"itemMajorVersionLimit,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
