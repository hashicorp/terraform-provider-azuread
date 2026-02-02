package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchHitsContainer struct {
	// Contains the collection of aggregations computed based on the provided aggregationOption specified in the request.
	Aggregations *[]SearchAggregation `json:"aggregations,omitempty"`

	// A collection of the search results.
	Hits *[]SearchHit `json:"hits,omitempty"`

	// Provides information if more results are available. Based on this information, you can adjust the from and size
	// properties of the searchRequest accordingly.
	MoreResultsAvailable nullable.Type[bool] `json:"moreResultsAvailable,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of results. Note this is not the number of results on the page, but the total number of results
	// satisfying the query.
	Total nullable.Type[int64] `json:"total,omitempty"`
}
