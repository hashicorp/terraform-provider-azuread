package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AiInteraction{}

type AiInteraction struct {
	// The data source for Copilot data. For example, IPM.SkypeTeams.Message.Copilot.Excel or
	// IPM.SkypeTeams.Message.Copilot.Loop.
	AppClass *string `json:"appClass,omitempty"`

	// The collection of documents attached to the interaction, such as cards and images.
	Attachments *[]AiInteractionAttachment `json:"attachments,omitempty"`

	// The body of the message, including the text of the body and its body type.
	Body *ItemBody `json:"body,omitempty"`

	// The identifer that maps to all contexts associated with an interaction.
	Contexts *[]AiInteractionContext `json:"contexts,omitempty"`

	// The type of the conversation. For example, appchat or bizchat.
	ConversationType nullable.Type[string] `json:"conversationType,omitempty"`

	// The time when the interaction was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The timestamp of when the interaction was last modified.
	Etag nullable.Type[string] `json:"etag,omitempty"`

	From            IdentitySet        `json:"from"`
	InteractionType *AiInteractionType `json:"interactionType,omitempty"`

	// The collection of links that appear in the interaction.
	Links *[]AiInteractionLink `json:"links,omitempty"`

	// The locale of the sender.
	Locale *string `json:"locale,omitempty"`

	// The collection of the entities that were mentioned in the interaction, including users, bots, and so on.
	Mentions *[]AiInteractionMention `json:"mentions,omitempty"`

	// The identifier that groups a user prompt with its Copilot response.
	RequestId *string `json:"requestId,omitempty"`

	// The thread ID or conversation identifier that maps to all Copilot sessions for the user.
	SessionId *string `json:"sessionId,omitempty"`

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

func (s AiInteraction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AiInteraction{}

func (s AiInteraction) MarshalJSON() ([]byte, error) {
	type wrapper AiInteraction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AiInteraction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AiInteraction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aiInteraction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AiInteraction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AiInteraction{}

func (s *AiInteraction) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppClass         *string                    `json:"appClass,omitempty"`
		Attachments      *[]AiInteractionAttachment `json:"attachments,omitempty"`
		Body             *ItemBody                  `json:"body,omitempty"`
		Contexts         *[]AiInteractionContext    `json:"contexts,omitempty"`
		ConversationType nullable.Type[string]      `json:"conversationType,omitempty"`
		CreatedDateTime  nullable.Type[string]      `json:"createdDateTime,omitempty"`
		Etag             nullable.Type[string]      `json:"etag,omitempty"`
		InteractionType  *AiInteractionType         `json:"interactionType,omitempty"`
		Links            *[]AiInteractionLink       `json:"links,omitempty"`
		Locale           *string                    `json:"locale,omitempty"`
		Mentions         *[]AiInteractionMention    `json:"mentions,omitempty"`
		RequestId        *string                    `json:"requestId,omitempty"`
		SessionId        *string                    `json:"sessionId,omitempty"`
		Id               *string                    `json:"id,omitempty"`
		ODataId          *string                    `json:"@odata.id,omitempty"`
		ODataType        *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppClass = decoded.AppClass
	s.Attachments = decoded.Attachments
	s.Body = decoded.Body
	s.Contexts = decoded.Contexts
	s.ConversationType = decoded.ConversationType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Etag = decoded.Etag
	s.InteractionType = decoded.InteractionType
	s.Links = decoded.Links
	s.Locale = decoded.Locale
	s.Mentions = decoded.Mentions
	s.RequestId = decoded.RequestId
	s.SessionId = decoded.SessionId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AiInteraction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["from"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'From' for 'AiInteraction': %+v", err)
		}
		s.From = impl
	}

	return nil
}
