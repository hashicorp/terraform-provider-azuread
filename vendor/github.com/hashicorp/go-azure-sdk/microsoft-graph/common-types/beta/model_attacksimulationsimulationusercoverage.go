package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationSimulationUserCoverage struct {
	// User in an attack simulation and training campaign.
	AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`

	// Number of link clicks in the received payloads by the user in attack simulation and training campaigns.
	ClickCount nullable.Type[int64] `json:"clickCount,omitempty"`

	// Number of compromising actions by the user in attack simulation and training campaigns.
	CompromisedCount nullable.Type[int64] `json:"compromisedCount,omitempty"`

	// Date and time of the latest attack simulation and training campaign that the user was included in.
	LatestSimulationDateTime nullable.Type[string] `json:"latestSimulationDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of attack simulation and training campaigns that the user was included in.
	SimulationCount nullable.Type[int64] `json:"simulationCount,omitempty"`
}
