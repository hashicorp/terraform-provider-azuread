package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingCompletionSummary struct {
	// The number of users who completed all the trainings before the due date.
	CompletedUsersCount nullable.Type[int64] `json:"completedUsersCount,omitempty"`

	// The number of users who started at least one training.
	InProgressUsersCount nullable.Type[int64] `json:"inProgressUsersCount,omitempty"`

	// The number of users who didn't complete all the trainings before the due date.
	NotCompletedUsersCount nullable.Type[int64] `json:"notCompletedUsersCount,omitempty"`

	// The number of users who didn't start any training.
	NotStartedUsersCount nullable.Type[int64] `json:"notStartedUsersCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of users who are already assigned the same training.
	PreviouslyAssignedUsersCount nullable.Type[int64] `json:"previouslyAssignedUsersCount,omitempty"`
}
