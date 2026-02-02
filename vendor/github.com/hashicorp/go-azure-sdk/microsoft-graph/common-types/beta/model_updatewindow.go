package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindow struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// End of a time window during which agents can receive updates
	UpdateWindowEndTime nullable.Type[string] `json:"updateWindowEndTime,omitempty"`

	// Start of a time window during which agents can receive updates
	UpdateWindowStartTime nullable.Type[string] `json:"updateWindowStartTime,omitempty"`
}
