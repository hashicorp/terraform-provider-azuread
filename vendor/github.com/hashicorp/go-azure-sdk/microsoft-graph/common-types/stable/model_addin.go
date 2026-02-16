package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddIn struct {
	// The unique identifier for the addIn object.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The collection of key-value pairs that define parameters that the consuming service can use or call. You must specify
	// this property when performing a POST or a PATCH operation on the addIns collection. Required.
	Properties []KeyValue `json:"properties"`

	// The unique name for the functionality exposed by the app.
	Type *string `json:"type,omitempty"`
}
