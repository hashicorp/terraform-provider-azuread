package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FolderView struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The method by which the folder should be sorted.
	SortBy nullable.Type[string] `json:"sortBy,omitempty"`

	// If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
	SortOrder nullable.Type[string] `json:"sortOrder,omitempty"`

	// The type of view that should be used to represent the folder.
	ViewType nullable.Type[string] `json:"viewType,omitempty"`
}
