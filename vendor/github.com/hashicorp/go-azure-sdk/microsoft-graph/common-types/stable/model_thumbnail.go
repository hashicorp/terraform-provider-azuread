package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Thumbnail struct {
	// The content stream for the thumbnail.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The height of the thumbnail, in pixels.
	Height nullable.Type[int64] `json:"height,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is
	// requested.
	SourceItemId nullable.Type[string] `json:"sourceItemId,omitempty"`

	// The URL used to fetch the thumbnail content.
	Url nullable.Type[string] `json:"url,omitempty"`

	// The width of the thumbnail, in pixels.
	Width nullable.Type[int64] `json:"width,omitempty"`
}
