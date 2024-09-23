package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeSearchEngineBase interface {
	EdgeSearchEngineBase() BaseEdgeSearchEngineBaseImpl
}

var _ EdgeSearchEngineBase = BaseEdgeSearchEngineBaseImpl{}

type BaseEdgeSearchEngineBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEdgeSearchEngineBaseImpl) EdgeSearchEngineBase() BaseEdgeSearchEngineBaseImpl {
	return s
}

var _ EdgeSearchEngineBase = RawEdgeSearchEngineBaseImpl{}

// RawEdgeSearchEngineBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEdgeSearchEngineBaseImpl struct {
	edgeSearchEngineBase BaseEdgeSearchEngineBaseImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawEdgeSearchEngineBaseImpl) EdgeSearchEngineBase() BaseEdgeSearchEngineBaseImpl {
	return s.edgeSearchEngineBase
}

func UnmarshalEdgeSearchEngineBaseImplementation(input []byte) (EdgeSearchEngineBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeSearchEngineBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeSearchEngine") {
		var out EdgeSearchEngine
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeSearchEngine: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeSearchEngineCustom") {
		var out EdgeSearchEngineCustom
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeSearchEngineCustom: %+v", err)
		}
		return out, nil
	}

	var parent BaseEdgeSearchEngineBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEdgeSearchEngineBaseImpl: %+v", err)
	}

	return RawEdgeSearchEngineBaseImpl{
		edgeSearchEngineBase: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
