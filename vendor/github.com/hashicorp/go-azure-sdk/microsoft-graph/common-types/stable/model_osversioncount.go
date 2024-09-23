package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OsVersionCount struct {
	// Count of devices with malware for the OS version
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The Timestamp of the last update for the device count in UTC
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// OS version
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`
}
