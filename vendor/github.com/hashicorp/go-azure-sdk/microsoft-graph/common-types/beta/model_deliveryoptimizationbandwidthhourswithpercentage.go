package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationBandwidth = DeliveryOptimizationBandwidthHoursWithPercentage{}

type DeliveryOptimizationBandwidthHoursWithPercentage struct {
	// Background download percentage hours.
	BandwidthBackgroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthBackgroundPercentageHours,omitempty"`

	// Foreground download percentage hours.
	BandwidthForegroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthForegroundPercentageHours,omitempty"`

	// Fields inherited from DeliveryOptimizationBandwidth

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationBandwidthHoursWithPercentage) DeliveryOptimizationBandwidth() BaseDeliveryOptimizationBandwidthImpl {
	return BaseDeliveryOptimizationBandwidthImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationBandwidthHoursWithPercentage{}

func (s DeliveryOptimizationBandwidthHoursWithPercentage) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationBandwidthHoursWithPercentage
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationBandwidthHoursWithPercentage: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationBandwidthHoursWithPercentage: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationBandwidthHoursWithPercentage"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationBandwidthHoursWithPercentage: %+v", err)
	}

	return encoded, nil
}
