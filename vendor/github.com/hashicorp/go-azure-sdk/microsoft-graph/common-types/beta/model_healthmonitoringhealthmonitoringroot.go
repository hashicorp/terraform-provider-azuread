package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HealthMonitoringHealthMonitoringRoot{}

type HealthMonitoringHealthMonitoringRoot struct {
	// The configuration of an alert type, which defines behavior that occurs when an alert is created.
	AlertConfigurations *[]HealthMonitoringAlertConfiguration `json:"alertConfigurations,omitempty"`

	// The collection of health monitoring system detected alerts for anomalous usage patterns found in a Microsoft Entra
	// tenant.
	Alerts *[]HealthMonitoringAlert `json:"alerts,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s HealthMonitoringHealthMonitoringRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HealthMonitoringHealthMonitoringRoot{}

func (s HealthMonitoringHealthMonitoringRoot) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringHealthMonitoringRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringHealthMonitoringRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringHealthMonitoringRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.healthMonitoringRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringHealthMonitoringRoot: %+v", err)
	}

	return encoded, nil
}
