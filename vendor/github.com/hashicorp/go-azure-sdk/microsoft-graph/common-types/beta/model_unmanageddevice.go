package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDevice struct {
	// Device name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Domain.
	Domain nullable.Type[string] `json:"domain,omitempty"`

	// IP address.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Last logged on user.
	LastLoggedOnUser nullable.Type[string] `json:"lastLoggedOnUser,omitempty"`

	// Last seen date and time.
	LastSeenDateTime *string `json:"lastSeenDateTime,omitempty"`

	// Location.
	Location nullable.Type[string] `json:"location,omitempty"`

	// MAC address.
	MacAddress nullable.Type[string] `json:"macAddress,omitempty"`

	// Manufacturer.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Model.
	Model nullable.Type[string] `json:"model,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Operating system.
	Os nullable.Type[string] `json:"os,omitempty"`

	// Operating system version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`
}
