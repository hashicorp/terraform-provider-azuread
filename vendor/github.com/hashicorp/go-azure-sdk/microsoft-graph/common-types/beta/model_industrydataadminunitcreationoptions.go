package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAdminUnitCreationOptions struct {
	// Indicates whether the administrative unit should be created based on the org.
	CreateBasedOnOrg *bool `json:"createBasedOnOrg,omitempty"`

	// Indicates whether the administrative unit should be created based on the org and role group.
	CreateBasedOnOrgPlusRoleGroup *bool `json:"createBasedOnOrgPlusRoleGroup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
