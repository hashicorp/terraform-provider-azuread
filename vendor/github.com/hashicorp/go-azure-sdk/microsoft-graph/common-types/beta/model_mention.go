package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Mention{}

type Mention struct {
	// The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
	Application nullable.Type[string] `json:"application,omitempty"`

	// A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for
	// message.
	ClientReference nullable.Type[string] `json:"clientReference,omitempty"`

	// The email information of the user who made the mention.
	CreatedBy EmailAddress `json:"createdBy"`

	// The date and time that the mention is created on the client.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for
	// message.
	DeepLink nullable.Type[string] `json:"deepLink,omitempty"`

	// Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property
	// of the message instead.
	MentionText nullable.Type[string] `json:"mentionText,omitempty"`

	Mentioned EmailAddress `json:"mentioned"`

	// The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
	ServerCreatedDateTime nullable.Type[string] `json:"serverCreatedDateTime,omitempty"`

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

func (s Mention) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Mention{}

func (s Mention) MarshalJSON() ([]byte, error) {
	type wrapper Mention
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Mention: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Mention: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mention"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Mention: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Mention{}

func (s *Mention) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Application           nullable.Type[string] `json:"application,omitempty"`
		ClientReference       nullable.Type[string] `json:"clientReference,omitempty"`
		CreatedDateTime       nullable.Type[string] `json:"createdDateTime,omitempty"`
		DeepLink              nullable.Type[string] `json:"deepLink,omitempty"`
		MentionText           nullable.Type[string] `json:"mentionText,omitempty"`
		ServerCreatedDateTime nullable.Type[string] `json:"serverCreatedDateTime,omitempty"`
		Id                    *string               `json:"id,omitempty"`
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Application = decoded.Application
	s.ClientReference = decoded.ClientReference
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DeepLink = decoded.DeepLink
	s.MentionText = decoded.MentionText
	s.ServerCreatedDateTime = decoded.ServerCreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Mention into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'Mention': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["mentioned"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Mentioned' for 'Mention': %+v", err)
		}
		s.Mentioned = impl
	}

	return nil
}
