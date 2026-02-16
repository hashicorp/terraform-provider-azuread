package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterSupportedProperty struct {
	// The data type of the property.
	DataType nullable.Type[string] `json:"dataType,omitempty"`

	// Indicates whether the property is a collection type or not.
	IsCollection *bool `json:"isCollection,omitempty"`

	// Name of the property.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Regex string to do validation on the property value.
	PropertyRegexConstraint nullable.Type[string] `json:"propertyRegexConstraint,omitempty"`

	// List of all supported operators on this property.
	SupportedOperators *[]AssignmentFilterOperator `json:"supportedOperators,omitempty"`

	// List of all supported values for this property, empty if everything is supported.
	SupportedValues *[]string `json:"supportedValues,omitempty"`
}
