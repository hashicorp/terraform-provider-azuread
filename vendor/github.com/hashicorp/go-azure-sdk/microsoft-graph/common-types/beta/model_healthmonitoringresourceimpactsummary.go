package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringResourceImpactSummary interface {
	HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl
}

var _ HealthMonitoringResourceImpactSummary = BaseHealthMonitoringResourceImpactSummaryImpl{}

type BaseHealthMonitoringResourceImpactSummaryImpl struct {
	// The number of resources impacted. The number could be an exhaustive count or a sampling count.
	ImpactedCount *string `json:"impactedCount,omitempty"`

	// Indicates whether impactedCount is exhaustive or a sampling. When this value is true, the limit was exceeded and
	// impactedCount represents a sampling; otherwise, impactedCount represents the true number of impacts.
	ImpactedCountLimitExceeded *bool `json:"impactedCountLimitExceeded,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of resource that was impacted. Examples include user, group, application, servicePrincipal, device.
	ResourceType *string `json:"resourceType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseHealthMonitoringResourceImpactSummaryImpl) HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl {
	return s
}

var _ HealthMonitoringResourceImpactSummary = RawHealthMonitoringResourceImpactSummaryImpl{}

// RawHealthMonitoringResourceImpactSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawHealthMonitoringResourceImpactSummaryImpl struct {
	healthMonitoringResourceImpactSummary BaseHealthMonitoringResourceImpactSummaryImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawHealthMonitoringResourceImpactSummaryImpl) HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl {
	return s.healthMonitoringResourceImpactSummary
}

func UnmarshalHealthMonitoringResourceImpactSummaryImplementation(input []byte) (HealthMonitoringResourceImpactSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringResourceImpactSummary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.directoryObjectImpactSummary") {
		var out HealthMonitoringDirectoryObjectImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringDirectoryObjectImpactSummary: %+v", err)
		}
		return out, nil
	}

	var parent BaseHealthMonitoringResourceImpactSummaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseHealthMonitoringResourceImpactSummaryImpl: %+v", err)
	}

	return RawHealthMonitoringResourceImpactSummaryImpl{
		healthMonitoringResourceImpactSummary: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
