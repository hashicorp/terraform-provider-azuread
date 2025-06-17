package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContainer struct {
	// The identifier of the resource that contains the plan. Optional.
	ContainerId nullable.Type[string] `json:"containerId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of the resource that contains the plan. For supported types, see the previous table. Possible values are:
	// group, unknownFutureValue, roster. Use the Prefer: include-unknown-enum-members request header to get the following
	// value in this evolvable enum: roster. Optional.
	Type *PlannerContainerType `json:"type,omitempty"`

	// The full canonical URL of the container. Optional.
	Url nullable.Type[string] `json:"url,omitempty"`
}
