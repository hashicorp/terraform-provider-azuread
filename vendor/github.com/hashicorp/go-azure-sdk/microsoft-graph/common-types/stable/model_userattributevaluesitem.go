package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAttributeValuesItem struct {
	// Determines whether the value is set as the default.
	IsDefault *bool `json:"isDefault,omitempty"`

	// The display name of the property displayed to the user in the user flow.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value that is set when this item is selected.
	Value nullable.Type[string] `json:"value,omitempty"`
}
