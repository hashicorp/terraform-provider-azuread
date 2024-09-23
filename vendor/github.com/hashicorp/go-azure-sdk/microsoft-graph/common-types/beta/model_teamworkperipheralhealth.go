package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkPeripheralHealth struct {
	// The connected state and time since the peripheral device was connected.
	Connection *TeamworkConnection `json:"connection,omitempty"`

	// True if the peripheral is optional. Used for health computation.
	IsOptional nullable.Type[bool] `json:"isOptional,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Peripheral *TeamworkPeripheral `json:"peripheral,omitempty"`
}
