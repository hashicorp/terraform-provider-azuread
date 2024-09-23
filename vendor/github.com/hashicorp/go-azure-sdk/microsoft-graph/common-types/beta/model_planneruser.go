package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerDelta = PlannerUser{}

type PlannerUser struct {
	All *[]PlannerDelta `json:"all,omitempty"`

	// A collection that contains the references to the plans that the user marked as favorites.
	FavoritePlanReferences *PlannerFavoritePlanReferenceCollection `json:"favoritePlanReferences,omitempty"`

	// Read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
	FavoritePlans *[]PlannerPlan `json:"favoritePlans,omitempty"`

	// Read-only. Nullable. Returns the plannerTasks to be shown in the My Day view of the user.
	MyDayTasks *[]PlannerTask `json:"myDayTasks,omitempty"`

	Plans *[]PlannerPlan `json:"plans,omitempty"`

	// A collection that contains references to the plans that the user recently viewed in apps that support recent plans.
	RecentPlanReferences *PlannerRecentPlanReferenceCollection `json:"recentPlanReferences,omitempty"`

	// Read-only. Nullable. Returns the plannerPlans that the user recently viewed in apps that support recent plans.
	RecentPlans *[]PlannerPlan `json:"recentPlans,omitempty"`

	// Read-only. Nullable. Returns the plannerPlans contained by the plannerRosters the user is a member.
	RosterPlans *[]PlannerPlan `json:"rosterPlans,omitempty"`

	// Read-only. Nullable. Returns the plannerTasks assigned to the user.
	Tasks *[]PlannerTask `json:"tasks,omitempty"`

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

func (s PlannerUser) PlannerDelta() BasePlannerDeltaImpl {
	return BasePlannerDeltaImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s PlannerUser) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerUser{}

func (s PlannerUser) MarshalJSON() ([]byte, error) {
	type wrapper PlannerUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerUser: %+v", err)
	}

	delete(decoded, "favoritePlans")
	delete(decoded, "myDayTasks")
	delete(decoded, "recentPlans")
	delete(decoded, "rosterPlans")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerUser: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerUser{}

func (s *PlannerUser) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FavoritePlanReferences *PlannerFavoritePlanReferenceCollection `json:"favoritePlanReferences,omitempty"`
		FavoritePlans          *[]PlannerPlan                          `json:"favoritePlans,omitempty"`
		Plans                  *[]PlannerPlan                          `json:"plans,omitempty"`
		RecentPlanReferences   *PlannerRecentPlanReferenceCollection   `json:"recentPlanReferences,omitempty"`
		RecentPlans            *[]PlannerPlan                          `json:"recentPlans,omitempty"`
		RosterPlans            *[]PlannerPlan                          `json:"rosterPlans,omitempty"`
		Id                     *string                                 `json:"id,omitempty"`
		ODataId                *string                                 `json:"@odata.id,omitempty"`
		ODataType              *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FavoritePlanReferences = decoded.FavoritePlanReferences
	s.FavoritePlans = decoded.FavoritePlans
	s.Plans = decoded.Plans
	s.RecentPlanReferences = decoded.RecentPlanReferences
	s.RecentPlans = decoded.RecentPlans
	s.RosterPlans = decoded.RosterPlans
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerUser into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["all"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling All into list []json.RawMessage: %+v", err)
		}

		output := make([]PlannerDelta, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPlannerDeltaImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'All' for 'PlannerUser': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.All = &output
	}

	if v, ok := temp["myDayTasks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MyDayTasks into list []json.RawMessage: %+v", err)
		}

		output := make([]PlannerTask, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPlannerTaskImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MyDayTasks' for 'PlannerUser': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MyDayTasks = &output
	}

	if v, ok := temp["tasks"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Tasks into list []json.RawMessage: %+v", err)
		}

		output := make([]PlannerTask, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPlannerTaskImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Tasks' for 'PlannerUser': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Tasks = &output
	}

	return nil
}
