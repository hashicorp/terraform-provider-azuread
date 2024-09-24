package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceKey struct {
	DeviceId    nullable.Type[string] `json:"deviceId,omitempty"`
	KeyMaterial nullable.Type[string] `json:"keyMaterial,omitempty"`
	KeyType     nullable.Type[string] `json:"keyType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
