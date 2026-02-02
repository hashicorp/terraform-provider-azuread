package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionStep struct {
	// A link to the documentation or Microsoft Entra admin center page that is associated with the action step.
	ActionUrl *ActionUrl `json:"actionUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the position for this action in the order of the collection of actions to be taken.
	StepNumber *int64 `json:"stepNumber,omitempty"`

	// Friendly description of the action to take.
	Text nullable.Type[string] `json:"text,omitempty"`
}
