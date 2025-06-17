package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEvent struct {
	// Indicates whether this action happens at an approved location. This property will be removed by November 20, 2027.
	// Use isAtApprovedLocation instead. atApprovedLocation and isAtApprovedLocation always have the same value, so setting
	// one automatically sets the value for the other. If both are included in the request with different values, the value
	// for isAtApprovedLocation takes precedence.
	AtApprovedLocation nullable.Type[bool] `json:"atApprovedLocation,omitempty"`

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
