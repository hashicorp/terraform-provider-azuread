package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTypeOrder struct {
	// Indicates whether this is the default content type
	Default nullable.Type[bool] `json:"default,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the position in which the content type appears in the selection UI.
	Position nullable.Type[int64] `json:"position,omitempty"`
}
