package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainer interface {
	PlannerPlanContainer() BasePlannerPlanContainerImpl
}

var _ PlannerPlanContainer = BasePlannerPlanContainerImpl{}

type BasePlannerPlanContainerImpl struct {
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

func (s BasePlannerPlanContainerImpl) PlannerPlanContainer() BasePlannerPlanContainerImpl {
	return s
}

var _ PlannerPlanContainer = RawPlannerPlanContainerImpl{}

// RawPlannerPlanContainerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlannerPlanContainerImpl struct {
	plannerPlanContainer BasePlannerPlanContainerImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawPlannerPlanContainerImpl) PlannerPlanContainer() BasePlannerPlanContainerImpl {
	return s.plannerPlanContainer
}

func UnmarshalPlannerPlanContainerImplementation(input []byte) (PlannerPlanContainer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerPlanContainer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerSharedWithContainer") {
		var out PlannerSharedWithContainer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerSharedWithContainer: %+v", err)
		}
		return out, nil
	}

	var parent BasePlannerPlanContainerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlannerPlanContainerImpl: %+v", err)
	}

	return RawPlannerPlanContainerImpl{
		plannerPlanContainer: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
