package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMention struct {
	// Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at
	// id='{index}'> tag in the message body.
	Id nullable.Type[int64] `json:"id,omitempty"`

	// String used to represent the mention. For example, a user's display name, a team name.
	MentionText nullable.Type[string] `json:"mentionText,omitempty"`

	// The entity (user, application, team, channel, or chat) that was @mentioned.
	Mentioned *ChatMessageMentionedIdentitySet `json:"mentioned,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
