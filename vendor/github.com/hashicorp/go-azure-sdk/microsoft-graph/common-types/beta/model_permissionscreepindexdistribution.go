package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PermissionsCreepIndexDistribution{}

type PermissionsCreepIndexDistribution struct {
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`

	// Defines when the PCI distribution was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	HighRiskProfile   *RiskProfile `json:"highRiskProfile,omitempty"`
	LowRiskProfile    *RiskProfile `json:"lowRiskProfile,omitempty"`
	MediumRiskProfile *RiskProfile `json:"mediumRiskProfile,omitempty"`

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

func (s PermissionsCreepIndexDistribution) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PermissionsCreepIndexDistribution{}

func (s PermissionsCreepIndexDistribution) MarshalJSON() ([]byte, error) {
	type wrapper PermissionsCreepIndexDistribution
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PermissionsCreepIndexDistribution: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PermissionsCreepIndexDistribution: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.permissionsCreepIndexDistribution"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PermissionsCreepIndexDistribution: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PermissionsCreepIndexDistribution{}

func (s *PermissionsCreepIndexDistribution) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime   *string      `json:"createdDateTime,omitempty"`
		HighRiskProfile   *RiskProfile `json:"highRiskProfile,omitempty"`
		LowRiskProfile    *RiskProfile `json:"lowRiskProfile,omitempty"`
		MediumRiskProfile *RiskProfile `json:"mediumRiskProfile,omitempty"`
		Id                *string      `json:"id,omitempty"`
		ODataId           *string      `json:"@odata.id,omitempty"`
		ODataType         *string      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.HighRiskProfile = decoded.HighRiskProfile
	s.LowRiskProfile = decoded.LowRiskProfile
	s.MediumRiskProfile = decoded.MediumRiskProfile
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PermissionsCreepIndexDistribution into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["authorizationSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuthorizationSystem' for 'PermissionsCreepIndexDistribution': %+v", err)
		}
		s.AuthorizationSystem = &impl
	}

	return nil
}
