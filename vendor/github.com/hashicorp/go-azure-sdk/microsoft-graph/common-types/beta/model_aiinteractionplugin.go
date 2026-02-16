package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AiInteractionPlugin struct {
	// The unique identifier of the plugin.
	Identifier nullable.Type[string] `json:"identifier,omitempty"`

	// The display name of the plugin.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The version of the plugin used.
	Version nullable.Type[string] `json:"version,omitempty"`
}
