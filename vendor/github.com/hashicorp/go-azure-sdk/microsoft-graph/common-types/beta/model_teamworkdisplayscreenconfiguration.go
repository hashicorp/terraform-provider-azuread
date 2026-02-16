package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayScreenConfiguration struct {
	// The brightness level on the device (0-100). Not applicable for Microsoft Teams Rooms devices.
	BacklightBrightness nullable.Type[int64] `json:"backlightBrightness,omitempty"`

	// Timeout for backlight (30-3600 secs). Not applicable for Teams Rooms devices.
	BacklightTimeout nullable.Type[string] `json:"backlightTimeout,omitempty"`

	// True if high contrast mode is enabled. Not applicable for Teams Rooms devices.
	IsHighContrastEnabled nullable.Type[bool] `json:"isHighContrastEnabled,omitempty"`

	// True if screensaver is enabled. Not applicable for Teams Rooms devices.
	IsScreensaverEnabled nullable.Type[bool] `json:"isScreensaverEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Screensaver timeout from 30 to 3600 secs. Not applicable for Teams Rooms devices.
	ScreensaverTimeout nullable.Type[string] `json:"screensaverTimeout,omitempty"`
}
