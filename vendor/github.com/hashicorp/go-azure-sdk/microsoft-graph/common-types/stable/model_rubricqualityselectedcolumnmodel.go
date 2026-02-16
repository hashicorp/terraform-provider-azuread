package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQualitySelectedColumnModel struct {
	// ID of the selected level for this quality.
	ColumnId nullable.Type[string] `json:"columnId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// ID of the associated quality.
	QualityId nullable.Type[string] `json:"qualityId,omitempty"`
}
