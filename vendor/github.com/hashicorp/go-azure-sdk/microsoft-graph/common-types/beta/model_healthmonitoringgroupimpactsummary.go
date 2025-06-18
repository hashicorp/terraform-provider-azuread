package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ HealthMonitoringDirectoryObjectImpactSummary = HealthMonitoringGroupImpactSummary{}

type HealthMonitoringGroupImpactSummary struct {

	// Fields inherited from HealthMonitoringDirectoryObjectImpactSummary

	ResourceSampling *[]DirectoryObject `json:"resourceSampling,omitempty"`

	// List of OData IDs for `ResourceSampling` to bind to this entity
	ResourceSampling_ODataBind *[]string `json:"resourceSampling@odata.bind,omitempty"`

	// Fields inherited from HealthMonitoringResourceImpactSummary

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

func (s HealthMonitoringGroupImpactSummary) HealthMonitoringDirectoryObjectImpactSummary() BaseHealthMonitoringDirectoryObjectImpactSummaryImpl {
	return BaseHealthMonitoringDirectoryObjectImpactSummaryImpl{
		ResourceSampling:           s.ResourceSampling,
		ResourceSampling_ODataBind: s.ResourceSampling_ODataBind,
		ImpactedCount:              s.ImpactedCount,
		ImpactedCountLimitExceeded: s.ImpactedCountLimitExceeded,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
		ResourceType:               s.ResourceType,
	}
}

func (s HealthMonitoringGroupImpactSummary) HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl {
	return BaseHealthMonitoringResourceImpactSummaryImpl{
		ImpactedCount:              s.ImpactedCount,
		ImpactedCountLimitExceeded: s.ImpactedCountLimitExceeded,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
		ResourceType:               s.ResourceType,
	}
}

var _ json.Marshaler = HealthMonitoringGroupImpactSummary{}

func (s HealthMonitoringGroupImpactSummary) MarshalJSON() ([]byte, error) {
	type wrapper HealthMonitoringGroupImpactSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HealthMonitoringGroupImpactSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringGroupImpactSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.groupImpactSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HealthMonitoringGroupImpactSummary: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &HealthMonitoringGroupImpactSummary{}

func (s *HealthMonitoringGroupImpactSummary) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ResourceSampling_ODataBind *[]string `json:"resourceSampling@odata.bind,omitempty"`
		ImpactedCount              *string   `json:"impactedCount,omitempty"`
		ImpactedCountLimitExceeded *bool     `json:"impactedCountLimitExceeded,omitempty"`
		ODataId                    *string   `json:"@odata.id,omitempty"`
		ODataType                  *string   `json:"@odata.type,omitempty"`
		ResourceType               *string   `json:"resourceType,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ImpactedCount = decoded.ImpactedCount
	s.ImpactedCountLimitExceeded = decoded.ImpactedCountLimitExceeded
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceSampling_ODataBind = decoded.ResourceSampling_ODataBind
	s.ResourceType = decoded.ResourceType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HealthMonitoringGroupImpactSummary into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resourceSampling"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResourceSampling into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResourceSampling' for 'HealthMonitoringGroupImpactSummary': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceSampling = &output
	}

	return nil
}
