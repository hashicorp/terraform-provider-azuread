package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactDataMatchStoreColumn struct {
	IgnoredDelimiters *[]string             `json:"ignoredDelimiters,omitempty"`
	IsCaseInsensitive nullable.Type[bool]   `json:"isCaseInsensitive,omitempty"`
	IsSearchable      nullable.Type[bool]   `json:"isSearchable,omitempty"`
	Name              nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
