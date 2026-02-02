package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlterationResponse struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the original user query string.
	OriginalQueryString nullable.Type[string] `json:"originalQueryString,omitempty"`

	// Defines the details of alteration information for the spelling correction.
	QueryAlteration *SearchAlteration `json:"queryAlteration,omitempty"`

	// Defines the type of the spelling correction. Possible values are suggestion, modification.
	QueryAlterationType *SearchAlterationType `json:"queryAlterationType,omitempty"`
}
