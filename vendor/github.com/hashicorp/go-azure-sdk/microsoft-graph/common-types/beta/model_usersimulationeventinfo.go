package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSimulationEventInfo struct {
	// Browser information from where the simulation event was initiated by a user in an attack simulation and training
	// campaign.
	Browser nullable.Type[string] `json:"browser,omitempty"`

	// Date and time of the simulation event by a user in an attack simulation and training campaign.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// Name of the simulation event by a user in an attack simulation and training campaign.
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	// IP address from where the simulation event was initiated by a user in an attack simulation and training campaign.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operating system, platform, and device details from where the simulation event was initiated by a user in an
	// attack simulation and training campaign.
	OsPlatformDeviceDetails nullable.Type[string] `json:"osPlatformDeviceDetails,omitempty"`
}
