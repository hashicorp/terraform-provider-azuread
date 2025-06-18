package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AiInteractionAttachment{}

type AiInteractionAttachment struct {
	// The identifier for the attachment. This identifier is only unique within the message scope.
	AttachmentId nullable.Type[string] `json:"attachmentId,omitempty"`

	// The content of the attachment.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The type of the content. For example, reference, file, and image/imageType.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// The URL of the content.
	ContentUrl nullable.Type[string] `json:"contentUrl,omitempty"`

	// The name of the attachment.
	Name nullable.Type[string] `json:"name,omitempty"`

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

func (s AiInteractionAttachment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AiInteractionAttachment{}

func (s AiInteractionAttachment) MarshalJSON() ([]byte, error) {
	type wrapper AiInteractionAttachment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AiInteractionAttachment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AiInteractionAttachment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aiInteractionAttachment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AiInteractionAttachment: %+v", err)
	}

	return encoded, nil
}
