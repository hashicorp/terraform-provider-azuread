package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ HealthMonitoringHealthMonitoringDictionary = HealthMonitoringDocumentation{}

type HealthMonitoringDocumentation struct {

	// Fields inherited from HealthMonitoringDictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s HealthMonitoringDocumentation) HealthMonitoringHealthMonitoringDictionary() BaseHealthMonitoringHealthMonitoringDictionaryImpl {
	return BaseHealthMonitoringHealthMonitoringDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s HealthMonitoringDocumentation) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return BaseHealthMonitoringDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HealthMonitoringDocumentation{}

func (s HealthMonitoringDocumentation) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringDocumentation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringDocumentation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringDocumentation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.documentation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringDocumentation: %+v", err)
	}

	return encoded, nil
}
