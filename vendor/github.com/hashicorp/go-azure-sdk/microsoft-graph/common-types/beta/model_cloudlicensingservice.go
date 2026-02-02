package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLicensingService struct {
	AssignableTo *CloudLicensingAssigneeTypes `json:"assignableTo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of the service plan that is equal to the servicePlanId property on the related servicePlanInfo
	// objects.
	PlanId *string `json:"planId,omitempty"`

	// The name of the service plan that is equal to the servicePlanName property on the related servicePlanInfo objects.
	PlanName nullable.Type[string] `json:"planName,omitempty"`
}
