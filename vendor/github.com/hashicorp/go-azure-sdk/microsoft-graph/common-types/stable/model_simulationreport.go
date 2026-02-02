package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationReport struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Overview of an attack simulation and training campaign.
	Overview *SimulationReportOverview `json:"overview,omitempty"`

	// The tenant users and their online actions in an attack simulation and training campaign.
	SimulationUsers *[]UserSimulationDetails `json:"simulationUsers,omitempty"`
}
