package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationMember interface {
	Entity
	ConversationMember() BaseConversationMemberImpl
}

var _ ConversationMember = BaseConversationMemberImpl{}

type BaseConversationMemberImpl struct {
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

func (s BaseConversationMemberImpl) ConversationMember() BaseConversationMemberImpl {
	return s
}

func (s BaseConversationMemberImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ConversationMember = RawConversationMemberImpl{}

// RawConversationMemberImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConversationMemberImpl struct {
	conversationMember BaseConversationMemberImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawConversationMemberImpl) ConversationMember() BaseConversationMemberImpl {
	return s.conversationMember
}

func (s RawConversationMemberImpl) Entity() BaseEntityImpl {
	return s.conversationMember.Entity()
}

var _ json.Marshaler = BaseConversationMemberImpl{}

func (s BaseConversationMemberImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseConversationMemberImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseConversationMemberImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseConversationMemberImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conversationMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseConversationMemberImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalConversationMemberImplementation(input []byte) (ConversationMember, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConversationMember into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aadUserConversationMember") {
		var out AadUserConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadUserConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.anonymousGuestConversationMember") {
		var out AnonymousGuestConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AnonymousGuestConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureCommunicationServicesUserConversationMember") {
		var out AzureCommunicationServicesUserConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureCommunicationServicesUserConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAccountUserConversationMember") {
		var out MicrosoftAccountUserConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAccountUserConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.skypeForBusinessUserConversationMember") {
		var out SkypeForBusinessUserConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkypeForBusinessUserConversationMember: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.skypeUserConversationMember") {
		var out SkypeUserConversationMember
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SkypeUserConversationMember: %+v", err)
		}
		return out, nil
	}

	var parent BaseConversationMemberImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConversationMemberImpl: %+v", err)
	}

	return RawConversationMemberImpl{
		conversationMember: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
