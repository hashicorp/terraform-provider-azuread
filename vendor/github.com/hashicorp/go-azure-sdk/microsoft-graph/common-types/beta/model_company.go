package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Company struct {
	Accounts                *[]Account                `json:"accounts,omitempty"`
	AgedAccountsPayable     *[]AgedAccountsPayable    `json:"agedAccountsPayable,omitempty"`
	AgedAccountsReceivable  *[]AgedAccountsReceivable `json:"agedAccountsReceivable,omitempty"`
	BusinessProfileId       nullable.Type[string]     `json:"businessProfileId,omitempty"`
	CompanyInformation      *[]CompanyInformation     `json:"companyInformation,omitempty"`
	CountriesRegions        *[]CountryRegion          `json:"countriesRegions,omitempty"`
	Currencies              *[]Currency               `json:"currencies,omitempty"`
	CustomerPaymentJournals *[]CustomerPaymentJournal `json:"customerPaymentJournals,omitempty"`
	CustomerPayments        *[]CustomerPayment        `json:"customerPayments,omitempty"`
	Customers               *[]Customer               `json:"customers,omitempty"`
	DimensionValues         *[]DimensionValue         `json:"dimensionValues,omitempty"`
	Dimensions              *[]Dimension              `json:"dimensions,omitempty"`
	DisplayName             nullable.Type[string]     `json:"displayName,omitempty"`
	Employees               *[]Employee               `json:"employees,omitempty"`
	GeneralLedgerEntries    *[]GeneralLedgerEntry     `json:"generalLedgerEntries,omitempty"`
	Id                      *string                   `json:"id,omitempty"`
	ItemCategories          *[]ItemCategory           `json:"itemCategories,omitempty"`
	Items                   *[]Item                   `json:"items,omitempty"`
	JournalLines            *[]JournalLine            `json:"journalLines,omitempty"`
	Journals                *[]Journal                `json:"journals,omitempty"`
	Name                    nullable.Type[string]     `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PaymentMethods       *[]PaymentMethod       `json:"paymentMethods,omitempty"`
	PaymentTerms         *[]PaymentTerm         `json:"paymentTerms,omitempty"`
	Picture              *[]Picture             `json:"picture,omitempty"`
	PurchaseInvoiceLines *[]PurchaseInvoiceLine `json:"purchaseInvoiceLines,omitempty"`
	PurchaseInvoices     *[]PurchaseInvoice     `json:"purchaseInvoices,omitempty"`
	SalesCreditMemoLines *[]SalesCreditMemoLine `json:"salesCreditMemoLines,omitempty"`
	SalesCreditMemos     *[]SalesCreditMemo     `json:"salesCreditMemos,omitempty"`
	SalesInvoiceLines    *[]SalesInvoiceLine    `json:"salesInvoiceLines,omitempty"`
	SalesInvoices        *[]SalesInvoice        `json:"salesInvoices,omitempty"`
	SalesOrderLines      *[]SalesOrderLine      `json:"salesOrderLines,omitempty"`
	SalesOrders          *[]SalesOrder          `json:"salesOrders,omitempty"`
	SalesQuoteLines      *[]SalesQuoteLine      `json:"salesQuoteLines,omitempty"`
	SalesQuotes          *[]SalesQuote          `json:"salesQuotes,omitempty"`
	ShipmentMethods      *[]ShipmentMethod      `json:"shipmentMethods,omitempty"`
	SystemVersion        nullable.Type[string]  `json:"systemVersion,omitempty"`
	TaxAreas             *[]TaxArea             `json:"taxAreas,omitempty"`
	TaxGroups            *[]TaxGroup            `json:"taxGroups,omitempty"`
	UnitsOfMeasure       *[]UnitOfMeasure       `json:"unitsOfMeasure,omitempty"`
	Vendors              *[]Vendor              `json:"vendors,omitempty"`
}
