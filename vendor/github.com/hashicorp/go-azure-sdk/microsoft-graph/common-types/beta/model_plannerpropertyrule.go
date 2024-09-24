package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPropertyRule interface {
	PlannerPropertyRule() BasePlannerPropertyRuleImpl
}

var _ PlannerPropertyRule = BasePlannerPropertyRuleImpl{}

type BasePlannerPropertyRuleImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Identifies which type of property rules is represented by this instance. The possible values are: taskRule,
	// bucketRule, planRule, unknownFutureValue.
	RuleKind *PlannerRuleKind `json:"ruleKind,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasePlannerPropertyRuleImpl) PlannerPropertyRule() BasePlannerPropertyRuleImpl {
	return s
}

var _ PlannerPropertyRule = RawPlannerPropertyRuleImpl{}

// RawPlannerPropertyRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlannerPropertyRuleImpl struct {
	plannerPropertyRule BasePlannerPropertyRuleImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawPlannerPropertyRuleImpl) PlannerPropertyRule() BasePlannerPropertyRuleImpl {
	return s.plannerPropertyRule
}

func UnmarshalPlannerPropertyRuleImplementation(input []byte) (PlannerPropertyRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPropertyRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerBucketPropertyRule") {
		var out PlannerBucketPropertyRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerBucketPropertyRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerPlanPropertyRule") {
		var out PlannerPlanPropertyRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerPlanPropertyRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerTaskPropertyRule") {
		var out PlannerTaskPropertyRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerTaskPropertyRule: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerPropertyRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerPropertyRuleImpl: %+v", err)
	}

	return RawPlannerPropertyRuleImpl{
		plannerPropertyRule: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
