package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferencedObject struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Name of the referenced object. Must match one of the objects in the directory definition.
	ReferencedObjectName nullable.Type[string] `json:"referencedObjectName,omitempty"`

	// Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
	ReferencedProperty nullable.Type[string] `json:"referencedProperty,omitempty"`
}
