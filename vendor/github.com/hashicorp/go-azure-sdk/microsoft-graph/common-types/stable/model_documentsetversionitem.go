package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentSetVersionItem struct {
	// The unique identifier for the item.
	ItemId nullable.Type[string] `json:"itemId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The title of the item.
	Title nullable.Type[string] `json:"title,omitempty"`

	// The version ID of the item.
	VersionId nullable.Type[string] `json:"versionId,omitempty"`
}
