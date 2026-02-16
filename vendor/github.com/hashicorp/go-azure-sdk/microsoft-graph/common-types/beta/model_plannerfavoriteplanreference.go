package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFavoritePlanReference struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Hint used to order items of this type in a list view. The format is defined in Using order hints in Planner.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

	// Title of the plan at the time the user marked it as a favorite.
	PlanTitle nullable.Type[string] `json:"planTitle,omitempty"`
}
