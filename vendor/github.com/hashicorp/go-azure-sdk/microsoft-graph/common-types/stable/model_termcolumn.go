package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermColumn struct {
	// Specifies whether the column allows more than one value.
	AllowMultipleValues nullable.Type[bool] `json:"allowMultipleValues,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ParentTerm *TermStoreTerm `json:"parentTerm,omitempty"`

	// Specifies whether to display the entire term path or only the term label.
	ShowFullyQualifiedName nullable.Type[bool] `json:"showFullyQualifiedName,omitempty"`

	TermSet *TermStoreSet `json:"termSet,omitempty"`
}
