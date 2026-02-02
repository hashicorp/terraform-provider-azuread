package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationMaxCacheSize = DeliveryOptimizationMaxCacheSizePercentage{}

type DeliveryOptimizationMaxCacheSizePercentage struct {
	// Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid
	// values 1 to 100
	MaximumCacheSizePercentage *int64 `json:"maximumCacheSizePercentage,omitempty"`

	// Fields inherited from DeliveryOptimizationMaxCacheSize

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationMaxCacheSizePercentage) DeliveryOptimizationMaxCacheSize() BaseDeliveryOptimizationMaxCacheSizeImpl {
	return BaseDeliveryOptimizationMaxCacheSizeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationMaxCacheSizePercentage{}

func (s DeliveryOptimizationMaxCacheSizePercentage) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationMaxCacheSizePercentage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationMaxCacheSizePercentage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationMaxCacheSizePercentage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationMaxCacheSizePercentage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationMaxCacheSizePercentage: %+v", err)
	}

	return encoded, nil
}
