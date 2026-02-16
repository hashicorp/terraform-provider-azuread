package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OidcAddressInboundClaims struct {
	// Country name.
	Country nullable.Type[string] `json:"country,omitempty"`

	// City or locality.
	Locality nullable.Type[string] `json:"locality,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Zip code or postal code.
	Postalcode nullable.Type[string] `json:"postal_code,omitempty"`

	// Country name.
	Region nullable.Type[string] `json:"region,omitempty"`

	// Full mailing address, formatted for display or use on a mailing label. This field MAY contain multiple lines,
	// separated by newlines. Newlines can be represented either as a carriage return/line feed pair ('/r/n') or as a single
	// line feed character ('/n').
	Streetaddress nullable.Type[string] `json:"street_address,omitempty"`
}
