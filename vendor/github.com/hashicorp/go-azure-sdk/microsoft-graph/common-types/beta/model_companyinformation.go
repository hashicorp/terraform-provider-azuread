package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyInformation struct {
	Address                    *PostalAddressType    `json:"address,omitempty"`
	CurrencyCode               nullable.Type[string] `json:"currencyCode,omitempty"`
	CurrentFiscalYearStartDate nullable.Type[string] `json:"currentFiscalYearStartDate,omitempty"`
	DisplayName                nullable.Type[string] `json:"displayName,omitempty"`
	Email                      nullable.Type[string] `json:"email,omitempty"`
	FaxNumber                  nullable.Type[string] `json:"faxNumber,omitempty"`
	Id                         *string               `json:"id,omitempty"`
	Industry                   nullable.Type[string] `json:"industry,omitempty"`
	LastModifiedDateTime       nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PhoneNumber           nullable.Type[string] `json:"phoneNumber,omitempty"`
	Picture               nullable.Type[string] `json:"picture,omitempty"`
	TaxRegistrationNumber nullable.Type[string] `json:"taxRegistrationNumber,omitempty"`
	Website               nullable.Type[string] `json:"website,omitempty"`
}
