package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Item struct {
	BaseUnitOfMeasureId  nullable.Type[string]  `json:"baseUnitOfMeasureId,omitempty"`
	Blocked              nullable.Type[bool]    `json:"blocked,omitempty"`
	DisplayName          nullable.Type[string]  `json:"displayName,omitempty"`
	Gtin                 nullable.Type[string]  `json:"gtin,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Inventory            nullable.Type[float64] `json:"inventory,omitempty"`
	ItemCategory         *ItemCategory          `json:"itemCategory,omitempty"`
	ItemCategoryCode     nullable.Type[string]  `json:"itemCategoryCode,omitempty"`
	ItemCategoryId       nullable.Type[string]  `json:"itemCategoryId,omitempty"`
	LastModifiedDateTime nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
	Number               nullable.Type[string]  `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Picture          *[]Picture             `json:"picture,omitempty"`
	PriceIncludesTax nullable.Type[bool]    `json:"priceIncludesTax,omitempty"`
	TaxGroupCode     nullable.Type[string]  `json:"taxGroupCode,omitempty"`
	TaxGroupId       nullable.Type[string]  `json:"taxGroupId,omitempty"`
	Type             nullable.Type[string]  `json:"type,omitempty"`
	UnitCost         nullable.Type[float64] `json:"unitCost,omitempty"`
	UnitPrice        nullable.Type[float64] `json:"unitPrice,omitempty"`
}
