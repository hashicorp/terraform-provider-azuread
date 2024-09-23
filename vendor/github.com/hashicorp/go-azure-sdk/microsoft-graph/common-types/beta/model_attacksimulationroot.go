package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AttackSimulationRoot{}

type AttackSimulationRoot struct {
	// Represents an end user's notification for an attack simulation training.
	EndUserNotifications *[]EndUserNotification `json:"endUserNotifications,omitempty"`

	// Represents an attack simulation training landing page.
	LandingPages *[]LandingPage `json:"landingPages,omitempty"`

	// Represents an attack simulation training login page.
	LoginPages *[]LoginPage `json:"loginPages,omitempty"`

	// Represents an attack simulation training operation.
	Operations *[]AttackSimulationOperation `json:"operations,omitempty"`

	// Represents an attack simulation training campaign payload in a tenant.
	Payloads *[]Payload `json:"payloads,omitempty"`

	// Represents simulation automation created to run on a tenant.
	SimulationAutomations *[]SimulationAutomation `json:"simulationAutomations,omitempty"`

	// Represents an attack simulation training campaign in a tenant.
	Simulations *[]Simulation `json:"simulations,omitempty"`

	// Represents a training campaign in a tenant.
	TrainingCampaigns *[]TrainingCampaign `json:"trainingCampaigns,omitempty"`

	// Represents details about attack simulation trainings.
	Trainings *[]Training `json:"trainings,omitempty"`

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

func (s AttackSimulationRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttackSimulationRoot{}

func (s AttackSimulationRoot) MarshalJSON() ([]byte, error) {
	type wrapper AttackSimulationRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttackSimulationRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttackSimulationRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attackSimulationRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttackSimulationRoot: %+v", err)
	}

	return encoded, nil
}
