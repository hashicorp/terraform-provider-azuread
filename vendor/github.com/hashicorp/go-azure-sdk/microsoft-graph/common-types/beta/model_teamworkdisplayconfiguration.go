package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDisplayConfiguration struct {
	// The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
	ConfiguredDisplays *[]TeamworkConfiguredPeripheral `json:"configuredDisplays,omitempty"`

	// Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
	DisplayCount nullable.Type[int64] `json:"displayCount,omitempty"`

	// Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
	InBuiltDisplayScreenConfiguration *TeamworkDisplayScreenConfiguration `json:"inBuiltDisplayScreenConfiguration,omitempty"`

	// True if content duplication is allowed. Applicable only for Teams Rooms devices.
	IsContentDuplicationAllowed nullable.Type[bool] `json:"isContentDuplicationAllowed,omitempty"`

	// True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both
	// front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft
	// Teams Rooms device. Applicable only for Teams Rooms devices.
	IsDualDisplayModeEnabled nullable.Type[bool] `json:"isDualDisplayModeEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
