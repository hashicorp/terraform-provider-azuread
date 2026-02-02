package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsUserAgent = CallRecordsServiceUserAgent{}

type CallRecordsServiceUserAgent struct {
	Role *CallRecordsServiceRole `json:"role,omitempty"`

	// Fields inherited from CallRecordsUserAgent

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

func (s CallRecordsServiceUserAgent) CallRecordsUserAgent() BaseCallRecordsUserAgentImpl {
	return BaseCallRecordsUserAgentImpl{
		ApplicationVersion: s.ApplicationVersion,
		HeaderValue:        s.HeaderValue,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsServiceUserAgent{}

func (s CallRecordsServiceUserAgent) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsServiceUserAgent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsServiceUserAgent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsServiceUserAgent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.serviceUserAgent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsServiceUserAgent: %+v", err)
	}

	return encoded, nil
}
