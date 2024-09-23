package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserAgent interface {
	CallRecordsUserAgent() BaseCallRecordsUserAgentImpl
}

var _ CallRecordsUserAgent = BaseCallRecordsUserAgentImpl{}

type BaseCallRecordsUserAgentImpl struct {
	// Identifies the version of application software used by this endpoint.
	ApplicationVersion nullable.Type[string] `json:"applicationVersion,omitempty"`

	// User-agent header value reported by this endpoint.
	HeaderValue nullable.Type[string] `json:"headerValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCallRecordsUserAgentImpl) CallRecordsUserAgent() BaseCallRecordsUserAgentImpl {
	return s
}

var _ CallRecordsUserAgent = RawCallRecordsUserAgentImpl{}

// RawCallRecordsUserAgentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCallRecordsUserAgentImpl struct {
	callRecordsUserAgent BaseCallRecordsUserAgentImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawCallRecordsUserAgentImpl) CallRecordsUserAgent() BaseCallRecordsUserAgentImpl {
	return s.callRecordsUserAgent
}

func UnmarshalCallRecordsUserAgentImplementation(input []byte) (CallRecordsUserAgent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsUserAgent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.clientUserAgent") {
		var out CallRecordsClientUserAgent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsClientUserAgent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecords.serviceUserAgent") {
		var out CallRecordsServiceUserAgent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordsServiceUserAgent: %+v", err)
		}
		return out, nil
	}

	var parent BaseCallRecordsUserAgentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCallRecordsUserAgentImpl: %+v", err)
	}

	return RawCallRecordsUserAgentImpl{
		callRecordsUserAgent: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
