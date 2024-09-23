package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Chat{}

type Chat struct {
	ChatType *ChatType `json:"chatType,omitempty"`

	// The user or application that created the chat. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date and time at which the chat was created. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// A collection of all the apps in the chat. Nullable.
	InstalledApps *[]TeamsAppInstallation `json:"installedApps,omitempty"`

	// Indicates whether the chat is hidden for all its members. Read-only.
	IsHiddenForAllMembers nullable.Type[bool] `json:"isHiddenForAllMembers,omitempty"`

	// Preview of the last message sent in the chat. Null if no messages are sent in the chat. Currently, only the list
	// chats operation supports this property.
	LastMessagePreview *ChatMessageInfo `json:"lastMessagePreview,omitempty"`

	// Date and time at which the chat was renamed or list of members were last changed. Read-only.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// A collection of all the members in the chat. Nullable.
	Members *[]ConversationMember `json:"members,omitempty"`

	// A collection of all the messages in the chat. Nullable.
	Messages *[]ChatMessage `json:"messages,omitempty"`

	// Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is
	// empty. Read-only.
	OnlineMeetingInfo *TeamworkOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty"`

	// A collection of all the Teams async operations that ran or are running on the chat. Nullable.
	Operations *[]TeamsAsyncOperation `json:"operations,omitempty"`

	// A collection of permissions granted to apps for the chat.
	PermissionGrants *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`

	// A collection of all the pinned messages in the chat. Nullable.
	PinnedMessages *[]PinnedChatMessageInfo `json:"pinnedMessages,omitempty"`

	// A collection of all the tabs in the chat. Nullable.
	Tabs *[]TeamsTab `json:"tabs,omitempty"`

	// The identifier of the tenant in which the chat was created. Read-only.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// (Optional) Subject or topic for the chat. Only available for group chats.
	Topic nullable.Type[string] `json:"topic,omitempty"`

	// Represents caller-specific information about the chat, such as last message read date and time. This property is
	// populated only when the request is made in a delegated context.
	Viewpoint *ChatViewpoint `json:"viewpoint,omitempty"`

	// The URL for the chat in Microsoft Teams. The URL should be treated as an opaque blob, and not parsed. Read-only.
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

func (s Chat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Chat{}

func (s Chat) MarshalJSON() ([]byte, error) {
	type wrapper Chat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Chat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Chat: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "isHiddenForAllMembers")
	delete(decoded, "lastUpdatedDateTime")
	delete(decoded, "onlineMeetingInfo")
	delete(decoded, "tenantId")
	delete(decoded, "webUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Chat: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Chat{}

func (s *Chat) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChatType              *ChatType                          `json:"chatType,omitempty"`
		CreatedDateTime       nullable.Type[string]              `json:"createdDateTime,omitempty"`
		IsHiddenForAllMembers nullable.Type[bool]                `json:"isHiddenForAllMembers,omitempty"`
		LastMessagePreview    *ChatMessageInfo                   `json:"lastMessagePreview,omitempty"`
		LastUpdatedDateTime   nullable.Type[string]              `json:"lastUpdatedDateTime,omitempty"`
		Messages              *[]ChatMessage                     `json:"messages,omitempty"`
		OnlineMeetingInfo     *TeamworkOnlineMeetingInfo         `json:"onlineMeetingInfo,omitempty"`
		Operations            *[]TeamsAsyncOperation             `json:"operations,omitempty"`
		PermissionGrants      *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
		PinnedMessages        *[]PinnedChatMessageInfo           `json:"pinnedMessages,omitempty"`
		Tabs                  *[]TeamsTab                        `json:"tabs,omitempty"`
		TenantId              nullable.Type[string]              `json:"tenantId,omitempty"`
		Topic                 nullable.Type[string]              `json:"topic,omitempty"`
		Viewpoint             *ChatViewpoint                     `json:"viewpoint,omitempty"`
		WebUrl                nullable.Type[string]              `json:"webUrl,omitempty"`
		Id                    *string                            `json:"id,omitempty"`
		ODataId               *string                            `json:"@odata.id,omitempty"`
		ODataType             *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChatType = decoded.ChatType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.IsHiddenForAllMembers = decoded.IsHiddenForAllMembers
	s.LastMessagePreview = decoded.LastMessagePreview
	s.LastUpdatedDateTime = decoded.LastUpdatedDateTime
	s.Messages = decoded.Messages
	s.OnlineMeetingInfo = decoded.OnlineMeetingInfo
	s.Operations = decoded.Operations
	s.PermissionGrants = decoded.PermissionGrants
	s.PinnedMessages = decoded.PinnedMessages
	s.Tabs = decoded.Tabs
	s.TenantId = decoded.TenantId
	s.Topic = decoded.Topic
	s.Viewpoint = decoded.Viewpoint
	s.WebUrl = decoded.WebUrl
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Chat into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'Chat': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["installedApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InstalledApps into list []json.RawMessage: %+v", err)
		}

		output := make([]TeamsAppInstallation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalTeamsAppInstallationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InstalledApps' for 'Chat': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InstalledApps = &output
	}

	if v, ok := temp["members"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Members into list []json.RawMessage: %+v", err)
		}

		output := make([]ConversationMember, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConversationMemberImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Members' for 'Chat': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Members = &output
	}

	return nil
}
