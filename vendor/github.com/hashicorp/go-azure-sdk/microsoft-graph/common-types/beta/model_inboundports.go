package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundPorts interface {
	InboundPorts() BaseInboundPortsImpl
}

var _ InboundPorts = BaseInboundPortsImpl{}

type BaseInboundPortsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseInboundPortsImpl) InboundPorts() BaseInboundPortsImpl {
	return s
}

var _ InboundPorts = RawInboundPortsImpl{}

// RawInboundPortsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInboundPortsImpl struct {
	inboundPorts BaseInboundPortsImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawInboundPortsImpl) InboundPorts() BaseInboundPortsImpl {
	return s.inboundPorts
}

func UnmarshalInboundPortsImplementation(input []byte) (InboundPorts, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InboundPorts into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allInboundPorts") {
		var out AllInboundPorts
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllInboundPorts: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedInboundPorts") {
		var out EnumeratedInboundPorts
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedInboundPorts: %+v", err)
		}
		return out, nil
	}

	var parent BaseInboundPortsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInboundPortsImpl: %+v", err)
	}

	return RawInboundPortsImpl{
		inboundPorts: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
