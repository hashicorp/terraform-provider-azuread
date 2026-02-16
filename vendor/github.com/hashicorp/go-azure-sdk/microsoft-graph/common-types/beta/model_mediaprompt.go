package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Prompt = MediaPrompt{}

type MediaPrompt struct {
	Loop      nullable.Type[int64] `json:"loop,omitempty"`
	MediaInfo *MediaInfo           `json:"mediaInfo,omitempty"`

	// Fields inherited from Prompt

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MediaPrompt) Prompt() BasePromptImpl {
	return BasePromptImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MediaPrompt{}

func (s MediaPrompt) MarshalJSON() ([]byte, error) {
	type wrapper MediaPrompt
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MediaPrompt: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MediaPrompt: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mediaPrompt"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MediaPrompt: %+v", err)
	}

	return encoded, nil
}
