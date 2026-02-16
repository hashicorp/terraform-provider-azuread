package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSystemInfo struct {
	// The UTC date and time the file was created on a client.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The UTC date and time the file was last accessed. Available for the recent file list only.
	LastAccessedDateTime nullable.Type[string] `json:"lastAccessedDateTime,omitempty"`

	// The UTC date and time the file was last modified on a client.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
