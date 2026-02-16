package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDefinition struct {
	// Defines attributes of the object.
	Attributes *[]AttributeDefinition `json:"attributes,omitempty"`

	// Metadata for the given object.
	Metadata *[]ObjectDefinitionMetadataEntry `json:"metadata,omitempty"`

	// Name of the object. Must be unique within a directory definition. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The API that the provisioning service queries to retrieve data for synchronization.
	SupportedApis *[]string `json:"supportedApis,omitempty"`
}
