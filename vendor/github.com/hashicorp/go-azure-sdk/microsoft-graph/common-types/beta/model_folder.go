package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Folder struct {
	// Number of children contained immediately within this container.
	ChildCount nullable.Type[int64] `json:"childCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A collection of properties defining the recommended view for the folder.
	View *FolderView `json:"view,omitempty"`
}
