package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Planner{}

type Planner struct {
	// Read-only. Nullable. Returns a collection of the specified buckets
	Buckets *[]PlannerBucket `json:"buckets,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified plans
	Plans *[]PlannerPlan `json:"plans,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified tasks
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

func (s Planner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Planner{}

func (s Planner) MarshalJSON() ([]byte, error) {
	type wrapper Planner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Planner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Planner: %+v", err)
	}

	delete(decoded, "buckets")
	delete(decoded, "plans")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.planner"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Planner: %+v", err)
	}

	return encoded, nil
}
