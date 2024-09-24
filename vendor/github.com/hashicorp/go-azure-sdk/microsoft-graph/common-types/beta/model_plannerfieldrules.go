package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerFieldRules struct {
	// The default rules that apply if no override matches to the current data.
	DefaultRules *[]string `json:"defaultRules,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Overrides that specify different rules for specific data associated with the field.
	Overrides *[]PlannerRuleOverride `json:"overrides,omitempty"`
}
