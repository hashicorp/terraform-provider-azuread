package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScoredEmailAddress struct {
	// The email address.
	Address nullable.Type[string] `json:"address,omitempty"`

	ItemId nullable.Type[string] `json:"itemId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SelectionLikelihood *SelectionLikelihoodInfo `json:"selectionLikelihood,omitempty"`
}
