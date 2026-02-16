package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDiagnosticStatus struct {
	// Instruction on how to mitigate a failed validation
	MitigationInstruction nullable.Type[string] `json:"mitigationInstruction,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The state of the operation
	State nullable.Type[string] `json:"state,omitempty"`

	// The validation friendly name
	ValidationName nullable.Type[string] `json:"validationName,omitempty"`
}
