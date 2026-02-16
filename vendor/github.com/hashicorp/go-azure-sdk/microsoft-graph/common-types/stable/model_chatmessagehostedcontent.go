package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamworkHostedContent = ChatMessageHostedContent{}

type ChatMessageHostedContent struct {

	// Fields inherited from TeamworkHostedContent

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

func (s ChatMessageHostedContent) TeamworkHostedContent() BaseTeamworkHostedContentImpl {
	return BaseTeamworkHostedContentImpl{
		ContentBytes: s.ContentBytes,
		ContentType:  s.ContentType,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s ChatMessageHostedContent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChatMessageHostedContent{}

func (s ChatMessageHostedContent) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessageHostedContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessageHostedContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessageHostedContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMessageHostedContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessageHostedContent: %+v", err)
	}

	return encoded, nil
}
