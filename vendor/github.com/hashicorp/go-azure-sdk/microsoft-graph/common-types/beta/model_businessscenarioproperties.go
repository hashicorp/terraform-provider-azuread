package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioProperties struct {
	// The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be
	// placed in the corresponding plannerBucket in the target plan. Required.
	ExternalBucketId nullable.Type[string] `json:"externalBucketId,omitempty"`

	// The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by
	// their externalContextId. Optional.
	ExternalContextId nullable.Type[string] `json:"externalContextId,omitempty"`

	// Application-specific identifier for the task. Every task for the same scenario must have a unique identifier
	// specified for this property. Required.
	ExternalObjectId nullable.Type[string] `json:"externalObjectId,omitempty"`

	// Application-specific version of the task. Optional.
	ExternalObjectVersion nullable.Type[string] `json:"externalObjectVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The URL to the application-specific experience for this task. Optional.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
