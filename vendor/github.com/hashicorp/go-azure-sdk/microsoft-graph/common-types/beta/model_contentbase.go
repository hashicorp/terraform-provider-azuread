package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentBase interface {
	ContentBase() BaseContentBaseImpl
}

var _ ContentBase = BaseContentBaseImpl{}

type BaseContentBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseContentBaseImpl) ContentBase() BaseContentBaseImpl {
	return s
}

var _ ContentBase = RawContentBaseImpl{}

// RawContentBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawContentBaseImpl struct {
	contentBase BaseContentBaseImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawContentBaseImpl) ContentBase() BaseContentBaseImpl {
	return s.contentBase
}

func (s RawContentBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalContentBaseImplementation(input []byte) (ContentBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.binaryContent") {
		var out BinaryContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BinaryContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.textContent") {
		var out TextContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseContentBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseContentBaseImpl: %+v", err)
	}

	return RawContentBaseImpl{
		contentBase: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
