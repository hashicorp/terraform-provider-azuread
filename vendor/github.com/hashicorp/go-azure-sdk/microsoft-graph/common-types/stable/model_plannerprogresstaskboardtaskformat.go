package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerProgressTaskBoardTaskFormat{}

type PlannerProgressTaskBoardTaskFormat struct {
	// Hint value used to order the task on the progress view of the task board. For details about the supported format, see
	// Using order hints in Planner.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

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

func (s PlannerProgressTaskBoardTaskFormat) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerProgressTaskBoardTaskFormat{}

func (s PlannerProgressTaskBoardTaskFormat) MarshalJSON() ([]byte, error) {
	type wrapper PlannerProgressTaskBoardTaskFormat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerProgressTaskBoardTaskFormat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerProgressTaskBoardTaskFormat: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerProgressTaskBoardTaskFormat"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerProgressTaskBoardTaskFormat: %+v", err)
	}

	return encoded, nil
}
