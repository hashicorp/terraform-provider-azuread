package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerPropertyRule = PlannerBucketPropertyRule{}

type PlannerBucketPropertyRule struct {
	Order *[]string `json:"order,omitempty"`
	Title *[]string `json:"title,omitempty"`

	// Fields inherited from PlannerPropertyRule

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

func (s PlannerBucketPropertyRule) PlannerPropertyRule() BasePlannerPropertyRuleImpl {
	return BasePlannerPropertyRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleKind:  s.RuleKind,
	}
}

var _ json.Marshaler = PlannerBucketPropertyRule{}

func (s PlannerBucketPropertyRule) MarshalJSON() ([]byte, error) {
	type wrapper PlannerBucketPropertyRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerBucketPropertyRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerBucketPropertyRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerBucketPropertyRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerBucketPropertyRule: %+v", err)
	}

	return encoded, nil
}
