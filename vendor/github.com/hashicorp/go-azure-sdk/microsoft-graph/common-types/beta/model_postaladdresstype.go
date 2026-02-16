package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PostalAddressType struct {
	City              nullable.Type[string] `json:"city,omitempty"`
	CountryLetterCode nullable.Type[string] `json:"countryLetterCode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PostalCode nullable.Type[string] `json:"postalCode,omitempty"`
	State      nullable.Type[string] `json:"state,omitempty"`
	Street     nullable.Type[string] `json:"street,omitempty"`
}
