package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

// RawDeliveryOptimizationGroupIdSourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeliveryOptimizationGroupIdSourceImpl struct {
	deliveryOptimizationGroupIdSource BaseDeliveryOptimizationGroupIdSourceImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawDeliveryOptimizationGroupIdSourceImpl) DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl {
	return s.deliveryOptimizationGroupIdSource
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
