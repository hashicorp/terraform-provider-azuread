package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerProcessedContent struct {
	// A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat
	// the values as HTML content and run services like safety checks, search index and link fixup on them.
	HtmlStrings *[]MetaDataKeyStringPair `json:"htmlStrings,omitempty"`

	// A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values
	// as image sources and run services like search index and link fixup on them.
	ImageSources *[]MetaDataKeyStringPair `json:"imageSources,omitempty"`

	// A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links
	// and run services like link fixup on them.
	Links *[]MetaDataKeyStringPair `json:"links,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A key-value map where keys are string identifiers and values are strings that should be search indexed.
	SearchablePlainTexts *[]MetaDataKeyStringPair `json:"searchablePlainTexts,omitempty"`
}
