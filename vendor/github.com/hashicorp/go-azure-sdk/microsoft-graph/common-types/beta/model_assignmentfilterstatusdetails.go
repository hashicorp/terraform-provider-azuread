package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterStatusDetails struct {
	// Device properties used for filter evaluation during device check-in time.
	DeviceProperties *[]KeyValuePair `json:"deviceProperties,omitempty"`

	// Evaluation result summaries for each filter associated to device and payload
	EvalutionSummaries *[]AssignmentFilterEvaluationSummary `json:"evalutionSummaries,omitempty"`

	// Unique identifier for the device object.
	ManagedDeviceId nullable.Type[string] `json:"managedDeviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier for payload object.
	PayloadId nullable.Type[string] `json:"payloadId,omitempty"`

	// Unique identifier for UserId object. Can be null
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
