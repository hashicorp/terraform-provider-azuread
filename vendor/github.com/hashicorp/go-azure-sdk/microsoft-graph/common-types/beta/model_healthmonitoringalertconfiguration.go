package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HealthMonitoringAlertConfiguration{}

type HealthMonitoringAlertConfiguration struct {
	// Defines the recipients of email notifications for an alert type. Currently, only one email notification configuration
	// is supported for an alert configuration, meaning only one group can receive notifications for an alert type.
	EmailNotificationConfigurations *[]HealthMonitoringEmailNotificationConfiguration `json:"emailNotificationConfigurations,omitempty"`

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

func (s HealthMonitoringAlertConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HealthMonitoringAlertConfiguration{}

func (s HealthMonitoringAlertConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringAlertConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringAlertConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringAlertConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.alertConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringAlertConfiguration: %+v", err)
	}

	return encoded, nil
}
