package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTemplateAction struct {
	Description nullable.Type[string] `json:"description,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	Licenses    *LicenseDetails       `json:"licenses,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Service          nullable.Type[string]    `json:"service,omitempty"`
	Settings         *[]ManagedTenantsSetting `json:"settings,omitempty"`
	TemplateActionId *string                  `json:"templateActionId,omitempty"`
}
