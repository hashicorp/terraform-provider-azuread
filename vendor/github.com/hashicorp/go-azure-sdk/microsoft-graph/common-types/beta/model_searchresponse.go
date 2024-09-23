package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchResponse struct {
	// A collection of search results.
	HitsContainers *[]SearchHitsContainer `json:"hitsContainers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Provides details of query alteration response for spelling correction.
	QueryAlterationResponse *AlterationResponse `json:"queryAlterationResponse,omitempty"`

	// A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result
	// templates.
	ResultTemplates *ResultTemplateDictionary `json:"resultTemplates,omitempty"`

	// Contains the search terms sent in the initial search query.
	SearchTerms *[]string `json:"searchTerms,omitempty"`
}
