package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetonationBehaviourDetails struct {
	// The status of the action performed during detonation (e.g., 'Successful', 'Failed', 'Blocked').
	ActionStatus nullable.Type[string] `json:"actionStatus,omitempty"`

	// Categorizes the capability or type of behavior observed.
	BehaviourCapability nullable.Type[string] `json:"behaviourCapability,omitempty"`

	// Groups related behaviors together for classification purposes.
	BehaviourGroup nullable.Type[string] `json:"behaviourGroup,omitempty"`

	// More contextual information about the observed behavior or action.
	Details nullable.Type[string] `json:"details,omitempty"`

	// The date and time when the behavior or action was observed during detonation.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The specific operation or action that was performed.
	Operation nullable.Type[string] `json:"operation,omitempty"`

	// The unique identifier of the process involved in the behavior.
	ProcessId nullable.Type[string] `json:"processId,omitempty"`

	// The name of the process that performed or was involved in the behavior.
	ProcessName nullable.Type[string] `json:"processName,omitempty"`

	// The target of the operation.
	Target nullable.Type[string] `json:"target,omitempty"`
}
