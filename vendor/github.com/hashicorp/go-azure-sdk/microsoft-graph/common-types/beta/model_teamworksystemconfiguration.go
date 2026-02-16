package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSystemConfiguration struct {
	// The date and time configurations for a device.
	DateTimeConfiguration *TeamworkDateTimeConfiguration `json:"dateTimeConfiguration,omitempty"`

	// The default password for the device. Write-Only.
	DefaultPassword nullable.Type[string] `json:"defaultPassword,omitempty"`

	// The device lock timeout in seconds.
	DeviceLockTimeout nullable.Type[string] `json:"deviceLockTimeout,omitempty"`

	// True if the device lock is enabled.
	IsDeviceLockEnabled nullable.Type[bool] `json:"isDeviceLockEnabled,omitempty"`

	// True if logging is enabled.
	IsLoggingEnabled nullable.Type[bool] `json:"isLoggingEnabled,omitempty"`

	// True if power saving is enabled.
	IsPowerSavingEnabled nullable.Type[bool] `json:"isPowerSavingEnabled,omitempty"`

	// True if screen capture is enabled.
	IsScreenCaptureEnabled nullable.Type[bool] `json:"isScreenCaptureEnabled,omitempty"`

	// True if silent mode is enabled.
	IsSilentModeEnabled nullable.Type[bool] `json:"isSilentModeEnabled,omitempty"`

	// The language option for the device.
	Language nullable.Type[string] `json:"language,omitempty"`

	// The pin that unlocks the device. Write-Only.
	LockPin nullable.Type[string] `json:"lockPin,omitempty"`

	// The logging level for the device.
	LoggingLevel nullable.Type[string] `json:"loggingLevel,omitempty"`

	// The network configuration for the device.
	NetworkConfiguration *TeamworkNetworkConfiguration `json:"networkConfiguration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
