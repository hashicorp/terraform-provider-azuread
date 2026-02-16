package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesCreditMemo struct {
	BillToCustomerId         nullable.Type[string]  `json:"billToCustomerId,omitempty"`
	BillToCustomerNumber     nullable.Type[string]  `json:"billToCustomerNumber,omitempty"`
	BillToName               nullable.Type[string]  `json:"billToName,omitempty"`
	BillingPostalAddress     *PostalAddressType     `json:"billingPostalAddress,omitempty"`
	CreditMemoDate           nullable.Type[string]  `json:"creditMemoDate,omitempty"`
	Currency                 *Currency              `json:"currency,omitempty"`
	CurrencyCode             nullable.Type[string]  `json:"currencyCode,omitempty"`
	CurrencyId               nullable.Type[string]  `json:"currencyId,omitempty"`
	Customer                 *Customer              `json:"customer,omitempty"`
	CustomerId               nullable.Type[string]  `json:"customerId,omitempty"`
	CustomerName             nullable.Type[string]  `json:"customerName,omitempty"`
	CustomerNumber           nullable.Type[string]  `json:"customerNumber,omitempty"`
	DiscountAmount           nullable.Type[float64] `json:"discountAmount,omitempty"`
	DiscountAppliedBeforeTax nullable.Type[bool]    `json:"discountAppliedBeforeTax,omitempty"`
	DueDate                  nullable.Type[string]  `json:"dueDate,omitempty"`
	Email                    nullable.Type[string]  `json:"email,omitempty"`
	ExternalDocumentNumber   nullable.Type[string]  `json:"externalDocumentNumber,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	InvoiceId                nullable.Type[string]  `json:"invoiceId,omitempty"`
	InvoiceNumber            nullable.Type[string]  `json:"invoiceNumber,omitempty"`
	LastModifiedDateTime     nullable.Type[string]  `json:"lastModifiedDateTime,omitempty"`
	Number                   nullable.Type[string]  `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PaymentTerm             *PaymentTerm           `json:"paymentTerm,omitempty"`
	PaymentTermsId          nullable.Type[string]  `json:"paymentTermsId,omitempty"`
	PhoneNumber             nullable.Type[string]  `json:"phoneNumber,omitempty"`
	PricesIncludeTax        nullable.Type[bool]    `json:"pricesIncludeTax,omitempty"`
	SalesCreditMemoLines    *[]SalesCreditMemoLine `json:"salesCreditMemoLines,omitempty"`
	Salesperson             nullable.Type[string]  `json:"salesperson,omitempty"`
	SellingPostalAddress    *PostalAddressType     `json:"sellingPostalAddress,omitempty"`
	Status                  nullable.Type[string]  `json:"status,omitempty"`
	TotalAmountExcludingTax nullable.Type[float64] `json:"totalAmountExcludingTax,omitempty"`
	TotalAmountIncludingTax nullable.Type[float64] `json:"totalAmountIncludingTax,omitempty"`
	TotalTaxAmount          nullable.Type[float64] `json:"totalTaxAmount,omitempty"`
}
