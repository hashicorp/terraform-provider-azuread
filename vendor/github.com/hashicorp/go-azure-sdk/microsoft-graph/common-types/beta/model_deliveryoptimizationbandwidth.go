package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationBandwidth interface {
	DeliveryOptimizationBandwidth() BaseDeliveryOptimizationBandwidthImpl
}

var _ DeliveryOptimizationBandwidth = BaseDeliveryOptimizationBandwidthImpl{}

type BaseDeliveryOptimizationBandwidthImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeliveryOptimizationBandwidthImpl) DeliveryOptimizationBandwidth() BaseDeliveryOptimizationBandwidthImpl {
	return s
}

var _ DeliveryOptimizationBandwidth = RawDeliveryOptimizationBandwidthImpl{}

// RawDeliveryOptimizationBandwidthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeliveryOptimizationBandwidthImpl struct {
	deliveryOptimizationBandwidth BaseDeliveryOptimizationBandwidthImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawDeliveryOptimizationBandwidthImpl) DeliveryOptimizationBandwidth() BaseDeliveryOptimizationBandwidthImpl {
	return s.deliveryOptimizationBandwidth
}

func UnmarshalDeliveryOptimizationBandwidthImplementation(input []byte) (DeliveryOptimizationBandwidth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationBandwidth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationBandwidthAbsolute") {
		var out DeliveryOptimizationBandwidthAbsolute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationBandwidthAbsolute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationBandwidthHoursWithPercentage") {
		var out DeliveryOptimizationBandwidthHoursWithPercentage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationBandwidthHoursWithPercentage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationBandwidthPercentage") {
		var out DeliveryOptimizationBandwidthPercentage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationBandwidthPercentage: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeliveryOptimizationBandwidthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeliveryOptimizationBandwidthImpl: %+v", err)
	}

	return RawDeliveryOptimizationBandwidthImpl{
		deliveryOptimizationBandwidth: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
