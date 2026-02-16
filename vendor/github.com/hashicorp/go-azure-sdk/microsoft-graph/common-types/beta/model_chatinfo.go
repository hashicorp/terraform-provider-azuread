package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInfo struct {
	// The unique identifier for a message in a Microsoft Teams channel.
	MessageId nullable.Type[string] `json:"messageId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the reply message.
	ReplyChainMessageId nullable.Type[string] `json:"replyChainMessageId,omitempty"`

	// The unique identifier for a thread in Microsoft Teams.
	ThreadId nullable.Type[string] `json:"threadId,omitempty"`
}
