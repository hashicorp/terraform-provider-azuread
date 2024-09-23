package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StringKeyAttributeMappingSourceValuePair struct {
	// The name of the parameter.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The value of the parameter.
	Value *AttributeMappingSource `json:"value,omitempty"`
}
