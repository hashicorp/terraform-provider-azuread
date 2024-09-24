package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedPlan struct {
	// For example, 'Enabled'.
	CapabilityStatus nullable.Type[string] `json:"capabilityStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For example, 'Success'.
	ProvisioningStatus nullable.Type[string] `json:"provisioningStatus,omitempty"`

	// The name of the service; for example, 'AccessControlS2S'
	Service nullable.Type[string] `json:"service,omitempty"`
}
