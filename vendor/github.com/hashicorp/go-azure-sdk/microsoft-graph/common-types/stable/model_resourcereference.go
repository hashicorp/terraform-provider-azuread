package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceReference struct {
	// The item's unique identifier.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
	Type nullable.Type[string] `json:"type,omitempty"`

	// A URL leading to the referenced item.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
