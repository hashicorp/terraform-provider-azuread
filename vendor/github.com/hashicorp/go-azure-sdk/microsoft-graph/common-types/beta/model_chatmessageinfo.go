package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ChatMessageInfo{}

type ChatMessageInfo struct {
	// Body of the chatMessage. This will still contain markers for @mentions and attachments even though the object doesn't
	// return @mentions and attachments.
	Body *ItemBody `json:"body,omitempty"`

	// Date time object representing the time at which message was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example,
	// members were added, and so on. For event messages, the messageType property is set to systemEventMessage.
	EventDetail *EventMessageDetail `json:"eventDetail,omitempty"`

	// Information about the sender of the message.
	From *ChatMessageFromIdentitySet `json:"from,omitempty"`

	// If set to true, the original message has been deleted.
	IsDeleted nullable.Type[bool] `json:"isDeleted,omitempty"`

	MessageType *ChatMessageType `json:"messageType,omitempty"`

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

func (s ChatMessageInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChatMessageInfo{}

func (s ChatMessageInfo) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessageInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessageInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessageInfo: %+v", err)
	}

	delete(decoded, "eventDetail")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMessageInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessageInfo: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChatMessageInfo{}

func (s *ChatMessageInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Body            *ItemBody                   `json:"body,omitempty"`
		CreatedDateTime nullable.Type[string]       `json:"createdDateTime,omitempty"`
		From            *ChatMessageFromIdentitySet `json:"from,omitempty"`
		IsDeleted       nullable.Type[bool]         `json:"isDeleted,omitempty"`
		MessageType     *ChatMessageType            `json:"messageType,omitempty"`
		Id              *string                     `json:"id,omitempty"`
		ODataId         *string                     `json:"@odata.id,omitempty"`
		ODataType       *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Body = decoded.Body
	s.CreatedDateTime = decoded.CreatedDateTime
	s.From = decoded.From
	s.IsDeleted = decoded.IsDeleted
	s.MessageType = decoded.MessageType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChatMessageInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["eventDetail"]; ok {
		impl, err := UnmarshalEventMessageDetailImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EventDetail' for 'ChatMessageInfo': %+v", err)
		}
		s.EventDetail = &impl
	}

	return nil
}
