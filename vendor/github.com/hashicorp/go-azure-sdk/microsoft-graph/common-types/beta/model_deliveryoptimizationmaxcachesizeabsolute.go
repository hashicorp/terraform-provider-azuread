package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationMaxCacheSize = DeliveryOptimizationMaxCacheSizeAbsolute{}

type DeliveryOptimizationMaxCacheSizeAbsolute struct {
	// Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
	MaximumCacheSizeInGigabytes *int64 `json:"maximumCacheSizeInGigabytes,omitempty"`

	// Fields inherited from DeliveryOptimizationMaxCacheSize

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationMaxCacheSizeAbsolute) DeliveryOptimizationMaxCacheSize() BaseDeliveryOptimizationMaxCacheSizeImpl {
	return BaseDeliveryOptimizationMaxCacheSizeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationMaxCacheSizeAbsolute{}

func (s DeliveryOptimizationMaxCacheSizeAbsolute) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationMaxCacheSizeAbsolute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationMaxCacheSizeAbsolute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationMaxCacheSizeAbsolute: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationMaxCacheSizeAbsolute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationMaxCacheSizeAbsolute: %+v", err)
	}

	return encoded, nil
}
