package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskConfigurationRoleBase interface {
	PlannerTaskConfigurationRoleBase() BasePlannerTaskConfigurationRoleBaseImpl
}

var _ PlannerTaskConfigurationRoleBase = BasePlannerTaskConfigurationRoleBaseImpl{}

type BasePlannerTaskConfigurationRoleBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RoleKind *PlannerUserRoleKind `json:"roleKind,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerTaskConfigurationRoleBaseImpl) PlannerTaskConfigurationRoleBase() BasePlannerTaskConfigurationRoleBaseImpl {
	return s
}

var _ PlannerTaskConfigurationRoleBase = RawPlannerTaskConfigurationRoleBaseImpl{}

// RawPlannerTaskConfigurationRoleBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawPlannerTaskConfigurationRoleBaseImpl struct {
	plannerTaskConfigurationRoleBase BasePlannerTaskConfigurationRoleBaseImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawPlannerTaskConfigurationRoleBaseImpl) PlannerTaskConfigurationRoleBase() BasePlannerTaskConfigurationRoleBaseImpl {
	return s.plannerTaskConfigurationRoleBase
}

func (s RawPlannerTaskConfigurationRoleBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalPlannerTaskConfigurationRoleBaseImplementation(input []byte) (PlannerTaskConfigurationRoleBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTaskConfigurationRoleBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerRelationshipBasedUserType") {
		var out PlannerRelationshipBasedUserType
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerRelationshipBasedUserType: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerTaskConfigurationRoleBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerTaskConfigurationRoleBaseImpl: %+v", err)
	}

	return RawPlannerTaskConfigurationRoleBaseImpl{
		plannerTaskConfigurationRoleBase: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
