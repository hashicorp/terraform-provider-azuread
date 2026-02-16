package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsEndpoint = CallRecordsServiceEndpoint{}

type CallRecordsServiceEndpoint struct {

	// Fields inherited from CallRecordsEndpoint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// User-agent reported by this endpoint.
	UserAgent CallRecordsUserAgent `json:"userAgent"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallRecordsServiceEndpoint) CallRecordsEndpoint() BaseCallRecordsEndpointImpl {
	return BaseCallRecordsEndpointImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		UserAgent: s.UserAgent,
	}
}

var _ json.Marshaler = CallRecordsServiceEndpoint{}

func (s CallRecordsServiceEndpoint) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsServiceEndpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsServiceEndpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsServiceEndpoint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.serviceEndpoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsServiceEndpoint: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordsServiceEndpoint{}

func (s *CallRecordsServiceEndpoint) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling CallRecordsServiceEndpoint into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["userAgent"]; ok {
		impl, err := UnmarshalCallRecordsUserAgentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'UserAgent' for 'CallRecordsServiceEndpoint': %+v", err)
		}
		s.UserAgent = impl
	}

	return nil
}
