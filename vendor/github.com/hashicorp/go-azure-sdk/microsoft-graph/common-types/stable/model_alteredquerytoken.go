package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlteredQueryToken struct {
	// Defines the length of a changed segment.
	Length nullable.Type[int64] `json:"length,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the offset of a changed segment.
	Offset nullable.Type[int64] `json:"offset,omitempty"`

	// Represents the corrected segment string.
	Suggestion nullable.Type[string] `json:"suggestion,omitempty"`
}
