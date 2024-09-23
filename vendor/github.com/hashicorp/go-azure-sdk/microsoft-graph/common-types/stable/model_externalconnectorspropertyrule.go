package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRule struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Operation *ExternalConnectorsRuleOperation `json:"operation,omitempty"`

	// The property from the externalItem schema. Required.
	Property string `json:"property"`

	// A collection with one or many strings. One or more specified strings are matched with the specified property using
	// the specified operation. Required.
	Values []string `json:"values"`

	ValuesJoinedBy *BinaryOperator `json:"valuesJoinedBy,omitempty"`
}
