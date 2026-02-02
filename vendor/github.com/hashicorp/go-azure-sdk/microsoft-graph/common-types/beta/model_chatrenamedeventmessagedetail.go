package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ChatRenamedEventMessageDetail{}

type ChatRenamedEventMessageDetail struct {
	// The updated name of the chat.
	ChatDisplayName nullable.Type[string] `json:"chatDisplayName,omitempty"`

	// Unique identifier of the chat.
	ChatId nullable.Type[string] `json:"chatId,omitempty"`

	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChatRenamedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChatRenamedEventMessageDetail{}

func (s ChatRenamedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ChatRenamedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatRenamedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatRenamedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatRenamedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatRenamedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChatRenamedEventMessageDetail{}

func (s *ChatRenamedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChatDisplayName nullable.Type[string] `json:"chatDisplayName,omitempty"`
		ChatId          nullable.Type[string] `json:"chatId,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChatDisplayName = decoded.ChatDisplayName
	s.ChatId = decoded.ChatId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChatRenamedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ChatRenamedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
