package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentClassification struct {
	Confidence nullable.Type[int64] `json:"confidence,omitempty"`
	Matches    *[]MatchLocation     `json:"matches,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SensitiveTypeId nullable.Type[string] `json:"sensitiveTypeId,omitempty"`
	UniqueCount     nullable.Type[int64]  `json:"uniqueCount,omitempty"`
}
