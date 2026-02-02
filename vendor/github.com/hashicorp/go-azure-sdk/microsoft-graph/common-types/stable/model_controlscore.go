package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlScore struct {
	// Control action category (Identity, Data, Device, Apps, Infrastructure).
	ControlCategory nullable.Type[string] `json:"controlCategory,omitempty"`

	// Control unique name.
	ControlName nullable.Type[string] `json:"controlName,omitempty"`

	// Description of the control.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
