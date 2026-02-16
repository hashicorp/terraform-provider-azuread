package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintMargin struct {
	// The margin in microns from the bottom edge.
	Bottom nullable.Type[int64] `json:"bottom,omitempty"`

	// The margin in microns from the left edge.
	Left nullable.Type[int64] `json:"left,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The margin in microns from the right edge.
	Right nullable.Type[int64] `json:"right,omitempty"`

	// The margin in microns from the top edge.
	Top nullable.Type[int64] `json:"top,omitempty"`
}
