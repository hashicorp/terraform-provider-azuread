package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedTrainingInfo struct {
	// Number of users who were assigned the training in an attack simulation and training campaign.
	AssignedUserCount nullable.Type[int64] `json:"assignedUserCount,omitempty"`

	// Number of users who completed the training in an attack simulation and training campaign.
	CompletedUserCount nullable.Type[int64] `json:"completedUserCount,omitempty"`

	// Display name of the training in an attack simulation and training campaign.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
