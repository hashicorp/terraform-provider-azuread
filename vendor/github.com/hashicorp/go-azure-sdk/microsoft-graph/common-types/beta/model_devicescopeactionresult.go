package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceScopeActionResult struct {
	// Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
	DeviceScopeAction *DeviceScopeAction `json:"deviceScopeAction,omitempty"`

	// The unique identifier of the device scope the action was triggered on.
	DeviceScopeId nullable.Type[string] `json:"deviceScopeId,omitempty"`

	// The message indicates the reason the device scope action failed to trigger.
	FailedMessage nullable.Type[string] `json:"failedMessage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the status of the attempted device scope action
	Status *DeviceScopeActionStatus `json:"status,omitempty"`
}
