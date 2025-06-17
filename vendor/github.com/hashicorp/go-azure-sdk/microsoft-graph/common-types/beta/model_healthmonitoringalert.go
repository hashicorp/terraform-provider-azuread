package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HealthMonitoringAlert{}

type HealthMonitoringAlert struct {
	AlertType *HealthMonitoringAlertType `json:"alertType,omitempty"`
	Category  *HealthMonitoringCategory  `json:"category,omitempty"`

	// The time when Microsoft Entra Health monitoring generated the alert. Supports $orderby.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A key-value pair that contains the name of and link to the documentation to aid in investigation of the alert.
	Documentation *HealthMonitoringDocumentation `json:"documentation,omitempty"`

	// Investigative information on the alert. This information typically includes counts of impacted objects, which include
	// directory objects such as users, groups, and devices, and a pointer to supporting data.
	Enrichment *HealthMonitoringEnrichment `json:"enrichment,omitempty"`

	Scenario *HealthMonitoringScenario `json:"scenario,omitempty"`

	// The collection of signals that were used in the generation of the alert. These signals are sourced from
	// serviceActivity APIs and are added to the alert as key-value pairs.
	Signals *HealthMonitoringSignals `json:"signals,omitempty"`

	State *HealthMonitoringAlertState `json:"state,omitempty"`

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

func (s HealthMonitoringAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HealthMonitoringAlert{}

func (s HealthMonitoringAlert) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.alert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringAlert: %+v", err)
	}

	return encoded, nil
}
