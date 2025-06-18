package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionItem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The display name of the owner of the action item.
	OwnerDisplayName nullable.Type[string] `json:"ownerDisplayName,omitempty"`

	// The text content of the action item.
	Text nullable.Type[string] `json:"text,omitempty"`

	// The title of the action item.
	Title nullable.Type[string] `json:"title,omitempty"`
}
