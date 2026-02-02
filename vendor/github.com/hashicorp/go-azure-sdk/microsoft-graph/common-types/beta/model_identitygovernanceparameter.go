package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceParameter struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ValueType *IdentityGovernanceValueType `json:"valueType,omitempty"`

	// The values of the parameter.
	Values *[]string `json:"values,omitempty"`
}
