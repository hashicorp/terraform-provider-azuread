package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryOptimizationGroupIdSource = DeliveryOptimizationGroupIdCustom{}

type DeliveryOptimizationGroupIdCustom struct {
	// Specifies an arbitrary group ID that the device belongs to
	GroupIdCustom *string `json:"groupIdCustom,omitempty"`

	// Fields inherited from DeliveryOptimizationGroupIdSource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeliveryOptimizationGroupIdCustom) DeliveryOptimizationGroupIdSource() BaseDeliveryOptimizationGroupIdSourceImpl {
	return BaseDeliveryOptimizationGroupIdSourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeliveryOptimizationGroupIdCustom{}

func (s DeliveryOptimizationGroupIdCustom) MarshalJSON() ([]byte, error) {
	type wrapper DeliveryOptimizationGroupIdCustom
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeliveryOptimizationGroupIdCustom: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeliveryOptimizationGroupIdCustom: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deliveryOptimizationGroupIdCustom"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeliveryOptimizationGroupIdCustom: %+v", err)
	}

	return encoded, nil
}
