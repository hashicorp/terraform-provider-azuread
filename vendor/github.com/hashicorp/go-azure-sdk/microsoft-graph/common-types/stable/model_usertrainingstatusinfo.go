package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingStatusInfo struct {
	// Date and time of assignment of the training to the user.
	AssignedDateTime nullable.Type[string] `json:"assignedDateTime,omitempty"`

	// Date and time of completion of the training by the user.
	CompletionDateTime nullable.Type[string] `json:"completionDateTime,omitempty"`

	// Display name of the assigned training.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed,
	// overdue, unknownFutureValue.
	TrainingStatus *TrainingStatus `json:"trainingStatus,omitempty"`
}
