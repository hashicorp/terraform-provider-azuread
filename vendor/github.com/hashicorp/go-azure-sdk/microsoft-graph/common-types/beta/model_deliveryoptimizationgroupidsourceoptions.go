package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationGroupIdSource = DeliveryOptimizationGroupIdSourceOptions{}

type DeliveryOptimizationGroupIdSourceOptions struct {
	// Possible values for the DeliveryOptimizationGroupIdOptionsType setting.
	GroupIdSourceOption *DeliveryOptimizationGroupIdOptionsType `json:"groupIdSourceOption,omitempty"`

	// Fields inherited from DeliveryOptimizationGroupIdSource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationGroupIdSourceOptions) DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl {
	return BaseDeliveryOptimizationGroupIdSourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationGroupIdSourceOptions{}

func (s DeliveryOptimizationGroupIdSourceOptions) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationGroupIdSourceOptions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationGroupIdSourceOptions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationGroupIdSourceOptions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationGroupIdSourceOptions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationGroupIdSourceOptions: %+v", err)
	}

	return encoded, nil
}
