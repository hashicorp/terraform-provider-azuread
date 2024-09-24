package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerAssignedToTaskBoardTaskFormat{}

type PlannerAssignedToTaskBoardTaskFormat struct {
	// Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the
	// users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
	OrderHintsByAssignee *PlannerOrderHintsByAssignee `json:"orderHintsByAssignee,omitempty"`

	// Hint value used to order the task on the AssignedTo view of the Task Board when the task isn't assigned to anyone, or
	// if the orderHintsByAssignee dictionary doesn't provide an order hint for the user the task is assigned to. The format
	// is defined as outlined here.
	UnassignedOrderHint nullable.Type[string] `json:"unassignedOrderHint,omitempty"`

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

func (s PlannerAssignedToTaskBoardTaskFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerAssignedToTaskBoardTaskFormat{}

func (s PlannerAssignedToTaskBoardTaskFormat) MarshalJSON() ([]byte, error) {
	type wrapper PlannerAssignedToTaskBoardTaskFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerAssignedToTaskBoardTaskFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerAssignedToTaskBoardTaskFormat: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerAssignedToTaskBoardTaskFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerAssignedToTaskBoardTaskFormat: %+v", err)
	}

	return encoded, nil
}
