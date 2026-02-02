package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemUploadableProperties struct {
	// Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Information about the drive item source. Read-write. Only on OneDrive for Business and SharePoint.
	DriveItemSource *DriveItemSource `json:"driveItemSource,omitempty"`

	// Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
	FileSize nullable.Type[int64] `json:"fileSize,omitempty"`

	// File system information on client. Read-write.
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`

	// Media source information. Read-write. Only on OneDrive for Business and SharePoint.
	MediaSource *MediaSource `json:"mediaSource,omitempty"`

	// The name of the item (filename and extension). Read-write.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
