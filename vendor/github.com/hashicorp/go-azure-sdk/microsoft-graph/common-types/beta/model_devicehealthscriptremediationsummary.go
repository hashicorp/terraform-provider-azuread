package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRemediationSummary struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of devices remediated by device health scripts.
	RemediatedDeviceCount nullable.Type[int64] `json:"remediatedDeviceCount,omitempty"`

	// The number of device health scripts deployed.
	ScriptCount nullable.Type[int64] `json:"scriptCount,omitempty"`
}
