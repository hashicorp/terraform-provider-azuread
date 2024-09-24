package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSpecificPermission struct {
	Description nullable.Type[string] `json:"description,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	Id          *string               `json:"id,omitempty"`
	IsEnabled   *bool                 `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Value nullable.Type[string] `json:"value,omitempty"`
}
