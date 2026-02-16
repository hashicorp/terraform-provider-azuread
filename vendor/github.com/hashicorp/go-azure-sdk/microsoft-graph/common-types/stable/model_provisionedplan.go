package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedPlan struct {
	// Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See
	// a detailed description of each value.
	CapabilityStatus nullable.Type[string] `json:"capabilityStatus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The possible values are:Success - Service is fully provisioned.Disabled - Service is disabled.Error - The service
	// plan isn't provisioned and is in an error state.PendingInput - The service isn't provisioned and is awaiting service
	// confirmation.PendingActivation - The service is provisioned but requires explicit activation by an administrator (for
	// example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it
	// isn't activated in the tenant.
	ProvisioningStatus nullable.Type[string] `json:"provisioningStatus,omitempty"`

	// The name of the service; for example, 'AccessControlS2S'.
	Service nullable.Type[string] `json:"service,omitempty"`
}
