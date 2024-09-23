package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileStorageContainerCustomPropertyValue struct {
	// Indicates whether the custom property is searchable. Optional. The default value is false.
	IsSearchable nullable.Type[bool] `json:"isSearchable,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value of the custom property. Required.
	Value string `json:"value"`
}
