package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerDelta = PlannerPlan{}

type PlannerPlan struct {
	// Read-only. Nullable. Contains information about who archived or unarchived the plan and why.
	ArchivalInfo *PlannerArchivalInfo `json:"archivalInfo,omitempty"`

	// Collection of buckets in the plan. Read-only. Nullable.
	Buckets *[]PlannerBucket `json:"buckets,omitempty"`

	// Identifies the container of the plan. Either specify all properties, or specify only the url, the containerId, and
	// type. After it's set, this property can’t be updated. It changes when a plan is moved from one container to
	// another, using plan move to container. Required.
	Container PlannerPlanContainer `json:"container"`

	// Read-only. Other user experiences in which this plan is used, represented as plannerPlanContext entries.
	Contexts *PlannerPlanContextCollection `json:"contexts,omitempty"`

	// Read-only. The user who created the plan.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains information about the origin of the plan.
	CreationSource PlannerPlanCreation `json:"creationSource"`

	// Extra details about the plan. Read-only. Nullable.
	Details *PlannerPlanDetails `json:"details,omitempty"`

	// Read-only. If set to true, the plan is archived. An archived plan is read-only.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Use the container property instead. ID of the group that owns the plan. After it's set, this property can’t be
	// updated. This property doesn't return a valid group ID if the container of the plan isn't a group.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// List of containers the plan is shared with.
	SharedWithContainers *[]PlannerSharedWithContainer `json:"sharedWithContainers,omitempty"`

	// Collection of tasks in the plan. Read-only. Nullable.
	Tasks *[]PlannerTask `json:"tasks,omitempty"`

	// Required. Title of the plan.
	Title string `json:"title"`

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

func (s PlannerPlan) PlannerDelta() BasePlannerDeltaImpl {
	return BasePlannerDeltaImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s PlannerPlan) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerPlan{}

func (s PlannerPlan) MarshalJSON() ([]byte, error) {
	type wrapper PlannerPlan
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerPlan: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlan: %+v", err)
	}

	delete(decoded, "archivalInfo")
	delete(decoded, "buckets")
	delete(decoded, "contexts")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "details")
	delete(decoded, "isArchived")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerPlan"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerPlan: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerPlan{}

func (s *PlannerPlan) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ArchivalInfo         *PlannerArchivalInfo          `json:"archivalInfo,omitempty"`
		Buckets              *[]PlannerBucket              `json:"buckets,omitempty"`
		Contexts             *PlannerPlanContextCollection `json:"contexts,omitempty"`
		CreatedDateTime      nullable.Type[string]         `json:"createdDateTime,omitempty"`
		Details              *PlannerPlanDetails           `json:"details,omitempty"`
		IsArchived           nullable.Type[bool]           `json:"isArchived,omitempty"`
		Owner                nullable.Type[string]         `json:"owner,omitempty"`
		SharedWithContainers *[]PlannerSharedWithContainer `json:"sharedWithContainers,omitempty"`
		Title                string                        `json:"title"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ArchivalInfo = decoded.ArchivalInfo
	s.Buckets = decoded.Buckets
	s.Contexts = decoded.Contexts
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Details = decoded.Details
	s.IsArchived = decoded.IsArchived
	s.Owner = decoded.Owner
	s.SharedWithContainers = decoded.SharedWithContainers
	s.Title = decoded.Title
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerPlan into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["container"]; ok {
		impl, err := UnmarshalPlannerPlanContainerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Container' for 'PlannerPlan': %+v", err)
		}
		s.Container = impl
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PlannerPlan': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["creationSource"]; ok {
		impl, err := UnmarshalPlannerPlanCreationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreationSource' for 'PlannerPlan': %+v", err)
		}
		s.CreationSource = impl
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
				return fmt.Errorf("unmarshaling index %d field 'Tasks' for 'PlannerPlan': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Tasks = &output
	}

	return nil
}
