package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModifiedProperty struct {
	// Indicates the property name of the target attribute that was changed.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates the updated value for the propery.
	NewValue nullable.Type[string] `json:"newValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the previous value (before the update) for the property.
	OldValue nullable.Type[string] `json:"oldValue,omitempty"`
}
