package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationRepeatOffender struct {
	// The user in an attack simulation and training campaign.
	AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of repeat offences of the user in attack simulation and training campaigns.
	RepeatOffenceCount nullable.Type[int64] `json:"repeatOffenceCount,omitempty"`
}
