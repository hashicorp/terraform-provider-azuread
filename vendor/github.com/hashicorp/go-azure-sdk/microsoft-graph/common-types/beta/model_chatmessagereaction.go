package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageReaction struct {
	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The hosted content URL for the custom reaction type.
	ReactionContentUrl nullable.Type[string] `json:"reactionContentUrl,omitempty"`

	// Supported values are Unicode characters and custom. Some backward-compatible reaction types include like, angry, sad,
	// laugh, heart, and surprised.
	ReactionType *string `json:"reactionType,omitempty"`

	User *ChatMessageReactionIdentitySet `json:"user,omitempty"`
}
