package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataOutboundProvisioningFlowSet{}

type IndustryDataOutboundProvisioningFlowSet struct {
	// The date and time when the flowSet was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name of the flowSet provided by the caller.
	DisplayName *string `json:"displayName,omitempty"`

	// The collection of provisioning filters applicable to all the flows under the given flowSet.
	Filter IndustryDataFilter `json:"filter"`

	// The date and time when the flowSet was most recently changed. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// A flow that provisions relevant records of a given entity type in the Microsoft 365 tenant.
	ProvisioningFlows *[]IndustryDataProvisioningFlow `json:"provisioningFlows,omitempty"`

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

func (s IndustryDataOutboundProvisioningFlowSet) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataOutboundProvisioningFlowSet{}

func (s IndustryDataOutboundProvisioningFlowSet) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataOutboundProvisioningFlowSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataOutboundProvisioningFlowSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataOutboundProvisioningFlowSet: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.outboundProvisioningFlowSet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataOutboundProvisioningFlowSet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataOutboundProvisioningFlowSet{}

func (s *IndustryDataOutboundProvisioningFlowSet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		DisplayName          *string               `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataOutboundProvisioningFlowSet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["filter"]; ok {
		impl, err := UnmarshalIndustryDataFilterImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Filter' for 'IndustryDataOutboundProvisioningFlowSet': %+v", err)
		}
		s.Filter = impl
	}

	if v, ok := temp["provisioningFlows"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProvisioningFlows into list []json.RawMessage: %+v", err)
		}

		output := make([]IndustryDataProvisioningFlow, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIndustryDataProvisioningFlowImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProvisioningFlows' for 'IndustryDataOutboundProvisioningFlowSet': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProvisioningFlows = &output
	}

	return nil
}
