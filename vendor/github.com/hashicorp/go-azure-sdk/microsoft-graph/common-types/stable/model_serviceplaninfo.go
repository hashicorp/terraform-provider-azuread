package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePlanInfo struct {
	// The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to
	// individual users.Company - service plan can be assigned to the entire tenant.
	AppliesTo nullable.Type[string] `json:"appliesTo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled
	// - Service is disabled.Error - The service plan isn't provisioned and is in an error state.PendingInput - The service
	// isn't provisioned and is awaiting service confirmation.PendingActivation - The service is provisioned but requires
	// explicit activation by an administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has
	// added a new service to the product SKU and it isn't activated in the tenant.
	ProvisioningStatus nullable.Type[string] `json:"provisioningStatus,omitempty"`

	// The unique identifier of the service plan.
	ServicePlanId nullable.Type[string] `json:"servicePlanId,omitempty"`

	// The name of the service plan.
	ServicePlanName nullable.Type[string] `json:"servicePlanName,omitempty"`
}
