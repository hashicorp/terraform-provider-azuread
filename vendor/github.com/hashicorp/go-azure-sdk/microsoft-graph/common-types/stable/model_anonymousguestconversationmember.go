package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConversationMember = AnonymousGuestConversationMember{}

type AnonymousGuestConversationMember struct {
	// Unique ID that represents the user. Note: This ID can change if the user leaves and rejoins the meeting, or joins
	// from a different device.
	AnonymousGuestId nullable.Type[string] `json:"anonymousGuestId,omitempty"`

	// Fields inherited from ConversationMember

	// The display name of the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The roles for that user. This property contains more qualifiers only when relevant - for example, if the member has
	// owner privileges, the roles property contains owner as one of the values. Similarly, if the member is an in-tenant
	// guest, the roles property contains guest as one of the values. A basic member shouldn't have any values specified in
	// the roles property. An Out-of-tenant external member is assigned the owner role.
	Roles *[]string `json:"roles,omitempty"`

	// The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is
	// settable only for members of a chat.
	VisibleHistoryStartDateTime nullable.Type[string] `json:"visibleHistoryStartDateTime,omitempty"`

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

func (s AnonymousGuestConversationMember) ConversationMember() BaseConversationMemberImpl {
	return BaseConversationMemberImpl{
		DisplayName:                 s.DisplayName,
		Roles:                       s.Roles,
		VisibleHistoryStartDateTime: s.VisibleHistoryStartDateTime,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s AnonymousGuestConversationMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AnonymousGuestConversationMember{}

func (s AnonymousGuestConversationMember) MarshalJSON() ([]byte, error) {
	type wrapper AnonymousGuestConversationMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AnonymousGuestConversationMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AnonymousGuestConversationMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.anonymousGuestConversationMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AnonymousGuestConversationMember: %+v", err)
	}

	return encoded, nil
}
