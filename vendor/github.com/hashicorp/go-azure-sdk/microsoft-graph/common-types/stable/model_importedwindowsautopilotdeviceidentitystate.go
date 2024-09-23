package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityState struct {
	// Device error code reported by Device Directory Service(DDS).
	DeviceErrorCode *int64 `json:"deviceErrorCode,omitempty"`

	// Device error name reported by Device Directory Service(DDS).
	DeviceErrorName nullable.Type[string] `json:"deviceErrorName,omitempty"`

	DeviceImportStatus *ImportedWindowsAutopilotDeviceIdentityImportStatus `json:"deviceImportStatus,omitempty"`

	// Device Registration ID for successfully added device reported by Device Directory Service(DDS).
	DeviceRegistrationId nullable.Type[string] `json:"deviceRegistrationId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
