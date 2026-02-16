package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MediaConfig = ServiceHostedMediaConfig{}

type ServiceHostedMediaConfig struct {
	// The list of media to pre-fetch.
	PreFetchMedia *[]MediaInfo `json:"preFetchMedia,omitempty"`

	// Fields inherited from MediaConfig

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServiceHostedMediaConfig) MediaConfig() BaseMediaConfigImpl {
	return BaseMediaConfigImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceHostedMediaConfig{}

func (s ServiceHostedMediaConfig) MarshalJSON() ([]byte, error) {
	type wrapper ServiceHostedMediaConfig
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceHostedMediaConfig: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceHostedMediaConfig: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceHostedMediaConfig"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceHostedMediaConfig: %+v", err)
	}

	return encoded, nil
}
