package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaConfig interface {
	MediaConfig() BaseMediaConfigImpl
}

var _ MediaConfig = BaseMediaConfigImpl{}

type BaseMediaConfigImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseMediaConfigImpl) MediaConfig() BaseMediaConfigImpl {
	return s
}

var _ MediaConfig = RawMediaConfigImpl{}

// RawMediaConfigImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawMediaConfigImpl struct {
	mediaConfig BaseMediaConfigImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawMediaConfigImpl) MediaConfig() BaseMediaConfigImpl {
	return s.mediaConfig
}

func (s RawMediaConfigImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalMediaConfigImplementation(input []byte) (MediaConfig, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MediaConfig into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appHostedMediaConfig") {
		var out AppHostedMediaConfig
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppHostedMediaConfig: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceHostedMediaConfig") {
		var out ServiceHostedMediaConfig
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceHostedMediaConfig: %+v", err)
		}
		return out, nil
	}

	var parent BaseMediaConfigImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMediaConfigImpl: %+v", err)
	}

	return RawMediaConfigImpl{
		mediaConfig: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
