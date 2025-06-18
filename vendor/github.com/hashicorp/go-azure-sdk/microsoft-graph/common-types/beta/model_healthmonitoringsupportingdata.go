package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ HealthMonitoringHealthMonitoringDictionary = HealthMonitoringSupportingData{}

type HealthMonitoringSupportingData struct {

	// Fields inherited from HealthMonitoringDictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s HealthMonitoringSupportingData) HealthMonitoringHealthMonitoringDictionary() BaseHealthMonitoringHealthMonitoringDictionaryImpl {
	return BaseHealthMonitoringHealthMonitoringDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s HealthMonitoringSupportingData) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return BaseHealthMonitoringDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HealthMonitoringSupportingData{}

func (s HealthMonitoringSupportingData) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringSupportingData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringSupportingData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringSupportingData: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.supportingData"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringSupportingData: %+v", err)
	}

	return encoded, nil
}
