package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SortProperty struct {
	// True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
	IsDescending nullable.Type[bool] `json:"isDescending,omitempty"`

	// The name of the property to sort on. Required.
	Name string `json:"name"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
