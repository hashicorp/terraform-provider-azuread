package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsUserAgent = CallRecordsClientUserAgent{}

type CallRecordsClientUserAgent struct {
	// The unique identifier of the Microsoft Entra application used by this endpoint.
	AzureADAppId nullable.Type[string] `json:"azureADAppId,omitempty"`

	// Immutable resource identifier of the Azure Communication Service associated with this endpoint based on Communication
	// Services APIs.
	CommunicationServiceId nullable.Type[string] `json:"communicationServiceId,omitempty"`

	Platform      *CallRecordsClientPlatform `json:"platform,omitempty"`
	ProductFamily *CallRecordsProductFamily  `json:"productFamily,omitempty"`

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

func (s CallRecordsClientUserAgent) CallRecordsUserAgent() BaseCallRecordsUserAgentImpl {
	return BaseCallRecordsUserAgentImpl{
		ApplicationVersion: s.ApplicationVersion,
		HeaderValue:        s.HeaderValue,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsClientUserAgent{}

func (s CallRecordsClientUserAgent) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsClientUserAgent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsClientUserAgent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsClientUserAgent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.clientUserAgent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsClientUserAgent: %+v", err)
	}

	return encoded, nil
}
