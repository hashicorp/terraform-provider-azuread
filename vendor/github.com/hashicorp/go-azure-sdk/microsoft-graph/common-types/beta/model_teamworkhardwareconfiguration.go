package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkHardwareConfiguration struct {
	Compute    *TeamworkPeripheral `json:"compute,omitempty"`
	HdmiIngest *TeamworkPeripheral `json:"hdmiIngest,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The CPU model on the device.
	ProcessorModel nullable.Type[string] `json:"processorModel,omitempty"`
}
