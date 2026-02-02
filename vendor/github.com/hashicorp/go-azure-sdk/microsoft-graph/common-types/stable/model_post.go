package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OutlookItem = Post{}

type Post struct {
	// Read-only. Nullable. Supports $expand.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// The contents of the post. This is a default property. This property can be null.
	Body *ItemBody `json:"body,omitempty"`

	// Unique ID of the conversation. Read-only.
	ConversationId nullable.Type[string] `json:"conversationId,omitempty"`

	// Unique ID of the conversation thread. Read-only.
	ConversationThreadId nullable.Type[string] `json:"conversationThreadId,omitempty"`

	// The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand.
	Extensions *[]Extension `json:"extensions,omitempty"`

	From Recipient `json:"from"`

	// Indicates whether the post has at least one attachment. This is a default property.
	HasAttachments *bool `json:"hasAttachments,omitempty"`

	// Read-only. Supports $expand.
	InReplyTo *Post `json:"inReplyTo,omitempty"`

	// The collection of multi-value extended properties defined for the post. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// Conversation participants that were added to the thread as part of this post.
	NewParticipants *[]Recipient `json:"newParticipants,omitempty"`

	// Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ReceivedDateTime *string `json:"receivedDateTime,omitempty"`

	// Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the
	// case when Sender is not specified. This is a default property.
	Sender Recipient `json:"sender"`

	// The collection of single-value extended properties defined for the post. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// Fields inherited from OutlookItem

	// The categories associated with the item
	Categories *[]string `json:"categories,omitempty"`

	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange
	// to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s Post) OutlookItem() BaseOutlookItemImpl {
	return BaseOutlookItemImpl{
		Categories:           s.Categories,
		ChangeKey:            s.ChangeKey,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Post) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Post{}

func (s Post) MarshalJSON() ([]byte, error) {
	type wrapper Post
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Post: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Post: %+v", err)
	}

	delete(decoded, "attachments")
	delete(decoded, "conversationId")
	delete(decoded, "conversationThreadId")
	delete(decoded, "extensions")
	delete(decoded, "inReplyTo")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.post"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Post: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Post{}

func (s *Post) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Body                          *ItemBody                            `json:"body,omitempty"`
		ConversationId                nullable.Type[string]                `json:"conversationId,omitempty"`
		ConversationThreadId          nullable.Type[string]                `json:"conversationThreadId,omitempty"`
		HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
		InReplyTo                     *Post                                `json:"inReplyTo,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		ReceivedDateTime              *string                              `json:"receivedDateTime,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		Categories                    *[]string                            `json:"categories,omitempty"`
		ChangeKey                     nullable.Type[string]                `json:"changeKey,omitempty"`
		CreatedDateTime               nullable.Type[string]                `json:"createdDateTime,omitempty"`
		LastModifiedDateTime          nullable.Type[string]                `json:"lastModifiedDateTime,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Body = decoded.Body
	s.ConversationId = decoded.ConversationId
	s.ConversationThreadId = decoded.ConversationThreadId
	s.HasAttachments = decoded.HasAttachments
	s.InReplyTo = decoded.InReplyTo
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.ReceivedDateTime = decoded.ReceivedDateTime
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.Categories = decoded.Categories
	s.ChangeKey = decoded.ChangeKey
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Post into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attachments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attachments into list []json.RawMessage: %+v", err)
		}

		output := make([]Attachment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAttachmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attachments' for 'Post': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attachments = &output
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'Post': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	if v, ok := temp["from"]; ok {
		impl, err := UnmarshalRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'From' for 'Post': %+v", err)
		}
		s.From = impl
	}

	if v, ok := temp["newParticipants"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling NewParticipants into list []json.RawMessage: %+v", err)
		}

		output := make([]Recipient, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRecipientImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'NewParticipants' for 'Post': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.NewParticipants = &output
	}

	if v, ok := temp["sender"]; ok {
		impl, err := UnmarshalRecipientImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Sender' for 'Post': %+v", err)
		}
		s.Sender = impl
	}

	return nil
}
