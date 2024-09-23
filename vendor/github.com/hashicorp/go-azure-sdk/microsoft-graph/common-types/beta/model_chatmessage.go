package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ChatMessage{}

type ChatMessage struct {
	// References to attached objects like files, tabs, meetings etc.
	Attachments *[]ChatMessageAttachment `json:"attachments,omitempty"`

	Body *ItemBody `json:"body,omitempty"`

	// If the message was sent in a channel, represents identity of the channel.
	ChannelIdentity *ChannelIdentity `json:"channelIdentity,omitempty"`

	// If the message was sent in a chat, represents the identity of the chat.
	ChatId nullable.Type[string] `json:"chatId,omitempty"`

	// Timestamp of when the chat message was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Read only. Timestamp at which the chat message was deleted, or null if not deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// Read-only. Version number of the chat message.
	Etag nullable.Type[string] `json:"etag,omitempty"`

	// Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example,
	// adding new members. For event messages, the messageType property will be set to systemEventMessage.
	EventDetail *EventMessageDetail `json:"eventDetail,omitempty"`

	// Details of the sender of the chat message. Can only be set during migration.
	From *ChatMessageFromIdentitySet `json:"from,omitempty"`

	// Content in a message hosted by Microsoft Teams - for example, images or code snippets.
	HostedContents *[]ChatMessageHostedContent `json:"hostedContents,omitempty"`

	Importance *ChatMessageImportance `json:"importance,omitempty"`

	// Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits
	// are made the value is null.
	LastEditedDateTime nullable.Type[string] `json:"lastEditedDateTime,omitempty"`

	// Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is
	// added or removed.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Locale of the chat message set by the client. Always set to en-us.
	Locale *string `json:"locale,omitempty"`

	// List of entities mentioned in the chat message. Supported entities are: user, bot, team, channel, and tag.
	Mentions *[]ChatMessageMention `json:"mentions,omitempty"`

	// List of activity history of a message item, including modification time and actions, such as reactionAdded,
	// reactionRemoved, or reaction changes, on the message.
	MessageHistory *[]ChatMessageHistoryItem `json:"messageHistory,omitempty"`

	MessageType *ChatMessageType `json:"messageType,omitempty"`

	// User attribution of the message when bot sends a message on behalf of a user.
	OnBehalfOf *ChatMessageFromIdentitySet `json:"onBehalfOf,omitempty"`

	// Defines the properties of a policy violation set by a data loss prevention (DLP) application.
	PolicyViolation *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`

	// Reactions for this chat message (for example, Like).
	Reactions *[]ChatMessageReaction `json:"reactions,omitempty"`

	// Replies for a specified message. Supports $expand for channel messages.
	Replies *[]ChatMessage `json:"replies,omitempty"`

	// Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in
	// channels, not chats.)
	ReplyToId nullable.Type[string] `json:"replyToId,omitempty"`

	// The subject of the chat message, in plaintext.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only
	// applies to channel chat messages, not chat messages in a chat.
	Summary nullable.Type[string] `json:"summary,omitempty"`

	// Read-only. Link to the message in Microsoft Teams.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChatMessage) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChatMessage{}

func (s ChatMessage) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessage: %+v", err)
	}

	delete(decoded, "etag")
	delete(decoded, "eventDetail")
	delete(decoded, "replyToId")
	delete(decoded, "webUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMessage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessage: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChatMessage{}

func (s *ChatMessage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Attachments          *[]ChatMessageAttachment    `json:"attachments,omitempty"`
		Body                 *ItemBody                   `json:"body,omitempty"`
		ChannelIdentity      *ChannelIdentity            `json:"channelIdentity,omitempty"`
		ChatId               nullable.Type[string]       `json:"chatId,omitempty"`
		CreatedDateTime      nullable.Type[string]       `json:"createdDateTime,omitempty"`
		DeletedDateTime      nullable.Type[string]       `json:"deletedDateTime,omitempty"`
		Etag                 nullable.Type[string]       `json:"etag,omitempty"`
		From                 *ChatMessageFromIdentitySet `json:"from,omitempty"`
		HostedContents       *[]ChatMessageHostedContent `json:"hostedContents,omitempty"`
		Importance           *ChatMessageImportance      `json:"importance,omitempty"`
		LastEditedDateTime   nullable.Type[string]       `json:"lastEditedDateTime,omitempty"`
		LastModifiedDateTime nullable.Type[string]       `json:"lastModifiedDateTime,omitempty"`
		Locale               *string                     `json:"locale,omitempty"`
		Mentions             *[]ChatMessageMention       `json:"mentions,omitempty"`
		MessageHistory       *[]ChatMessageHistoryItem   `json:"messageHistory,omitempty"`
		MessageType          *ChatMessageType            `json:"messageType,omitempty"`
		OnBehalfOf           *ChatMessageFromIdentitySet `json:"onBehalfOf,omitempty"`
		PolicyViolation      *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`
		Reactions            *[]ChatMessageReaction      `json:"reactions,omitempty"`
		Replies              *[]ChatMessage              `json:"replies,omitempty"`
		ReplyToId            nullable.Type[string]       `json:"replyToId,omitempty"`
		Subject              nullable.Type[string]       `json:"subject,omitempty"`
		Summary              nullable.Type[string]       `json:"summary,omitempty"`
		WebUrl               nullable.Type[string]       `json:"webUrl,omitempty"`
		Id                   *string                     `json:"id,omitempty"`
		ODataId              *string                     `json:"@odata.id,omitempty"`
		ODataType            *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Attachments = decoded.Attachments
	s.Body = decoded.Body
	s.ChannelIdentity = decoded.ChannelIdentity
	s.ChatId = decoded.ChatId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Etag = decoded.Etag
	s.From = decoded.From
	s.HostedContents = decoded.HostedContents
	s.Importance = decoded.Importance
	s.LastEditedDateTime = decoded.LastEditedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Locale = decoded.Locale
	s.Mentions = decoded.Mentions
	s.MessageHistory = decoded.MessageHistory
	s.MessageType = decoded.MessageType
	s.OnBehalfOf = decoded.OnBehalfOf
	s.PolicyViolation = decoded.PolicyViolation
	s.Reactions = decoded.Reactions
	s.Replies = decoded.Replies
	s.ReplyToId = decoded.ReplyToId
	s.Subject = decoded.Subject
	s.Summary = decoded.Summary
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChatMessage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["eventDetail"]; ok {
		impl, err := UnmarshalEventMessageDetailImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EventDetail' for 'ChatMessage': %+v", err)
		}
		s.EventDetail = &impl
	}

	return nil
}
