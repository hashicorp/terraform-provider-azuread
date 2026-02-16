package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingParameterSchema struct {
	// The given parameter can be provided multiple times (for example, multiple input strings in the
	// Concatenate(string,string,...) function).
	AllowMultipleOccurrences *bool `json:"allowMultipleOccurrences,omitempty"`

	// Parameter name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// true if the parameter is required; otherwise false.
	Required *bool `json:"required,omitempty"`

	Type *AttributeType `json:"type,omitempty"`
}
