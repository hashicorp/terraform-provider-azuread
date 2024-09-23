package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PurchaseInvoice struct {
	BuyFromAddress           *PostalAddressType     `json:"buyFromAddress,omitempty"`
	Currency                 *Currency              `json:"currency,omitempty"`
	CurrencyCode             nullable.Type[string]  `json:"currencyCode,omitempty"`
	CurrencyId               nullable.Type[string]  `json:"currencyId,omitempty"`
	DiscountAmount           nullable.Type[float64] `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax nullable.Type[bool]    `json:"discountAppliedBeforeTax,omitempty"`
	DueDate                  nullable.Type[string]  `json:"dueDate,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	InvoiceDate              nullable.Type[string]  `json:"invoiceDate,omitempty"`
	LastModifiedDateTime     nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
	Number                   nullable.Type[string]  `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PayToAddress            *PostalAddressType     `json:"payToAddress,omitempty"`
	PayToContact            nullable.Type[string]  `json:"payToContact,omitempty"`
	PayToName               nullable.Type[string]  `json:"payToName,omitempty"`
	PayToVendorId           nullable.Type[string]  `json:"payToVendorId,omitempty"`
	PayToVendorNumber       nullable.Type[string]  `json:"payToVendorNumber,omitempty"`
	PricesIncludeTax        nullable.Type[bool]    `json:"pricesIncludeTax,omitempty"`
	PurchaseInvoiceLines    *[]PurchaseInvoiceLine `json:"purchaseInvoiceLines,omitempty"`
	ShipToAddress           *PostalAddressType     `json:"shipToAddress,omitempty"`
	ShipToContact           nullable.Type[string]  `json:"shipToContact,omitempty"`
	ShipToName              nullable.Type[string]  `json:"shipToName,omitempty"`
	Status                  nullable.Type[string]  `json:"status,omitempty"`
	TotalAmountExcludingTax nullable.Type[float64] `json:"totalAmountExcludingTax,omitempty"`
	TotalAmountIncludingTax nullable.Type[float64] `json:"totalAmountIncludingTax,omitempty"`
	TotalTaxAmount          nullable.Type[float64] `json:"totalTaxAmount,omitempty"`
	Vendor                  *Vendor                `json:"vendor,omitempty"`
	VendorId                nullable.Type[string]  `json:"vendorId,omitempty"`
	VendorInvoiceNumber     nullable.Type[string]  `json:"vendorInvoiceNumber,omitempty"`
	VendorName              nullable.Type[string]  `json:"vendorName,omitempty"`
	VendorNumber            nullable.Type[string]  `json:"vendorNumber,omitempty"`
}
