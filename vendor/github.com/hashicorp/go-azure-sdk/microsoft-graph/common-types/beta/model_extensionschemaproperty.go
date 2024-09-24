package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionSchemaProperty struct {
	// The name of the strongly typed property defined as part of a schema extension.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of the property that is defined as part of a schema extension. Allowed values are Binary, Boolean, DateTime,
	// Integer, or String. For more information, see Supported property data types.
	Type nullable.Type[string] `json:"type,omitempty"`
}
