package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemAccessOperationsViewpoint struct {
	// Indicates whether the user can comment on this item.
	CanComment nullable.Type[bool] `json:"canComment,omitempty"`

	// Indicates whether the user can create files within this object. Returned only on folders.
	CanCreateFile nullable.Type[bool] `json:"canCreateFile,omitempty"`

	// Indicates whether the user can create folders within this object. Returned only on folders.
	CanCreateFolder nullable.Type[bool] `json:"canCreateFolder,omitempty"`

	// Indicates whether the user can delete this item.
	CanDelete nullable.Type[bool] `json:"canDelete,omitempty"`

	// Indicates whether the user can download this item.
	CanDownload nullable.Type[bool] `json:"canDownload,omitempty"`

	// Indicates whether the user can read this item.
	CanRead nullable.Type[bool] `json:"canRead,omitempty"`

	// Indicates whether the user can update this item.
	CanUpdate nullable.Type[bool] `json:"canUpdate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
