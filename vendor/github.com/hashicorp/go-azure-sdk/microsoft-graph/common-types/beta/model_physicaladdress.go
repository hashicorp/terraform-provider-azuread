package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalAddress struct {
	// The city.
	City nullable.Type[string] `json:"city,omitempty"`

	// The country or region. It's a free-format string value, for example, 'United States'.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The post office box number.
	PostOfficeBox nullable.Type[string] `json:"postOfficeBox,omitempty"`

	// The postal code.
	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`

	// The state.
	State nullable.Type[string] `json:"state,omitempty"`

	// The street.
	Street nullable.Type[string] `json:"street,omitempty"`

	// The type of address. Possible values are: unknown, home, business, other.
	Type *PhysicalAddressType `json:"type,omitempty"`
}
