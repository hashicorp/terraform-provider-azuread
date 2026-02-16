package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InboundPorts = EnumeratedInboundPorts{}

type EnumeratedInboundPorts struct {
	// Collection of ports that allow inbound traffic.
	Ports *[]string `json:"ports,omitempty"`

	// Fields inherited from InboundPorts

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EnumeratedInboundPorts) InboundPorts() BaseInboundPortsImpl {
	return BaseInboundPortsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnumeratedInboundPorts{}

func (s EnumeratedInboundPorts) MarshalJSON() ([]byte, error) {
	type wrapper EnumeratedInboundPorts
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnumeratedInboundPorts: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnumeratedInboundPorts: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enumeratedInboundPorts"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnumeratedInboundPorts: %+v", err)
	}

	return encoded, nil
}
