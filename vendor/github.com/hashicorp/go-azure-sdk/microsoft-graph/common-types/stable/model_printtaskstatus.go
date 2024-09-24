package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskStatus struct {
	// A human-readable description of the current processing state of the printTask.
	Description *string `json:"description,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *PrintTaskProcessingState `json:"state,omitempty"`
}
