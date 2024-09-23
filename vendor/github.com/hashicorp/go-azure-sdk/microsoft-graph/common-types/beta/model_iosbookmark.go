package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosBookmark struct {
	// The folder into which the bookmark should be added in Safari
	BookmarkFolder nullable.Type[string] `json:"bookmarkFolder,omitempty"`

	// The display name of the bookmark
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// URL allowed to access
	Url *string `json:"url,omitempty"`
}
