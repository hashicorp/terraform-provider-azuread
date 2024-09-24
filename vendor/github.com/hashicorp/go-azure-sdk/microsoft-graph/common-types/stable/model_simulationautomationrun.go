package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SimulationAutomationRun{}

type SimulationAutomationRun struct {
	// Date and time when the run ends in an attack simulation automation.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Unique identifier for the attack simulation campaign initiated in the attack simulation automation run.
	SimulationId nullable.Type[string] `json:"simulationId,omitempty"`

	// Date and time when the run starts in an attack simulation automation.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Status of the attack simulation automation run. The possible values are: unknown, running, succeeded, failed,
	// skipped, unknownFutureValue.
	Status *SimulationAutomationRunStatus `json:"status,omitempty"`

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

func (s SimulationAutomationRun) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SimulationAutomationRun{}

func (s SimulationAutomationRun) MarshalJSON() ([]byte, error) {
	type wrapper SimulationAutomationRun
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SimulationAutomationRun: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SimulationAutomationRun: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.simulationAutomationRun"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SimulationAutomationRun: %+v", err)
	}

	return encoded, nil
}
