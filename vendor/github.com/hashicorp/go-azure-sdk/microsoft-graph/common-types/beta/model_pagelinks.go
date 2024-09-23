package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PageLinks struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Opens the page in the OneNote native client if it's installed.
	OneNoteClientUrl *ExternalLink `json:"oneNoteClientUrl,omitempty"`

	// Opens the page in OneNote on the web.
	OneNoteWebUrl *ExternalLink `json:"oneNoteWebUrl,omitempty"`
}
