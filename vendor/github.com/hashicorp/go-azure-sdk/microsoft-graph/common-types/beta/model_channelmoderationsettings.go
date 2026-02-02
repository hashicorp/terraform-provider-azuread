package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelModerationSettings struct {
	// Indicates whether bots are allowed to post messages.
	AllowNewMessageFromBots nullable.Type[bool] `json:"allowNewMessageFromBots,omitempty"`

	// Indicates whether connectors are allowed to post messages.
	AllowNewMessageFromConnectors nullable.Type[bool] `json:"allowNewMessageFromConnectors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates who is allowed to reply to the teams channel. Possible values are: everyone, authorAndModerators,
	// unknownFutureValue.
	ReplyRestriction *ReplyRestriction `json:"replyRestriction,omitempty"`

	// Indicates who is allowed to post messages to teams channel. Possible values are: everyone, everyoneExceptGuests,
	// moderators, unknownFutureValue.
	UserNewMessageRestriction *UserNewMessageRestriction `json:"userNewMessageRestriction,omitempty"`
}
