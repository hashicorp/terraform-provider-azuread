package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDocumentUploadProperties struct {
	// The document's content (MIME) type.
	ContentType *string `json:"contentType,omitempty"`

	// The document's name.
	DocumentName *string `json:"documentName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The document's size in bytes.
	Size *int64 `json:"size,omitempty"`
}
