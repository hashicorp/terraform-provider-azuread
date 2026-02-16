package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerRoster{}

type PlannerRoster struct {
	// The sensitivity label applied to the roster. If mandatory labeling is enabled for the user and no label is specified,
	// the user can't create the roster. Also, if labels are mandatory for the user, the user can't change the label of the
	// roster to null. Possible values are: standard, privileged, auto, unknownFutureValue.
	AssignedSensitivityLabel *SensitivityLabelAssignment `json:"assignedSensitivityLabel,omitempty"`

	// Retrieves the members of the plannerRoster.
	Members *[]PlannerRosterMember `json:"members,omitempty"`

	// Retrieves the plans contained by the plannerRoster.
	Plans *[]PlannerPlan `json:"plans,omitempty"`

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

func (s PlannerRoster) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerRoster{}

func (s PlannerRoster) MarshalJSON() ([]byte, error) {
	type wrapper PlannerRoster
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerRoster: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerRoster: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerRoster"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerRoster: %+v", err)
	}

	return encoded, nil
}
