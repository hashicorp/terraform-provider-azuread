package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluateDynamicMembershipResult struct {
	// If a group ID is provided, the value is the membership rule for the group. If a group ID isn't provided, the value is
	// the membership rule that was provided as a parameter. For more information, see Dynamic membership rules for groups
	// in Microsoft Entra ID.
	MembershipRule nullable.Type[string] `json:"membershipRule,omitempty"`

	// Provides a detailed analysis of the membership evaluation result.
	MembershipRuleEvaluationDetails *ExpressionEvaluationDetails `json:"membershipRuleEvaluationDetails,omitempty"`

	// The value is true if the user or device is a member of the group. The value can also be true if a membership rule was
	// provided and the user or device passes the rule evaluation; otherwise false.
	MembershipRuleEvaluationResult *bool `json:"membershipRuleEvaluationResult,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
