package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsProperty struct {
	// A set of aliases or a friendly name for the property. Maximum 32 characters. Only alphanumeric characters allowed.
	// For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [,
	// ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
	Aliases *[]string `json:"aliases,omitempty"`

	// Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries.
	// Optional.
	IsQueryable nullable.Type[bool] `json:"isQueryable,omitempty"`

	// Specifies if the property is refinable. Refinable properties can be used to filter search results in the Search API
	// and add a refiner control in the Microsoft Search user experience. Optional.
	IsRefinable nullable.Type[bool] `json:"isRefinable,omitempty"`

	// Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are
	// returned by the search API. Retrievable properties are also available to add to the display template used to render
	// search results. Optional.
	IsRetrievable nullable.Type[bool] `json:"isRetrievable,omitempty"`

	// Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable.
	// Nonsearchable properties aren't added to the search index. Optional.
	IsSearchable nullable.Type[bool] `json:"isSearchable,omitempty"`

	// Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics
	// of the data in the connection. Adding appropriate labels would result in an enhanced search experience (for example,
	// better relevance). Optional.The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime,
	// lastModifiedDateTime, fileName, fileExtension, unknownFutureValue, iconUrl. You must use the Prefer:
	// include-unknown-enum-members request header to get the following value in this evolvable enum: iconUrl.
	Labels *[]ExternalConnectorsLabel `json:"labels,omitempty"`

	// The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may
	// not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &,
	// ?, @, #, /, ~, ', ', <, >, `, ^. Required.
	Name string `json:"name"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Type *ExternalConnectorsPropertyType `json:"type,omitempty"`
}
