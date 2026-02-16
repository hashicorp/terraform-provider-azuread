package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveContentLocation struct {
	Confidence nullable.Type[int64]        `json:"confidence,omitempty"`
	Evidences  *[]SensitiveContentEvidence `json:"evidences,omitempty"`
	IdMatch    nullable.Type[string]       `json:"idMatch,omitempty"`
	Length     nullable.Type[int64]        `json:"length,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Offset nullable.Type[int64] `json:"offset,omitempty"`
}
