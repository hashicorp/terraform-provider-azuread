package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OathTokenMetadata struct {
	Enabled                nullable.Type[bool]   `json:"enabled,omitempty"`
	Manufacturer           nullable.Type[string] `json:"manufacturer,omitempty"`
	ManufacturerProperties *[]KeyValue           `json:"manufacturerProperties,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`
	TokenType    *string               `json:"tokenType,omitempty"`
}
