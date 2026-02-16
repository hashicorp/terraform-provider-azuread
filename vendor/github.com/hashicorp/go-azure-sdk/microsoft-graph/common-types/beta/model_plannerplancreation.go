package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanCreation interface {
	PlannerPlanCreation() BasePlannerPlanCreationImpl
}

var _ PlannerPlanCreation = BasePlannerPlanCreationImpl{}

type BasePlannerPlanCreationImpl struct {
	// Specifies what kind of creation source the plan is created with. The possible values are: external, publication and
	// unknownFutureValue.
	CreationSourceKind *PlannerCreationSourceKind `json:"creationSourceKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerPlanCreationImpl) PlannerPlanCreation() BasePlannerPlanCreationImpl {
	return s
}

var _ PlannerPlanCreation = RawPlannerPlanCreationImpl{}

// RawPlannerPlanCreationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlannerPlanCreationImpl struct {
	plannerPlanCreation BasePlannerPlanCreationImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawPlannerPlanCreationImpl) PlannerPlanCreation() BasePlannerPlanCreationImpl {
	return s.plannerPlanCreation
}

func UnmarshalPlannerPlanCreationImplementation(input []byte) (PlannerPlanCreation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanCreation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerExternalPlanSource") {
		var out PlannerExternalPlanSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerExternalPlanSource: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerPlanCreationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerPlanCreationImpl: %+v", err)
	}

	return RawPlannerPlanCreationImpl{
		plannerPlanCreation: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
