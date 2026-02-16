package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEvent struct {
	// The time the entry is recorded.
	DateTime *string `json:"dateTime,omitempty"`

	// Indicates whether this action happens at an approved location.
	IsAtApprovedLocation nullable.Type[bool] `json:"isAtApprovedLocation,omitempty"`

	// Notes about the timeCardEvent.
	Notes *ItemBody `json:"notes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
