package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationBandwidth = DeliveryOptimizationBandwidthPercentage{}

type DeliveryOptimizationBandwidthPercentage struct {
	// Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download
	// activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumBackgroundBandwidthPercentage nullable.Type[int64] `json:"maximumBackgroundBandwidthPercentage,omitempty"`

	// Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download
	// activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumForegroundBandwidthPercentage nullable.Type[int64] `json:"maximumForegroundBandwidthPercentage,omitempty"`

	// Fields inherited from DeliveryOptimizationBandwidth

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationBandwidthPercentage) DeliveryOptimizationBandwidth() BaseDeliveryOptimizationBandwidthImpl {
	return BaseDeliveryOptimizationBandwidthImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationBandwidthPercentage{}

func (s DeliveryOptimizationBandwidthPercentage) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationBandwidthPercentage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationBandwidthPercentage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationBandwidthPercentage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationBandwidthPercentage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationBandwidthPercentage: %+v", err)
	}

	return encoded, nil
}
