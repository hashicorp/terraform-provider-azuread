package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConversationMember = AadUserConversationMember{}

type AadUserConversationMember struct {
	// The email address of the user.
	Email nullable.Type[string] `json:"email,omitempty"`

	// TenantId which the Microsoft Entra user belongs to.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	User *User `json:"user,omitempty"`

	// The GUID of the user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from ConversationMember

	// The display name of the user.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The roles for that user. This property contains additional qualifiers only when relevant - for example, if the member
	// has owner privileges, the roles property contains owner as one of the values. Similarly, if the member is an
	// in-tenant guest, the roles property contains guest as one of the values. A basic member should not have any values
	// specified in the roles property. An Out-of-tenant external member is assigned the owner role.
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

func (s AadUserConversationMember) ConversationMember() BaseConversationMemberImpl {
	return BaseConversationMemberImpl{
		DisplayName:                 s.DisplayName,
		Roles:                       s.Roles,
		VisibleHistoryStartDateTime: s.VisibleHistoryStartDateTime,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s AadUserConversationMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AadUserConversationMember{}

func (s AadUserConversationMember) MarshalJSON() ([]byte, error) {
	type wrapper AadUserConversationMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AadUserConversationMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AadUserConversationMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aadUserConversationMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AadUserConversationMember: %+v", err)
	}

	return encoded, nil
}
