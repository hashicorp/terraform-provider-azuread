package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHostedContent interface {
	Entity
	TeamworkHostedContent() BaseTeamworkHostedContentImpl
}

var _ TeamworkHostedContent = BaseTeamworkHostedContentImpl{}

type BaseTeamworkHostedContentImpl struct {
	// Write only. Bytes for the hosted content (such as images).
	ContentBytes nullable.Type[string] `json:"contentBytes,omitempty"`

	// Write only. Content type. such as image/png, image/jpg.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

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

func (s BaseTeamworkHostedContentImpl) TeamworkHostedContent() BaseTeamworkHostedContentImpl {
	return s
}

func (s BaseTeamworkHostedContentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ TeamworkHostedContent = RawTeamworkHostedContentImpl{}

// RawTeamworkHostedContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTeamworkHostedContentImpl struct {
	teamworkHostedContent BaseTeamworkHostedContentImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawTeamworkHostedContentImpl) TeamworkHostedContent() BaseTeamworkHostedContentImpl {
	return s.teamworkHostedContent
}

func (s RawTeamworkHostedContentImpl) Entity() BaseEntityImpl {
	return s.teamworkHostedContent.Entity()
}

var _ json.Marshaler = BaseTeamworkHostedContentImpl{}

func (s BaseTeamworkHostedContentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseTeamworkHostedContentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseTeamworkHostedContentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseTeamworkHostedContentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamworkHostedContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseTeamworkHostedContentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalTeamworkHostedContentImplementation(input []byte) (TeamworkHostedContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkHostedContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageHostedContent") {
		var out ChatMessageHostedContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageHostedContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeamworkHostedContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeamworkHostedContentImpl: %+v", err)
	}

	return RawTeamworkHostedContentImpl{
		teamworkHostedContent: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
