package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainRegistrant struct {
	CountryOrRegionCode nullable.Type[string] `json:"countryOrRegionCode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Organization nullable.Type[string] `json:"organization,omitempty"`
	Url          nullable.Type[string] `json:"url,omitempty"`
	Vendor       nullable.Type[string] `json:"vendor,omitempty"`
}
