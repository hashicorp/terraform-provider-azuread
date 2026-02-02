package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationTrainingUserCoverage struct {
	// User in an attack simulation and training campaign.
	AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// List of assigned trainings and their statuses for the user.
	UserTrainings *[]UserTrainingStatusInfo `json:"userTrainings,omitempty"`
}
