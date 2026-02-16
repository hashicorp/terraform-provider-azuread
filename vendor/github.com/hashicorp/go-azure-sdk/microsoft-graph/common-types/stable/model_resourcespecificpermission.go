package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSpecificPermission struct {
	// Describes the level of access that the resource-specific permission represents.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the resource-specific permission.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The unique identifier for the resource-specific application permission.
	Id *string `json:"id,omitempty"`

	// Indicates whether the permission is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value of the permission.
	Value nullable.Type[string] `json:"value,omitempty"`
}
