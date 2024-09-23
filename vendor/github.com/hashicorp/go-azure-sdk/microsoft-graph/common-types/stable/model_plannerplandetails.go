package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerPlanDetails{}

type PlannerPlanDetails struct {
	// An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan.
	CategoryDescriptions *PlannerCategoryDescriptions `json:"categoryDescriptions,omitempty"`

	// Set of user IDs that this plan is shared with. If you're using Microsoft 365 groups, use the Groups API to manage
	// group membership to share the group's plan. You can also add existing members of the group to this collection,
	// although it isn't required for them to access the plan owned by the group.
	SharedWith *PlannerUserIds `json:"sharedWith,omitempty"`

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

func (s PlannerPlanDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerPlanDetails{}

func (s PlannerPlanDetails) MarshalJSON() ([]byte, error) {
	type wrapper PlannerPlanDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerPlanDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerPlanDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerPlanDetails: %+v", err)
	}

	return encoded, nil
}
