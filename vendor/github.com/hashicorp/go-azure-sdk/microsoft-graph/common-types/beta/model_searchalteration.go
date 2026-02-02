package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAlteration struct {
	// Defines the altered highlighted query string with spelling correction. The annotation around the corrected segment is
	// (/ue000, /ue001)
	AlteredHighlightedQueryString nullable.Type[string] `json:"alteredHighlightedQueryString,omitempty"`

	// Defines the altered query string with spelling correction.
	AlteredQueryString nullable.Type[string] `json:"alteredQueryString,omitempty"`

	// Represents changed segments with respect to original query.
	AlteredQueryTokens *[]AlteredQueryToken `json:"alteredQueryTokens,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
