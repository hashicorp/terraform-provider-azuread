package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerPropertyRule = PlannerTaskPropertyRule{}

type PlannerTaskPropertyRule struct {
	// Rules and restrictions for applied categories. This value doesn't currently support overrides. Accepted values for
	// the default rule and individual overrides are allow, block.
	AppliedCategories *PlannerFieldRules `json:"appliedCategories,omitempty"`

	// Rules and restrictions for approval. Allowed overrides are userCreated and applicationCreated. Accepted values for
	// the default rule and individual overrides are: allow, add, remove, block.
	ApprovalAttachment *PlannerFieldRules `json:"approvalAttachment,omitempty"`

	// Rules and restrictions for assignments. Allowed overrides are userCreated and applicationCreated. Accepted values for
	// the default rule and individual overrides are allow, add, addSelf, addOther, remove, removeSelf, removeOther, block.
	Assignments *PlannerFieldRules `json:"assignments,omitempty"`

	// Rules and restrictions for checklist. Allowed overrides are userCreated and applicationCreated. Accepted values for
	// the default rule and individual overrides are allow, add, remove, update, check, reorder, block.
	CheckLists *PlannerFieldRules `json:"checkLists,omitempty"`

	// Rules and restrictions for completion requirements of the task. Accepted values are allow, add, remove, edit, and
	// block.
	CompletionRequirements *[]string `json:"completionRequirements,omitempty"`

	// Rules and restrictions for deleting the task. Accepted values are allow and block.
	Delete *[]string `json:"delete,omitempty"`

	// Rules and restrictions for changing the due date of the task. Accepted values are allow and block.
	DueDate *[]string `json:"dueDate,omitempty"`

	// Rules and restrictions for forms. Allowed overrides are userCreated and applicationCreated. The following are the
	// accepted values for the default rule and individual overrides: allow, add, addResponse, remove, update, block.
	Forms *PlannerFieldRules `json:"forms,omitempty"`

	// Rules and restrictions for moving the task between buckets or plans. Accepted values are allow, moveBetweenPlans,
	// moveBetweenBuckets, and block.
	Move *[]string `json:"move,omitempty"`

	// Rules and restrictions for changing the notes of the task. Accepted values are allow and block.
	Notes *[]string `json:"notes,omitempty"`

	// Rules and restrictions for changing the order of the task. Accepted values are allow and block.
	Order *[]string `json:"order,omitempty"`

	// Rules and restrictions for changing the completion percentage of the task. Accepted values are allow, setToComplete,
	// overrideRequirements, setToNotStarted, setToInProgress, and block.
	PercentComplete *[]string `json:"percentComplete,omitempty"`

	// Rules and restrictions for changing the preview type of the task. Accepted values are allow and block.
	PreviewType *[]string `json:"previewType,omitempty"`

	// Rules and restrictions for changing the priority of the task. Accepted values are allow and block.
	Priority *[]string `json:"priority,omitempty"`

	// Rules and restrictions for references. Allowed overrides are userCreated and applicationCreated. Accepted values for
	// the default rule and individual overrides are allow, add, remove, block.
	References *PlannerFieldRules `json:"references,omitempty"`

	// Rules and restrictions for changing the start date of the task. Accepted values are allow and block.
	StartDate *[]string `json:"startDate,omitempty"`

	// Rules and restrictions for changing the title of the task. Accepted values are allow and block.
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

func (s PlannerTaskPropertyRule) PlannerPropertyRule() BasePlannerPropertyRuleImpl {
	return BasePlannerPropertyRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleKind:  s.RuleKind,
	}
}

var _ json.Marshaler = PlannerTaskPropertyRule{}

func (s PlannerTaskPropertyRule) MarshalJSON() ([]byte, error) {
	type wrapper PlannerTaskPropertyRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerTaskPropertyRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTaskPropertyRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerTaskPropertyRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerTaskPropertyRule: %+v", err)
	}

	return encoded, nil
}
