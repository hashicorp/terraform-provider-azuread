package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ConversationThread{}

type ConversationThread struct {
	// The Cc: recipients for the thread. Returned only on $select.
	CcRecipients *[]Recipient `json:"ccRecipients,omitempty"`

	// Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// Indicates if the thread is locked. Returned by default.
	IsLocked *bool `json:"isLocked,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
	LastDeliveredDateTime *string `json:"lastDeliveredDateTime,omitempty"`

	Posts *[]Post `json:"posts,omitempty"`

	// A short summary from the body of the latest post in this conversation. Returned by default.
	Preview *string `json:"preview,omitempty"`

	// The To: recipients for the thread. Returned only on $select.
	ToRecipients *[]Recipient `json:"toRecipients,omitempty"`

	// The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
	// Returned by default.
	Topic *string `json:"topic,omitempty"`

	// All the users that sent a message to this thread. Returned by default.
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

func (s ConversationThread) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConversationThread{}

func (s ConversationThread) MarshalJSON() ([]byte, error) {
	type wrapper ConversationThread
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConversationThread: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConversationThread: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conversationThread"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConversationThread: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ConversationThread{}

func (s *ConversationThread) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		HasAttachments        *bool     `json:"hasAttachments,omitempty"`
		IsLocked              *bool     `json:"isLocked,omitempty"`
		LastDeliveredDateTime *string   `json:"lastDeliveredDateTime,omitempty"`
		Posts                 *[]Post   `json:"posts,omitempty"`
		Preview               *string   `json:"preview,omitempty"`
		Topic                 *string   `json:"topic,omitempty"`
		UniqueSenders         *[]string `json:"uniqueSenders,omitempty"`
		Id                    *string   `json:"id,omitempty"`
		ODataId               *string   `json:"@odata.id,omitempty"`
		ODataType             *string   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.HasAttachments = decoded.HasAttachments
	s.IsLocked = decoded.IsLocked
	s.LastDeliveredDateTime = decoded.LastDeliveredDateTime
	s.Posts = decoded.Posts
	s.Preview = decoded.Preview
	s.Topic = decoded.Topic
	s.UniqueSenders = decoded.UniqueSenders
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConversationThread into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["ccRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CcRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CcRecipients' for 'ConversationThread': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CcRecipients = &output
	}

	if v, ok := temp["toRecipients"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ToRecipients into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ToRecipients' for 'ConversationThread': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ToRecipients = &output
	}

	return nil
}
