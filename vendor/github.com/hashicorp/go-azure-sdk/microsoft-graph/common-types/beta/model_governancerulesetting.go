package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRuleSetting struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The id of the rule. For example, ExpirationRule and MfaRule.
	RuleIdentifier nullable.Type[string] `json:"ruleIdentifier,omitempty"`

	// The settings of the rule. The value is a JSON string with a list of pairs in the format of
	// ParameterName:ParameterValue. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
	Setting nullable.Type[string] `json:"setting,omitempty"`
}
