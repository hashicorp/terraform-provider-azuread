package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentInfo struct {
	// The type of the attachment. The possible values are: file, item, reference. Required.
	AttachmentType AttachmentType `json:"attachmentType"`

	// The nature of the data in the attachment. Optional.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

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
