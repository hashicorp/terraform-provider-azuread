package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageHistoryItem struct {
	Actions *ChatMessageActions `json:"actions,omitempty"`

	// The date and time when the message was modified.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The reaction in the modified message.
	Reaction *ChatMessageReaction `json:"reaction,omitempty"`
}
