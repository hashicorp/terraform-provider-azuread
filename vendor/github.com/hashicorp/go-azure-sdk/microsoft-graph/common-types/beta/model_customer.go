package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Customer struct {
	Address              *PostalAddressType    `json:"address,omitempty"`
	Blocked              nullable.Type[string] `json:"blocked,omitempty"`
	Currency             *Currency             `json:"currency,omitempty"`
	CurrencyCode         nullable.Type[string] `json:"currencyCode,omitempty"`
	CurrencyId           nullable.Type[string] `json:"currencyId,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	Email                nullable.Type[string] `json:"email,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
	Number               nullable.Type[string] `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PaymentMethod         *PaymentMethod        `json:"paymentMethod,omitempty"`
	PaymentMethodId       nullable.Type[string] `json:"paymentMethodId,omitempty"`
	PaymentTerm           *PaymentTerm          `json:"paymentTerm,omitempty"`
	PaymentTermsId        nullable.Type[string] `json:"paymentTermsId,omitempty"`
	PhoneNumber           nullable.Type[string] `json:"phoneNumber,omitempty"`
	Picture               *[]Picture            `json:"picture,omitempty"`
	ShipmentMethod        *ShipmentMethod       `json:"shipmentMethod,omitempty"`
	ShipmentMethodId      nullable.Type[string] `json:"shipmentMethodId,omitempty"`
	TaxAreaDisplayName    nullable.Type[string] `json:"taxAreaDisplayName,omitempty"`
	TaxAreaId             nullable.Type[string] `json:"taxAreaId,omitempty"`
	TaxLiable             nullable.Type[bool]   `json:"taxLiable,omitempty"`
	TaxRegistrationNumber nullable.Type[string] `json:"taxRegistrationNumber,omitempty"`
	Type                  nullable.Type[string] `json:"type,omitempty"`
	Website               nullable.Type[string] `json:"website,omitempty"`
}
