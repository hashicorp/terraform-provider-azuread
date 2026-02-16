package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerDelta = PlannerBucket{}

type PlannerBucket struct {
	// Read-only. Nullable. Contains information about who archived or unarchived the bucket and why.
	ArchivalInfo *PlannerArchivalInfo `json:"archivalInfo,omitempty"`

	// Contains information about the origin of the bucket.
	CreationSource PlannerBucketCreation `json:"creationSource"`

	// Read-only. If set totrue, the bucket is archived. An archived bucket is read-only.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Name of the bucket.
	Name *string `json:"name,omitempty"`

	// Hint used to order items of this type in a list view. For details about the supported format, see Using order hints
	// in Planner.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

	// Plan ID to which the bucket belongs.
	PlanId nullable.Type[string] `json:"planId,omitempty"`

	// Read-only. Nullable. The collection of tasks in the bucket.
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

func (s PlannerBucket) PlannerDelta() BasePlannerDeltaImpl {
	return BasePlannerDeltaImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s PlannerBucket) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerBucket{}

func (s PlannerBucket) MarshalJSON() ([]byte, error) {
	type wrapper PlannerBucket
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerBucket: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerBucket: %+v", err)
	}

	delete(decoded, "archivalInfo")
	delete(decoded, "isArchived")
	delete(decoded, "tasks")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerBucket"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerBucket: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerBucket{}

func (s *PlannerBucket) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ArchivalInfo *PlannerArchivalInfo  `json:"archivalInfo,omitempty"`
		IsArchived   nullable.Type[bool]   `json:"isArchived,omitempty"`
		Name         *string               `json:"name,omitempty"`
		OrderHint    nullable.Type[string] `json:"orderHint,omitempty"`
		PlanId       nullable.Type[string] `json:"planId,omitempty"`
		Id           *string               `json:"id,omitempty"`
		ODataId      *string               `json:"@odata.id,omitempty"`
		ODataType    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ArchivalInfo = decoded.ArchivalInfo
	s.IsArchived = decoded.IsArchived
	s.Name = decoded.Name
	s.OrderHint = decoded.OrderHint
	s.PlanId = decoded.PlanId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerBucket into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["creationSource"]; ok {
		impl, err := UnmarshalPlannerBucketCreationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreationSource' for 'PlannerBucket': %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Tasks' for 'PlannerBucket': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Tasks = &output
	}

	return nil
}
