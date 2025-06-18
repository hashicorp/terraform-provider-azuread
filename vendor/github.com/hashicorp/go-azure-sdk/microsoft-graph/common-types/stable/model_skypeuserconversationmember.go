package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConversationMember = SkypeUserConversationMember{}

type SkypeUserConversationMember struct {
	// Skype ID of the user.
	SkypeId nullable.Type[string] `json:"skypeId,omitempty"`

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

func (s SkypeUserConversationMember) ConversationMember() BaseConversationMemberImpl {
	return BaseConversationMemberImpl{
		DisplayName:                 s.DisplayName,
		Roles:                       s.Roles,
		VisibleHistoryStartDateTime: s.VisibleHistoryStartDateTime,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s SkypeUserConversationMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SkypeUserConversationMember{}

func (s SkypeUserConversationMember) MarshalJSON() ([]byte, error) {
	type wrapper SkypeUserConversationMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SkypeUserConversationMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SkypeUserConversationMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.skypeUserConversationMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SkypeUserConversationMember: %+v", err)
	}

	return encoded, nil
}
