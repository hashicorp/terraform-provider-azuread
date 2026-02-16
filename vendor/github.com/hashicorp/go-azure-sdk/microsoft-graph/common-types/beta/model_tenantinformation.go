package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantInformation struct {
	// Primary domain name of a Microsoft Entra tenant.
	DefaultDomainName nullable.Type[string] `json:"defaultDomainName,omitempty"`

	// Display name of a Microsoft Entra tenant.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Name shown to users that sign in to a Microsoft Entra tenant.
	FederationBrandName nullable.Type[string] `json:"federationBrandName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique identifier of a Microsoft Entra tenant.
	TenantId *string `json:"tenantId,omitempty"`
}
