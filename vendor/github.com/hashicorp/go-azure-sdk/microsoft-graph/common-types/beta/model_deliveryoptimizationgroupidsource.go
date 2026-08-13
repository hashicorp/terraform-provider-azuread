package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationGroupIdSource interface {
	DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl
}

var _ DeliveryOptimizationGroupIdSource = BaseDeliveryOptimizationGroupIdSourceImpl{}

type BaseDeliveryOptimizationGroupIdSourceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeliveryOptimizationGroupIdSourceImpl) DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl {
	return s
}

var _ DeliveryOptimizationGroupIdSource = RawDeliveryOptimizationGroupIdSourceImpl{}

// RawDeliveryOptimizationGroupIdSourceImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDeliveryOptimizationGroupIdSourceImpl struct {
	deliveryOptimizationGroupIdSource BaseDeliveryOptimizationGroupIdSourceImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawDeliveryOptimizationGroupIdSourceImpl) DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl {
	return s.deliveryOptimizationGroupIdSource
}

func (s RawDeliveryOptimizationGroupIdSourceImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDeliveryOptimizationGroupIdSourceImplementation(input []byte) (DeliveryOptimizationGroupIdSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationGroupIdSource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationGroupIdCustom") {
		var out DeliveryOptimizationGroupIdCustom
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationGroupIdCustom: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deliveryOptimizationGroupIdSourceOptions") {
		var out DeliveryOptimizationGroupIdSourceOptions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeliveryOptimizationGroupIdSourceOptions: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeliveryOptimizationGroupIdSourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeliveryOptimizationGroupIdSourceImpl: %+v", err)
	}

	return RawDeliveryOptimizationGroupIdSourceImpl{
		deliveryOptimizationGroupIdSource: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
