package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestriction struct {
	// Collection of blocked Manufacturers.
	BlockedManufacturers *[]string `json:"blockedManufacturers,omitempty"`

	// Collection of blocked Skus.
	BlockedSkus *[]string `json:"blockedSkus,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Max OS version supported
	OsMaximumVersion nullable.Type[string] `json:"osMaximumVersion,omitempty"`

	// Min OS version supported
	OsMinimumVersion nullable.Type[string] `json:"osMinimumVersion,omitempty"`

	// Block personally owned devices from enrolling
	PersonalDeviceEnrollmentBlocked *bool `json:"personalDeviceEnrollmentBlocked,omitempty"`

	// Block the platform from enrolling
	PlatformBlocked *bool `json:"platformBlocked,omitempty"`
}
