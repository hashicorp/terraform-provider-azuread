package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPReferenceData struct {
	Asn                 nullable.Type[int64]  `json:"asn,omitempty"`
	City                nullable.Type[string] `json:"city,omitempty"`
	CountryOrRegionCode nullable.Type[string] `json:"countryOrRegionCode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Organization nullable.Type[string] `json:"organization,omitempty"`
	State        nullable.Type[string] `json:"state,omitempty"`
	Vendor       nullable.Type[string] `json:"vendor,omitempty"`
}
