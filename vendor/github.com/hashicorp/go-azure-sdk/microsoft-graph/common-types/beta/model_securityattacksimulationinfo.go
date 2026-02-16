package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAttackSimulationInfo struct {
	// The date and time of the attack simulation.
	AttackSimDateTime nullable.Type[string] `json:"attackSimDateTime,omitempty"`

	// The duration (in time) for the attack simulation.
	AttackSimDurationTime nullable.Type[string] `json:"attackSimDurationTime,omitempty"`

	// The activity ID for the attack simulation.
	AttackSimId nullable.Type[string] `json:"attackSimId,omitempty"`

	// The unique identifier for the user who got the attack simulation email.
	AttackSimUserId nullable.Type[string] `json:"attackSimUserId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
