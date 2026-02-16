package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkDetails struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the Private Link policy.
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// The name of the Private Link policy in Microsoft Entra ID.
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`

	// The tenant identifier of the Microsoft Entra tenant the Private Link policy belongs to.
	PolicyTenantId nullable.Type[string] `json:"policyTenantId,omitempty"`

	// The Azure Resource Manager (ARM) path for the Private Link policy resource.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`
}
