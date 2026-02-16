package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerDelta interface {
	Entity
	PlannerDelta() BasePlannerDeltaImpl
}

var _ PlannerDelta = BasePlannerDeltaImpl{}

type BasePlannerDeltaImpl struct {

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

func (s BasePlannerDeltaImpl) PlannerDelta() BasePlannerDeltaImpl {
	return s
}

func (s BasePlannerDeltaImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PlannerDelta = RawPlannerDeltaImpl{}

// RawPlannerDeltaImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlannerDeltaImpl struct {
	plannerDelta BasePlannerDeltaImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawPlannerDeltaImpl) PlannerDelta() BasePlannerDeltaImpl {
	return s.plannerDelta
}

func (s RawPlannerDeltaImpl) Entity() BaseEntityImpl {
	return s.plannerDelta.Entity()
}

var _ json.Marshaler = BasePlannerDeltaImpl{}

func (s BasePlannerDeltaImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePlannerDeltaImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePlannerDeltaImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePlannerDeltaImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerDelta"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePlannerDeltaImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalPlannerDeltaImplementation(input []byte) (PlannerDelta, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerDelta into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerAssignedToTaskBoardTaskFormat") {
		var out PlannerAssignedToTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerAssignedToTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucket") {
		var out PlannerBucket
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucket: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucketTaskBoardTaskFormat") {
		var out PlannerBucketTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucketTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlan") {
		var out PlannerPlan
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlan: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanDetails") {
		var out PlannerPlanDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerProgressTaskBoardTaskFormat") {
		var out PlannerProgressTaskBoardTaskFormat
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerProgressTaskBoardTaskFormat: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTask") {
		var out PlannerTask
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTask: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskDetails") {
		var out PlannerTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerUser") {
		var out PlannerUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerUser: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerDeltaImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerDeltaImpl: %+v", err)
	}

	return RawPlannerDeltaImpl{
		plannerDelta: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
