package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Property struct {
	Aliases       *[]string           `json:"aliases,omitempty"`
	IsQueryable   nullable.Type[bool] `json:"isQueryable,omitempty"`
	IsRefinable   nullable.Type[bool] `json:"isRefinable,omitempty"`
	IsRetrievable nullable.Type[bool] `json:"isRetrievable,omitempty"`
	IsSearchable  nullable.Type[bool] `json:"isSearchable,omitempty"`
	Labels        *[]Label            `json:"labels,omitempty"`
	Name          *string             `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Type *PropertyType `json:"type,omitempty"`
}
