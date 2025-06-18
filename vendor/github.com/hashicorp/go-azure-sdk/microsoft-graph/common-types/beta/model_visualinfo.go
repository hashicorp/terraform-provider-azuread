package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VisualInfo struct {
	// Optional. JSON object used to represent an icon which represents the application used to generate the activity
	Attribution *ImageInfo `json:"attribution,omitempty"`

	// Optional. Background color used to render the activity in the UI - brand color for the application source of the
	// activity. Must be a valid hex color
	BackgroundColor nullable.Type[string] `json:"backgroundColor,omitempty"`

	// Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows
	// Shell UI
	Content *Json `json:"content,omitempty"`

	// Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or
	// metadata)
	Description nullable.Type[string] `json:"description,omitempty"`

	// Required. Short text description of the user's unique activity (for example, document name in cases where an activity
	// refers to document creation)
	DisplayText string `json:"displayText"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
