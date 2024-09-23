package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedAction struct {
	// Web URL to the recommended action.
	ActionWebUrl nullable.Type[string] `json:"actionWebUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Title of the recommended action.
	Title nullable.Type[string] `json:"title,omitempty"`
}
