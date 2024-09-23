package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSimulationDetails struct {
	// Number of trainings assigned to a user in an attack simulation and training campaign.
	AssignedTrainingsCount nullable.Type[int64] `json:"assignedTrainingsCount,omitempty"`

	// Number of trainings completed by a user in an attack simulation and training campaign.
	CompletedTrainingsCount nullable.Type[int64] `json:"completedTrainingsCount,omitempty"`

	// Date and time of the compromising online action by a user in an attack simulation and training campaign.
	CompromisedDateTime nullable.Type[string] `json:"compromisedDateTime,omitempty"`

	// Number of trainings in progress by a user in an attack simulation and training campaign.
	InProgressTrainingsCount nullable.Type[int64] `json:"inProgressTrainingsCount,omitempty"`

	// Indicates whether a user was compromised in an attack simulation and training campaign.
	IsCompromised nullable.Type[bool] `json:"isCompromised,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Date and time when a user reported the delivered payload as phishing in the attack simulation and training campaign.
	ReportedPhishDateTime nullable.Type[string] `json:"reportedPhishDateTime,omitempty"`

	// List of simulation events of a user in the attack simulation and training campaign.
	SimulationEvents *[]UserSimulationEventInfo `json:"simulationEvents,omitempty"`

	// User in an attack simulation and training campaign.
	SimulationUser *AttackSimulationUser `json:"simulationUser,omitempty"`

	// List of training events of a user in the attack simulation and training campaign.
	TrainingEvents *[]UserTrainingEventInfo `json:"trainingEvents,omitempty"`
}
