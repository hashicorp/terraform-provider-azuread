package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionEndpointConfiguration = HttpRequestEndpoint{}

type HttpRequestEndpoint struct {
	// The HTTP endpoint that a custom extension calls.
	TargetUrl nullable.Type[string] `json:"targetUrl,omitempty"`

	// Fields inherited from CustomExtensionEndpointConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s HttpRequestEndpoint) CustomExtensionEndpointConfiguration() BaseCustomExtensionEndpointConfigurationImpl {
	return BaseCustomExtensionEndpointConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HttpRequestEndpoint{}

func (s HttpRequestEndpoint) MarshalJSON() ([]byte, error) {
	type wrapper HttpRequestEndpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HttpRequestEndpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HttpRequestEndpoint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.httpRequestEndpoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HttpRequestEndpoint: %+v", err)
	}

	return encoded, nil
}
