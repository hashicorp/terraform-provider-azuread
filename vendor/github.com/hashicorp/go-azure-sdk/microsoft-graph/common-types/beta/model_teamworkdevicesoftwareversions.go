package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceSoftwareVersions struct {
	// The software version for the admin agent running on the device.
	AdminAgentSoftwareVersion nullable.Type[string] `json:"adminAgentSoftwareVersion,omitempty"`

	// The software version for the firmware running on the device.
	FirmwareSoftwareVersion nullable.Type[string] `json:"firmwareSoftwareVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The software version for the operating system on the device.
	OperatingSystemSoftwareVersion nullable.Type[string] `json:"operatingSystemSoftwareVersion,omitempty"`

	// The software version for the partner agent running on the device.
	PartnerAgentSoftwareVersion nullable.Type[string] `json:"partnerAgentSoftwareVersion,omitempty"`

	// The software version for the Teams client running on the device.
	TeamsClientSoftwareVersion nullable.Type[string] `json:"teamsClientSoftwareVersion,omitempty"`
}
