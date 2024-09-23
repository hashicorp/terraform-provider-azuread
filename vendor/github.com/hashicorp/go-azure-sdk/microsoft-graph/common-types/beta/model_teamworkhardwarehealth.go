package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHardwareHealth struct {
	// The system health details for a teamworkDevice.
	ComputeHealth *TeamworkPeripheralHealth `json:"computeHealth,omitempty"`

	// The health details about the HDMI ingest of a device.
	HdmiIngestHealth *TeamworkPeripheralHealth `json:"hdmiIngestHealth,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
