package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassSettings struct {
	// Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
	IsDialInBypassEnabled nullable.Type[bool] `json:"isDialInBypassEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
	Scope *LobbyBypassScope `json:"scope,omitempty"`
}
