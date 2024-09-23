package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionInputObject struct {
	// Definition of the test object.
	Definition *ObjectDefinition `json:"definition,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Property values of the test object.
	Properties *[]StringKeyObjectValuePair `json:"properties,omitempty"`
}
