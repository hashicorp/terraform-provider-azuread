package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyToEvaluate struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Provides the property name.
	PropertyName nullable.Type[string] `json:"propertyName,omitempty"`

	// Provides the property value.
	PropertyValue nullable.Type[string] `json:"propertyValue,omitempty"`
}
