package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeHomeButtonConfiguration interface {
	EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl
}

var _ EdgeHomeButtonConfiguration = BaseEdgeHomeButtonConfigurationImpl{}

type BaseEdgeHomeButtonConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEdgeHomeButtonConfigurationImpl) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return s
}

var _ EdgeHomeButtonConfiguration = RawEdgeHomeButtonConfigurationImpl{}

// RawEdgeHomeButtonConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEdgeHomeButtonConfigurationImpl struct {
	edgeHomeButtonConfiguration BaseEdgeHomeButtonConfigurationImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawEdgeHomeButtonConfigurationImpl) EdgeHomeButtonConfiguration() BaseEdgeHomeButtonConfigurationImpl {
	return s.edgeHomeButtonConfiguration
}

func UnmarshalEdgeHomeButtonConfigurationImplementation(input []byte) (EdgeHomeButtonConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EdgeHomeButtonConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeHomeButtonHidden") {
		var out EdgeHomeButtonHidden
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeHomeButtonHidden: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeHomeButtonLoadsStartPage") {
		var out EdgeHomeButtonLoadsStartPage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeHomeButtonLoadsStartPage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeHomeButtonOpensCustomURL") {
		var out EdgeHomeButtonOpensCustomURL
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeHomeButtonOpensCustomURL: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.edgeHomeButtonOpensNewTab") {
		var out EdgeHomeButtonOpensNewTab
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdgeHomeButtonOpensNewTab: %+v", err)
		}
		return out, nil
	}

	var parent BaseEdgeHomeButtonConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEdgeHomeButtonConfigurationImpl: %+v", err)
	}

	return RawEdgeHomeButtonConfigurationImpl{
		edgeHomeButtonConfiguration: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
