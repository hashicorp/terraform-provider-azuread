package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinition struct {
	// true if the attribute should be used as the anchor for the object. Anchor attributes must have a unique value
	// identifying an object, and must be immutable. Default is false. One, and only one, of the object's attributes must be
	// designated as the anchor to support synchronization.
	Anchor *bool `json:"anchor,omitempty"`

	ApiExpressions *[]StringKeyStringValuePair `json:"apiExpressions,omitempty"`

	// true if value of this attribute should be treated as case-sensitive. This setting affects how the synchronization
	// engine detects changes for the attribute.
	CaseExact *bool `json:"caseExact,omitempty"`

	// The default value of the attribute.
	DefaultValue nullable.Type[string] `json:"defaultValue,omitempty"`

	// 'true' to allow null values for attributes.
	FlowNullValues *bool `json:"flowNullValues,omitempty"`

	// Metadata for the given object.
	Metadata *[]AttributeDefinitionMetadataEntry `json:"metadata,omitempty"`

	// true if an attribute can have multiple values. Default is false.
	Multivalued *bool `json:"multivalued,omitempty"`

	Mutability *Mutability `json:"mutability,omitempty"`

	// Name of the attribute. Must be unique within the object definition. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// For attributes with reference type, lists referenced objects (for example, the manager attribute would list User as
	// the referenced object).
	ReferencedObjects *[]ReferencedObject `json:"referencedObjects,omitempty"`

	// true if attribute is required. Object can not be created if any of the required attributes are missing. If during
	// synchronization, the required attribute has no value, the default value will be used. If default the value was not
	// set, synchronization will record an error.
	Required *bool `json:"required,omitempty"`

	Type *AttributeType `json:"type,omitempty"`
}
