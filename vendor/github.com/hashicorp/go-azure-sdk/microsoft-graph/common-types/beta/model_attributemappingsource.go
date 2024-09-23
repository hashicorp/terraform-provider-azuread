package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingSource struct {
	// Equivalent expression representation of this attributeMappingSource object.
	Expression nullable.Type[string] `json:"expression,omitempty"`

	// Name parameter of the mapping source. Depending on the type property value, this can be the name of the function, the
	// name of the source attribute, or a constant value to be used.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If this object represents a function, lists function parameters. Parameters consist of attributeMappingSource objects
	// themselves, allowing for complex expressions. If type isn't Function, this property is null/empty array.
	Parameters *[]StringKeyAttributeMappingSourceValuePair `json:"parameters,omitempty"`

	Type *AttributeMappingSourceType `json:"type,omitempty"`
}
