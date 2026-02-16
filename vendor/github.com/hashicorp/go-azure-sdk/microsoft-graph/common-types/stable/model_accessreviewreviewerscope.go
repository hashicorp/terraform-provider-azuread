package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewReviewerScope struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The query specifying who will be the reviewer.
	Query nullable.Type[string] `json:"query,omitempty"`

	// In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative
	// source of the query. This property is only required if a relative query, for example, ./manager, is specified.
	// Possible value: decisions.
	QueryRoot nullable.Type[string] `json:"queryRoot,omitempty"`

	// The type of query. Examples include MicrosoftGraph and ARM.
	QueryType nullable.Type[string] `json:"queryType,omitempty"`
}
