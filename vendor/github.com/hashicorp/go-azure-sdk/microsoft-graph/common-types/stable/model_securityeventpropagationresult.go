package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventPropagationResult struct {
	// The name of the specific location in the workload associated with the event.
	Location nullable.Type[string] `json:"location,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the workload associated with the event.
	ServiceName nullable.Type[string] `json:"serviceName,omitempty"`

	// Indicates the status of the event creation request. The possible values are: none, inProcessing, failed, success,
	// unknownFutureValue.
	Status *SecurityEventPropagationStatus `json:"status,omitempty"`

	// Additional information about the status of the event creation request.
	StatusInformation nullable.Type[string] `json:"statusInformation,omitempty"`
}
