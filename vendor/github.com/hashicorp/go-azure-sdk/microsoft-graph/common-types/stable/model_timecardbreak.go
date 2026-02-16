package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardBreak struct {
	// ID of the timeCardBreak.
	BreakId nullable.Type[string] `json:"breakId,omitempty"`

	// The start event of the timeCardBreak.
	End *TimeCardEvent `json:"end,omitempty"`

	// Notes about the timeCardBreak.
	Notes *ItemBody `json:"notes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Start *TimeCardEvent `json:"start,omitempty"`
}
