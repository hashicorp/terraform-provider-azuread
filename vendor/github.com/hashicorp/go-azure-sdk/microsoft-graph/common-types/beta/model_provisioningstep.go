package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStep struct {
	// Summary of what occurred during the step.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Details of what occurred during the step.
	Details *DetailsInfo `json:"details,omitempty"`

	// Name of the step.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of step. Possible values are: import, scoping, matching, processing, referenceResolution, export,
	// unknownFutureValue.
	ProvisioningStepType *ProvisioningStepType `json:"provisioningStepType,omitempty"`

	// Status of the step. Possible values are: success, warning, failure, skipped, unknownFutureValue.
	Status *ProvisioningResult `json:"status,omitempty"`
}
