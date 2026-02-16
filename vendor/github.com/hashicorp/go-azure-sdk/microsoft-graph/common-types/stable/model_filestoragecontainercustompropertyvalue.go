package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerCustomPropertyValue struct {
	// Indicates whether the custom property is searchable. Optional. The default value is false.
	IsSearchable nullable.Type[bool] `json:"isSearchable,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the custom property. Required.
	Value string `json:"value"`
}
