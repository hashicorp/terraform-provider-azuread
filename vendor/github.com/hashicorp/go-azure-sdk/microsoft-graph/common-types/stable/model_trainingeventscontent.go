package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingEventsContent struct {
	// List of assigned trainings and their information in an attack simulation and training campaign.
	AssignedTrainingsInfos *[]AssignedTrainingInfo `json:"assignedTrainingsInfos,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of users who were assigned trainings in an attack simulation and training campaign.
	TrainingsAssignedUserCount nullable.Type[int64] `json:"trainingsAssignedUserCount,omitempty"`
}
