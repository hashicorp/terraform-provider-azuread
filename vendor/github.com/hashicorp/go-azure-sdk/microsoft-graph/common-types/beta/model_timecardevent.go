package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardEvent struct {
	// Indicates whether the entry was recorded at the approved location.
	AtApprovedLocation nullable.Type[bool] `json:"atApprovedLocation,omitempty"`

	// The time the entry is recorded.
	DateTime *string `json:"dateTime,omitempty"`

	// Notes about the timeCardEvent.
	Notes *ItemBody `json:"notes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
