package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerPlan{}

type PlannerPlan struct {
	// Read-only. Nullable. Collection of buckets in the plan.
	Buckets *[]PlannerBucket `json:"buckets,omitempty"`

	// Identifies the container of the plan. Specify only the url, the containerId and type, or all properties. After it's
	// set, this property can’t be updated. Required.
	Container PlannerPlanContainer `json:"container"`

	// Read-only. The user who created the plan.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Read-only. Nullable. Extra details about the plan.
	Details *PlannerPlanDetails `json:"details,omitempty"`

	// Use the container property instead. ID of the group that owns the plan. After it's set, this property can’t be
	// updated. This property won't return a valid group ID if the container of the plan isn't a group.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// Read-only. Nullable. Collection of tasks in the plan.
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

	delete(decoded, "buckets")
	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "details")
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
		Buckets         *[]PlannerBucket      `json:"buckets,omitempty"`
		Container       PlannerPlanContainer  `json:"container"`
		CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`
		Details         *PlannerPlanDetails   `json:"details,omitempty"`
		Owner           nullable.Type[string] `json:"owner,omitempty"`
		Tasks           *[]PlannerTask        `json:"tasks,omitempty"`
		Title           string                `json:"title"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Buckets = decoded.Buckets
	s.Container = decoded.Container
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Details = decoded.Details
	s.Owner = decoded.Owner
	s.Tasks = decoded.Tasks
	s.Title = decoded.Title
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerPlan into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PlannerPlan': %+v", err)
		}
		s.CreatedBy = &impl
	}

	return nil
}
