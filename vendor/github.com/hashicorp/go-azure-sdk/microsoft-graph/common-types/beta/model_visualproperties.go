package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VisualProperties struct {
	// The body of a visual user notification. Body is optional.
	Body nullable.Type[string] `json:"body,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The title of a visual user notification. This field is required for visual notification payloads.
	Title nullable.Type[string] `json:"title,omitempty"`
}
