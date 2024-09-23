package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SalesCreditMemoLine{}

type SalesCreditMemoLine struct {
	Account                   *Account               `json:"account,omitempty"`
	AccountId                 nullable.Type[string]  `json:"accountId,omitempty"`
	AmountExcludingTax        nullable.Type[float64] `json:"amountExcludingTax,omitempty"`
	AmountIncludingTax        nullable.Type[float64] `json:"amountIncludingTax,omitempty"`
	Description               nullable.Type[string]  `json:"description,omitempty"`
	DiscountAmount            nullable.Type[float64] `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax  nullable.Type[bool]    `json:"discountAppliedBeforeTax,omitempty"`
	DiscountPercent           nullable.Type[float64] `json:"discountPercent,omitempty"`
	DocumentId                nullable.Type[string]  `json:"documentId,omitempty"`
	InvoiceDiscountAllocation nullable.Type[float64] `json:"invoiceDiscountAllocation,omitempty"`
	Item                      *Item                  `json:"item,omitempty"`
	ItemId                    nullable.Type[string]  `json:"itemId,omitempty"`
	LineType                  nullable.Type[string]  `json:"lineType,omitempty"`
	NetAmount                 nullable.Type[float64] `json:"netAmount,omitempty"`
	NetAmountIncludingTax     nullable.Type[float64] `json:"netAmountIncludingTax,omitempty"`
	NetTaxAmount              nullable.Type[float64] `json:"netTaxAmount,omitempty"`
	Quantity                  nullable.Type[float64] `json:"quantity,omitempty"`
	Sequence                  nullable.Type[int64]   `json:"sequence,omitempty"`
	ShipmentDate              nullable.Type[string]  `json:"shipmentDate,omitempty"`
	TaxCode                   nullable.Type[string]  `json:"taxCode,omitempty"`
	TaxPercent                nullable.Type[float64] `json:"taxPercent,omitempty"`
	TotalTaxAmount            nullable.Type[float64] `json:"totalTaxAmount,omitempty"`
	UnitOfMeasureId           nullable.Type[string]  `json:"unitOfMeasureId,omitempty"`
	UnitPrice                 nullable.Type[float64] `json:"unitPrice,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SalesCreditMemoLine) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SalesCreditMemoLine{}

func (s SalesCreditMemoLine) MarshalJSON() ([]byte, error) {
	type wrapper SalesCreditMemoLine
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SalesCreditMemoLine: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SalesCreditMemoLine: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.salesCreditMemoLine"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SalesCreditMemoLine: %+v", err)
	}

	return encoded, nil
}
