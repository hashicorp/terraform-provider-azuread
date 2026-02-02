package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceVisualization struct {
	// A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying
	// the owner of the OneDrive storing the item.
	ContainerDisplayName nullable.Type[string] `json:"containerDisplayName,omitempty"`

	// Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
	ContainerType nullable.Type[string] `json:"containerType,omitempty"`

	// A path leading to the folder in which the item is stored.
	ContainerWebUrl nullable.Type[string] `json:"containerWebUrl,omitempty"`

	// The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime
	// Types. Not all Media Mime Types are supported.
	MediaType nullable.Type[string] `json:"mediaType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A URL leading to the preview image for the item.
	PreviewImageUrl nullable.Type[string] `json:"previewImageUrl,omitempty"`

	// A preview text for the item.
	PreviewText nullable.Type[string] `json:"previewText,omitempty"`

	// The item's title text.
	Title nullable.Type[string] `json:"title,omitempty"`

	// The item's media type. Can be used for filtering for a specific file based on a specific type. See the section Type
	// property values for supported types.
	Type nullable.Type[string] `json:"type,omitempty"`
}
