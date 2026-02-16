package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointSearchResponse struct {
	// Contains alist of protection units with no restore points.
	NoResultProtectionUnitIds *[]string `json:"noResultProtectionUnitIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of the search response.
	SearchResponseId nullable.Type[string] `json:"searchResponseId,omitempty"`

	// Contains a collection of restore points.
	SearchResults *[]RestorePointSearchResult `json:"searchResults,omitempty"`
}
