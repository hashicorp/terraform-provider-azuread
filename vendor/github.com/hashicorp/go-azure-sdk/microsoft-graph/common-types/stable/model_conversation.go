package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Conversation{}

type Conversation struct {
	// Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne)
	// and $search.
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastDeliveredDateTime *string `json:"lastDeliveredDateTime,omitempty"`

	// A short summary from the body of the latest post in this conversation. Supports $filter (eq, ne, le, ge).
	Preview *string `json:"preview,omitempty"`

	// A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
	Threads *[]ConversationThread `json:"threads,omitempty"`

	// The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
	Topic *string `json:"topic,omitempty"`

	// All the users that sent a message to this Conversation.
	UniqueSenders *[]string `json:"uniqueSenders,omitempty"`

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

func (s Conversation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Conversation{}

func (s Conversation) MarshalJSON() ([]byte, error) {
	type wrapper Conversation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Conversation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Conversation: %+v", err)
	}

	delete(decoded, "threads")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conversation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Conversation: %+v", err)
	}

	return encoded, nil
}
