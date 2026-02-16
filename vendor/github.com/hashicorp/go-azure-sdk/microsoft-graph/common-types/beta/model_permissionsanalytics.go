package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PermissionsAnalytics{}

type PermissionsAnalytics struct {
	// The output of the permissions usage data analysis performed by Permissions Management to assess risk with identities
	// and resources.
	Findings *[]Finding `json:"findings,omitempty"`

	// Represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the
	// classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
	PermissionsCreepIndexDistributions *[]PermissionsCreepIndexDistribution `json:"permissionsCreepIndexDistributions,omitempty"`

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

func (s PermissionsAnalytics) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PermissionsAnalytics{}

func (s PermissionsAnalytics) MarshalJSON() ([]byte, error) {
	type wrapper PermissionsAnalytics
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionsAnalytics: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsAnalytics: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.permissionsAnalytics"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionsAnalytics: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PermissionsAnalytics{}

func (s *PermissionsAnalytics) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		PermissionsCreepIndexDistributions *[]PermissionsCreepIndexDistribution `json:"permissionsCreepIndexDistributions,omitempty"`
		Id                                 *string                              `json:"id,omitempty"`
		ODataId                            *string                              `json:"@odata.id,omitempty"`
		ODataType                          *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.PermissionsCreepIndexDistributions = decoded.PermissionsCreepIndexDistributions
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PermissionsAnalytics into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["findings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Findings into list []json.RawMessage: %+v", err)
		}

		output := make([]Finding, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalFindingImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Findings' for 'PermissionsAnalytics': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Findings = &output
	}

	return nil
}
