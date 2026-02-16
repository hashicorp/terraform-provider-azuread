package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionPageViewConfiguration struct {
	// The description of the page.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display configuration of attributes being collected on the attribute collection page. You must specify all
	// attributes that you want to retain, otherwise they're removed from the user flow.
	Inputs *[]AuthenticationAttributeCollectionInputConfiguration `json:"inputs,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The title of the attribute collection page.
	Title nullable.Type[string] `json:"title,omitempty"`
}
