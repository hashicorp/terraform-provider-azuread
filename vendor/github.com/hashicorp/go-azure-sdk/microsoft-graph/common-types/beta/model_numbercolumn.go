package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberColumn struct {
	// How many decimal places to display. See below for information about the possible values.
	DecimalPlaces nullable.Type[string] `json:"decimalPlaces,omitempty"`

	// How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
	DisplayAs nullable.Type[string] `json:"displayAs,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
