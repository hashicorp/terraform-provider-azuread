package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentItem struct {
	// The type of attachment. Possible values are: file, item, reference. Required.
	AttachmentType AttachmentType `json:"attachmentType"`

	// The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img
	// src='cid:contentId'> tag in HTML messages. Optional.
	ContentId nullable.Type[string] `json:"contentId,omitempty"`

	// The nature of the data in the attachment. Optional.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// true if the attachment is an inline attachment; otherwise, false. Optional.
	IsInline nullable.Type[bool] `json:"isInline,omitempty"`

	// The display name of the attachment. This can be a descriptive string and doesn't have to be the actual file name.
	// Required.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The length of the attachment in bytes. Required.
	Size nullable.Type[int64] `json:"size,omitempty"`
}
