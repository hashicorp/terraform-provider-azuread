package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTroubleshootingErrorDetails struct {
	Context nullable.Type[string] `json:"context,omitempty"`
	Failure nullable.Type[string] `json:"failure,omitempty"`

	// The detailed description of what went wrong.
	FailureDetails nullable.Type[string] `json:"failureDetails,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The detailed description of how to remediate this issue.
	Remediation nullable.Type[string] `json:"remediation,omitempty"`

	// Links to helpful documentation about this failure.
	Resources *[]DeviceManagementTroubleshootingErrorResource `json:"resources,omitempty"`
}
