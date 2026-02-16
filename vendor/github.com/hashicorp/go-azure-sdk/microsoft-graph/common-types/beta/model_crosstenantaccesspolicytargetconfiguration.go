package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetConfiguration struct {
	// Defines whether access is allowed or blocked. The possible values are: allowed, blocked, unknownFutureValue.
	AccessType *CrossTenantAccessPolicyTargetConfigurationAccessType `json:"accessType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies whether to target users, groups, or applications with this rule.
	Targets *[]CrossTenantAccessPolicyTarget `json:"targets,omitempty"`
}
