package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecuritySensor{}

type SecuritySensor struct {
	CreatedDateTime       *string                     `json:"createdDateTime,omitempty"`
	DeploymentStatus      *SecurityDeploymentStatus   `json:"deploymentStatus,omitempty"`
	DisplayName           *string                     `json:"displayName,omitempty"`
	DomainName            *string                     `json:"domainName,omitempty"`
	HealthIssues          *[]SecurityHealthIssue      `json:"healthIssues,omitempty"`
	HealthStatus          *SecuritySensorHealthStatus `json:"healthStatus,omitempty"`
	OpenHealthIssuesCount *int64                      `json:"openHealthIssuesCount,omitempty"`
	SensorType            *SecuritySensorType         `json:"sensorType,omitempty"`
	Settings              *SecuritySensorSettings     `json:"settings,omitempty"`
	Version               *string                     `json:"version,omitempty"`

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

func (s SecuritySensor) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecuritySensor{}

func (s SecuritySensor) MarshalJSON() ([]byte, error) {
	type wrapper SecuritySensor
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecuritySensor: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecuritySensor: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.sensor"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecuritySensor: %+v", err)
	}

	return encoded, nil
}
