package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerPlanContainer = PlannerSharedWithContainer{}

type PlannerSharedWithContainer struct {
	AccessLevel *PlannerPlanAccessLevel `json:"accessLevel,omitempty"`

	// Fields inherited from PlannerPlanContainer

	// The identifier of the resource that contains the plan. Optional.
	ContainerId nullable.Type[string] `json:"containerId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of the resource that contains the plan. For supported types, see the previous table. Possible values are:
	// group, unknownFutureValue, roster, project, driveItem, user, teamsChannel, and plannerTask. Use the Prefer:
	// include-unknown-enum-members request header to get the following values in this evolvable enum: roster, project,
	// driveItem, user, teamsChannel, and plannerTask. Optional.
	Type *PlannerContainerType `json:"type,omitempty"`

	// The full canonical URL of the container. Optional.
	Url nullable.Type[string] `json:"url,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PlannerSharedWithContainer) PlannerPlanContainer() BasePlannerPlanContainerImpl {
	return BasePlannerPlanContainerImpl{
		ContainerId: s.ContainerId,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Type:        s.Type,
		Url:         s.Url,
	}
}

var _ json.Marshaler = PlannerSharedWithContainer{}

func (s PlannerSharedWithContainer) MarshalJSON() ([]byte, error) {
	type wrapper PlannerSharedWithContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerSharedWithContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerSharedWithContainer: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerSharedWithContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerSharedWithContainer: %+v", err)
	}

	return encoded, nil
}
