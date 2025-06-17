package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceSummarizedAppState struct {
	// The unique identifier (DeviceId) associated with the device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the type of execution status of the device management script. This status provides insights into whether
	// the script has been successfully executed, encountered errors, or is pending execution.
	SummarizedAppState *DeviceManagementScriptRunState `json:"summarizedAppState,omitempty"`
}
