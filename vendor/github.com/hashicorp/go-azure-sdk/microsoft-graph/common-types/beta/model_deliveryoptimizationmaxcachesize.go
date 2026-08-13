package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationMaxCacheSize interface {
	DeliveryOptimizationMaxCacheSize() BaseDeliveryOptimizationMaxCacheSizeImpl
}

var _ DeliveryOptimizationMaxCacheSize = BaseDeliveryOptimizationMaxCacheSizeImpl{}

type BaseDeliveryOptimizationMaxCacheSizeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeliveryOptimizationMaxCacheSizeImpl) DeliveryOptimizationMaxCacheSize() BaseDeliveryOptimizationMaxCacheSizeImpl {
	return s
}

var _ DeliveryOptimizationMaxCacheSize = RawDeliveryOptimizationMaxCacheSizeImpl{}

// RawDeliveryOptimizationMaxCacheSizeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDeliveryOptimizationMaxCacheSizeImpl struct {
	deliveryOptimizationMaxCacheSize BaseDeliveryOptimizationMaxCacheSizeImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawDeliveryOptimizationMaxCacheSizeImpl) DeliveryOptimizationMaxCacheSize() BaseDeliveryOptimizationMaxCacheSizeImpl {
	return s.deliveryOptimizationMaxCacheSize
}

func (s RawDeliveryOptimizationMaxCacheSizeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDeliveryOptimizationMaxCacheSizeImplementation(input []byte) (DeliveryOptimizationMaxCacheSize, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationMaxCacheSize into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationMaxCacheSizeAbsolute") {
		var out DeliveryOptimizationMaxCacheSizeAbsolute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationMaxCacheSizeAbsolute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationMaxCacheSizePercentage") {
		var out DeliveryOptimizationMaxCacheSizePercentage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationMaxCacheSizePercentage: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeliveryOptimizationMaxCacheSizeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeliveryOptimizationMaxCacheSizeImpl: %+v", err)
	}

	return RawDeliveryOptimizationMaxCacheSizeImpl{
		deliveryOptimizationMaxCacheSize: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
