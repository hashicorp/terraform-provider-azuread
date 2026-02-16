package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditProperty struct {
	// The display name for this property.
	DisplayName *string `json:"displayName,omitempty"`

	// The new value for this property.
	NewValue *string `json:"newValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The old value for this property.
	OldValue nullable.Type[string] `json:"oldValue,omitempty"`
}
