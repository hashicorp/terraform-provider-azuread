package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInformation struct {
	CommerceUrl nullable.Type[string] `json:"commerceUrl,omitempty"`
	CompanyName nullable.Type[string] `json:"companyName,omitempty"`
	CompanyType *PartnerTenantType    `json:"companyType,omitempty"`
	HelpUrl     nullable.Type[string] `json:"helpUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PartnerTenantId   *string               `json:"partnerTenantId,omitempty"`
	SupportEmails     *[]string             `json:"supportEmails,omitempty"`
	SupportTelephones *[]string             `json:"supportTelephones,omitempty"`
	SupportUrl        nullable.Type[string] `json:"supportUrl,omitempty"`
}
