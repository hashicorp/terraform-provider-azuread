package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHardwareDetail struct {
	// MAC address.
	MacAddresses *[]string `json:"macAddresses,omitempty"`

	// Device manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Devie model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Device serial number.
	SerialNumber nullable.Type[string] `json:"serialNumber,omitempty"`

	// The unique identifier for the device.
	UniqueId nullable.Type[string] `json:"uniqueId,omitempty"`
}
