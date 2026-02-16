package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataSecurityGroupCreationOptions struct {
	// Indicates whether the security group should be created based on the org and role group.
	CreateBasedOnOrgPlusRoleGroup *bool `json:"createBasedOnOrgPlusRoleGroup,omitempty"`

	// A Boolean choice indicating whether the security group should be created based on the role group
	CreateBasedOnRoleGroup *bool `json:"createBasedOnRoleGroup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
