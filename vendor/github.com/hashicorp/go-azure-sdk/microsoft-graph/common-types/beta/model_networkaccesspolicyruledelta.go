package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyRuleDelta struct {
	Action *NetworkaccessForwardingRuleAction `json:"action,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The identifier of the policy rule to update.
	RuleId *string `json:"ruleId,omitempty"`
}
