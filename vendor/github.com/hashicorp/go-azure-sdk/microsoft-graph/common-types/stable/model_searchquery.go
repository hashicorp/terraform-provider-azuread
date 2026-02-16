package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQuery struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The search query containing the search terms. Required.
	QueryString string `json:"queryString"`

	// Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
	QueryTemplate nullable.Type[string] `json:"queryTemplate,omitempty"`
}
