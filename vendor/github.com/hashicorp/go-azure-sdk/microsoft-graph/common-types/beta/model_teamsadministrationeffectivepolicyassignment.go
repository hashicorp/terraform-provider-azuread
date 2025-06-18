package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAdministrationEffectivePolicyAssignment struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PolicyAssignment *TeamsAdministrationPolicyAssignment `json:"policyAssignment,omitempty"`

	// The type of the assigned policy; for example, TeamsMeetingPolicy and TeamsCallingPolicy.
	PolicyType *string `json:"policyType,omitempty"`
}
