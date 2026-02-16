package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportFileMetadata struct {
	DownloadUrl nullable.Type[string] `json:"downloadUrl,omitempty"`
	FileName    nullable.Type[string] `json:"fileName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Size nullable.Type[int64] `json:"size,omitempty"`
}
