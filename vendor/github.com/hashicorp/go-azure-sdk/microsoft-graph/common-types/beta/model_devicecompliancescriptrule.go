package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRule struct {
	// Data types for rules.
	DataType *DataType `json:"dataType,omitempty"`

	// Operator for rules.
	DeviceComplianceScriptRulOperator *DeviceComplianceScriptRulOperator `json:"deviceComplianceScriptRulOperator,omitempty"`

	// Data types for rules.
	DeviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDataType `json:"deviceComplianceScriptRuleDataType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Operand specified in the rule.
	Operand nullable.Type[string] `json:"operand,omitempty"`

	// Operator for rules.
	Operator *Operator `json:"operator,omitempty"`

	// Setting name specified in the rule.
	SettingName nullable.Type[string] `json:"settingName,omitempty"`
}
