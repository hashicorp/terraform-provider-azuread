package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalOfficeAddress struct {
	// The city.
	City nullable.Type[string] `json:"city,omitempty"`

	// The country or region. It's a free-format string value, for example, 'United States'.
	CountryOrRegion nullable.Type[string] `json:"countryOrRegion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Office location such as building and office number for an organizational contact.
	OfficeLocation nullable.Type[string] `json:"officeLocation,omitempty"`

	// The postal code.
	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`

	// The state.
	State nullable.Type[string] `json:"state,omitempty"`

	// The street.
	Street nullable.Type[string] `json:"street,omitempty"`
}
