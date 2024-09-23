package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipCustomerParticipant struct {
	// The display name of the customer tenant as set by Microsoft Entra ID. Read-only
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Microsoft Entra ID-assigned tenant ID of the customer tenant.
	TenantId *string `json:"tenantId,omitempty"`
}
