package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnsupportedDeviceConfigurationDetail struct {
	// A message explaining why an entity is unsupported.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If message is related to a specific property in the original entity, then the name of that property.
	PropertyName nullable.Type[string] `json:"propertyName,omitempty"`
}
