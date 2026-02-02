package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAlterationOptions struct {
	// Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected
	// query when there are no results for the original query with typos and get the spelling modification information in
	// queryAlterationResponse property of the response. Optional.
	EnableModification nullable.Type[bool] `json:"enableModification,omitempty"`

	// Indicates whether spelling suggestions are enabled. If enabled, the user will get the search results for the original
	// search query and suggestions for spelling correction in the queryAlterationResponse property of the response for the
	// typos in the query. Optional.
	EnableSuggestion nullable.Type[bool] `json:"enableSuggestion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
