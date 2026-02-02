package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRoute struct {
	Final IdentitySet `json:"final"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Original    IdentitySet  `json:"original"`
	RoutingType *RoutingType `json:"routingType,omitempty"`
}

var _ json.Unmarshaler = &CallRoute{}

func (s *CallRoute) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId     *string      `json:"@odata.id,omitempty"`
		ODataType   *string      `json:"@odata.type,omitempty"`
		RoutingType *RoutingType `json:"routingType,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RoutingType = decoded.RoutingType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRoute into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["final"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Final' for 'CallRoute': %+v", err)
		}
		s.Final = impl
	}

	if v, ok := temp["original"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Original' for 'CallRoute': %+v", err)
		}
		s.Original = impl
	}

	return nil
}
