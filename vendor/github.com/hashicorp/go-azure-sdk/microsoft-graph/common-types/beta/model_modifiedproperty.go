package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModifiedProperty struct {
	// Name of property that was modified.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// New property value.
	NewValue nullable.Type[string] `json:"newValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Old property value.
	OldValue nullable.Type[string] `json:"oldValue,omitempty"`
}
