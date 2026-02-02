package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsEndpoint interface {
	CallRecordsEndpoint() BaseCallRecordsEndpointImpl
}

var _ CallRecordsEndpoint = BaseCallRecordsEndpointImpl{}

type BaseCallRecordsEndpointImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// User-agent reported by this endpoint.
	UserAgent CallRecordsUserAgent `json:"userAgent"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCallRecordsEndpointImpl) CallRecordsEndpoint() BaseCallRecordsEndpointImpl {
	return s
}

var _ CallRecordsEndpoint = RawCallRecordsEndpointImpl{}

// RawCallRecordsEndpointImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCallRecordsEndpointImpl struct {
	callRecordsEndpoint BaseCallRecordsEndpointImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawCallRecordsEndpointImpl) CallRecordsEndpoint() BaseCallRecordsEndpointImpl {
	return s.callRecordsEndpoint
}

var _ json.Unmarshaler = &BaseCallRecordsEndpointImpl{}

func (s *BaseCallRecordsEndpointImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseCallRecordsEndpointImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["userAgent"]; ok {
		impl, err := UnmarshalCallRecordsUserAgentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'UserAgent' for 'BaseCallRecordsEndpointImpl': %+v", err)
		}
		s.UserAgent = impl
	}

	return nil
}

func UnmarshalCallRecordsEndpointImplementation(input []byte) (CallRecordsEndpoint, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsEndpoint into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.participantEndpoint") {
		var out CallRecordsParticipantEndpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsParticipantEndpoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.serviceEndpoint") {
		var out CallRecordsServiceEndpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsServiceEndpoint: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallRecordsEndpointImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallRecordsEndpointImpl: %+v", err)
	}

	return RawCallRecordsEndpointImpl{
		callRecordsEndpoint: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
