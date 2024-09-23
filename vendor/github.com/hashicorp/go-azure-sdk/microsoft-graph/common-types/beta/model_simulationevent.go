package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationEvent struct {
	// Count of the simulation event occurrence in an attack simulation and training campaign.
	Count nullable.Type[int64] `json:"count,omitempty"`

	// Name of the simulation event in an attack simulation and training campaign.
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
