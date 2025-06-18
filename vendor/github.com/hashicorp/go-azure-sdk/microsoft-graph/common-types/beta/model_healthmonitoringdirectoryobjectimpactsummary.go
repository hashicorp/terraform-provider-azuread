package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringDirectoryObjectImpactSummary interface {
	HealthMonitoringResourceImpactSummary
	HealthMonitoringDirectoryObjectImpactSummary() BaseHealthMonitoringDirectoryObjectImpactSummaryImpl
}

var _ HealthMonitoringDirectoryObjectImpactSummary = BaseHealthMonitoringDirectoryObjectImpactSummaryImpl{}

type BaseHealthMonitoringDirectoryObjectImpactSummaryImpl struct {
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

func (s BaseHealthMonitoringDirectoryObjectImpactSummaryImpl) HealthMonitoringDirectoryObjectImpactSummary() BaseHealthMonitoringDirectoryObjectImpactSummaryImpl {
	return s
}

func (s BaseHealthMonitoringDirectoryObjectImpactSummaryImpl) HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl {
	return BaseHealthMonitoringResourceImpactSummaryImpl{
		ImpactedCount:              s.ImpactedCount,
		ImpactedCountLimitExceeded: s.ImpactedCountLimitExceeded,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
		ResourceType:               s.ResourceType,
	}
}

var _ HealthMonitoringDirectoryObjectImpactSummary = RawHealthMonitoringDirectoryObjectImpactSummaryImpl{}

// RawHealthMonitoringDirectoryObjectImpactSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawHealthMonitoringDirectoryObjectImpactSummaryImpl struct {
	healthMonitoringDirectoryObjectImpactSummary BaseHealthMonitoringDirectoryObjectImpactSummaryImpl
	Type                                         string
	Values                                       map[string]interface{}
}

func (s RawHealthMonitoringDirectoryObjectImpactSummaryImpl) HealthMonitoringDirectoryObjectImpactSummary() BaseHealthMonitoringDirectoryObjectImpactSummaryImpl {
	return s.healthMonitoringDirectoryObjectImpactSummary
}

func (s RawHealthMonitoringDirectoryObjectImpactSummaryImpl) HealthMonitoringResourceImpactSummary() BaseHealthMonitoringResourceImpactSummaryImpl {
	return s.healthMonitoringDirectoryObjectImpactSummary.HealthMonitoringResourceImpactSummary()
}

var _ json.Marshaler = BaseHealthMonitoringDirectoryObjectImpactSummaryImpl{}

func (s BaseHealthMonitoringDirectoryObjectImpactSummaryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseHealthMonitoringDirectoryObjectImpactSummaryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseHealthMonitoringDirectoryObjectImpactSummaryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseHealthMonitoringDirectoryObjectImpactSummaryImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.directoryObjectImpactSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseHealthMonitoringDirectoryObjectImpactSummaryImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseHealthMonitoringDirectoryObjectImpactSummaryImpl{}

func (s *BaseHealthMonitoringDirectoryObjectImpactSummaryImpl) UnmarshalJSON(bytes []byte) error {
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

	s.ResourceSampling_ODataBind = decoded.ResourceSampling_ODataBind
	s.ImpactedCount = decoded.ImpactedCount
	s.ImpactedCountLimitExceeded = decoded.ImpactedCountLimitExceeded
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ResourceType = decoded.ResourceType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseHealthMonitoringDirectoryObjectImpactSummaryImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'ResourceSampling' for 'BaseHealthMonitoringDirectoryObjectImpactSummaryImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceSampling = &output
	}

	return nil
}

func UnmarshalHealthMonitoringDirectoryObjectImpactSummaryImplementation(input []byte) (HealthMonitoringDirectoryObjectImpactSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringDirectoryObjectImpactSummary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.applicationImpactSummary") {
		var out HealthMonitoringApplicationImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringApplicationImpactSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.deviceImpactSummary") {
		var out HealthMonitoringDeviceImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringDeviceImpactSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.groupImpactSummary") {
		var out HealthMonitoringGroupImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringGroupImpactSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.servicePrincipalImpactSummary") {
		var out HealthMonitoringServicePrincipalImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringServicePrincipalImpactSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.userImpactSummary") {
		var out HealthMonitoringUserImpactSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringUserImpactSummary: %+v", err)
		}
		return out, nil
	}

	var parent BaseHealthMonitoringDirectoryObjectImpactSummaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseHealthMonitoringDirectoryObjectImpactSummaryImpl: %+v", err)
	}

	return RawHealthMonitoringDirectoryObjectImpactSummaryImpl{
		healthMonitoringDirectoryObjectImpactSummary: parent,
		Type:   value,
		Values: temp,
	}, nil

}
