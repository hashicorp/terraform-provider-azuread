package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectedContent struct {
	// The content id
	Cid nullable.Type[string] `json:"cid,omitempty"`

	// The content format.
	Format nullable.Type[string] `json:"format,omitempty"`

	// The unique identifier for the sensitivity label applied to the content.
	LabelId *string `json:"labelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
