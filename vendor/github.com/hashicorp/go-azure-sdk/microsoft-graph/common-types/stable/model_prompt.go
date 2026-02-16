package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Prompt interface {
	Prompt() BasePromptImpl
}

var _ Prompt = BasePromptImpl{}

type BasePromptImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePromptImpl) Prompt() BasePromptImpl {
	return s
}

var _ Prompt = RawPromptImpl{}

// RawPromptImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPromptImpl struct {
	prompt BasePromptImpl
	Type   string
	Values map[string]interface{}
}

func (s RawPromptImpl) Prompt() BasePromptImpl {
	return s.prompt
}

func UnmarshalPromptImplementation(input []byte) (Prompt, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Prompt into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.mediaPrompt") {
		var out MediaPrompt
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MediaPrompt: %+v", err)
		}
		return out, nil
	}

	var parent BasePromptImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePromptImpl: %+v", err)
	}

	return RawPromptImpl{
		prompt: parent,
		Type:   value,
		Values: temp,
	}, nil

}
