package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataIndustryDataRoot{}

type IndustryDataIndustryDataRoot struct {
	// Set of connectors for importing data from source systems.
	DataConnectors *[]IndustryDataIndustryDataConnector `json:"dataConnectors,omitempty"`

	// Set of data import flow activities to bring data into the canonical store via a connector.
	InboundFlows *[]IndustryDataInboundFlow `json:"inboundFlows,omitempty"`

	// Set of ephemeral operations that the system runs currently. Read-only.
	Operations *[]LongRunningOperation `json:"operations,omitempty"`

	OutboundProvisioningFlowSets *[]IndustryDataOutboundProvisioningFlowSet `json:"outboundProvisioningFlowSets,omitempty"`

	// Set of user modifiable system picker types.
	ReferenceDefinitions *[]IndustryDataReferenceDefinition `json:"referenceDefinitions,omitempty"`

	// Set of groups of individual roles that makes role-based admin simpler.
	RoleGroups *[]IndustryDataRoleGroup `json:"roleGroups,omitempty"`

	// Set of ephemeral runs which present the point-in-time that diagnostic state of activities performed by the system.
	// Read-only.
	Runs *[]IndustryDataIndustryDataRun `json:"runs,omitempty"`

	// Set of source definitions that represents real-world external systems.
	SourceSystems *[]IndustryDataSourceSystemDefinition `json:"sourceSystems,omitempty"`

	// Set of years represented in the system.
	Years *[]IndustryDataYearTimePeriodDefinition `json:"years,omitempty"`

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

func (s IndustryDataIndustryDataRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataIndustryDataRoot{}

func (s IndustryDataIndustryDataRoot) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIndustryDataRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIndustryDataRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRoot: %+v", err)
	}

	delete(decoded, "operations")
	delete(decoded, "runs")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.industryDataRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIndustryDataRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataIndustryDataRoot{}

func (s *IndustryDataIndustryDataRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		OutboundProvisioningFlowSets *[]IndustryDataOutboundProvisioningFlowSet `json:"outboundProvisioningFlowSets,omitempty"`
		ReferenceDefinitions         *[]IndustryDataReferenceDefinition         `json:"referenceDefinitions,omitempty"`
		RoleGroups                   *[]IndustryDataRoleGroup                   `json:"roleGroups,omitempty"`
		Runs                         *[]IndustryDataIndustryDataRun             `json:"runs,omitempty"`
		SourceSystems                *[]IndustryDataSourceSystemDefinition      `json:"sourceSystems,omitempty"`
		Years                        *[]IndustryDataYearTimePeriodDefinition    `json:"years,omitempty"`
		Id                           *string                                    `json:"id,omitempty"`
		ODataId                      *string                                    `json:"@odata.id,omitempty"`
		ODataType                    *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.OutboundProvisioningFlowSets = decoded.OutboundProvisioningFlowSets
	s.ReferenceDefinitions = decoded.ReferenceDefinitions
	s.RoleGroups = decoded.RoleGroups
	s.Runs = decoded.Runs
	s.SourceSystems = decoded.SourceSystems
	s.Years = decoded.Years
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataIndustryDataRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataConnectors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DataConnectors into list []json.RawMessage: %+v", err)
		}

		output := make([]IndustryDataIndustryDataConnector, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIndustryDataIndustryDataConnectorImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DataConnectors' for 'IndustryDataIndustryDataRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DataConnectors = &output
	}

	if v, ok := temp["inboundFlows"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InboundFlows into list []json.RawMessage: %+v", err)
		}

		output := make([]IndustryDataInboundFlow, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIndustryDataInboundFlowImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InboundFlows' for 'IndustryDataIndustryDataRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InboundFlows = &output
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]LongRunningOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLongRunningOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'IndustryDataIndustryDataRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
