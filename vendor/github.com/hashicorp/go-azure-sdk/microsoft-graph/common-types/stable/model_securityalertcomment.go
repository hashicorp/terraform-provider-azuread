package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertComment struct {
	// The comment text.
	Comment nullable.Type[string] `json:"comment,omitempty"`

	// The person or app name that submitted the comment.
	CreatedByDisplayName nullable.Type[string] `json:"createdByDisplayName,omitempty"`

	// The time when the comment was submitted.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
