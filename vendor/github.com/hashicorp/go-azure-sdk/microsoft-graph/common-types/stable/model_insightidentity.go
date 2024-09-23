package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightIdentity struct {
	// The email address of the user who shared the item.
	Address nullable.Type[string] `json:"address,omitempty"`

	// The display name of the user who shared the item.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The ID of the user who shared the item.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
