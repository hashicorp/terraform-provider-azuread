package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingEventInfo struct {
	// Display name of the training.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed,
	// overdue, unknownFutureValue.
	LatestTrainingStatus *TrainingStatus `json:"latestTrainingStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Event details of the training when it was assigned to the user.
	TrainingAssignedProperties *UserTrainingContentEventInfo `json:"trainingAssignedProperties,omitempty"`

	// Event details of the training when it was completed by the user.
	TrainingCompletedProperties *UserTrainingContentEventInfo `json:"trainingCompletedProperties,omitempty"`

	// Event details of the training when it was updated/in-progress by the user.
	TrainingUpdatedProperties *UserTrainingContentEventInfo `json:"trainingUpdatedProperties,omitempty"`
}
