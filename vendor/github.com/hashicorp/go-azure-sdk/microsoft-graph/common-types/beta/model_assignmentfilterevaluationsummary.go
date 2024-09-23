package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummary struct {
	// The admin defined name for assignment filter.
	AssignmentFilterDisplayName nullable.Type[string] `json:"assignmentFilterDisplayName,omitempty"`

	// Unique identifier for the assignment filter object
	AssignmentFilterId nullable.Type[string] `json:"assignmentFilterId,omitempty"`

	// The time the assignment filter was last modified.
	AssignmentFilterLastModifiedDateTime *string `json:"assignmentFilterLastModifiedDateTime,omitempty"`

	// Supported platform types.
	AssignmentFilterPlatform *DevicePlatformType `json:"assignmentFilterPlatform,omitempty"`

	// Represents type of the assignment filter.
	AssignmentFilterType *DeviceAndAppManagementAssignmentFilterType `json:"assignmentFilterType,omitempty"`

	// A collection of filter types and their corresponding evaluation results.
	AssignmentFilterTypeAndEvaluationResults *[]AssignmentFilterTypeAndEvaluationResult `json:"assignmentFilterTypeAndEvaluationResults,omitempty"`

	// The time assignment filter was evaluated.
	EvaluationDateTime *string `json:"evaluationDateTime,omitempty"`

	// Supported evaluation results for filter.
	EvaluationResult *AssignmentFilterEvaluationResult `json:"evaluationResult,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
