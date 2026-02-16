package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatViewpoint struct {
	// Indicates whether the chat is hidden for the current user.
	IsHidden nullable.Type[bool] `json:"isHidden,omitempty"`

	// Represents the dateTime up until which the current user has read chatMessages in a specific chat.
	LastMessageReadDateTime nullable.Type[string] `json:"lastMessageReadDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
