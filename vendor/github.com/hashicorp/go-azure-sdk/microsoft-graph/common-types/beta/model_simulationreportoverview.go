package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationReportOverview struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// List of recommended actions for a tenant to improve its security posture based on the attack simulation and training
	// campaign attack type.
	RecommendedActions *[]RecommendedAction `json:"recommendedActions,omitempty"`

	// Number of valid users in the attack simulation and training campaign.
	ResolvedTargetsCount nullable.Type[int64] `json:"resolvedTargetsCount,omitempty"`

	// Summary of simulation events in the attack simulation and training campaign.
	SimulationEventsContent *SimulationEventsContent `json:"simulationEventsContent,omitempty"`

	// Summary of assigned trainings in the attack simulation and training campaign.
	TrainingEventsContent *TrainingEventsContent `json:"trainingEventsContent,omitempty"`
}
