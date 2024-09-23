package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesFilter struct {
	// Determines whether devices that satisfy the rule should be allowed or blocked. The possible values are: allowed,
	// blocked, unknownFutureValue.
	Mode *CrossTenantAccessPolicyTargetConfigurationAccessType `json:"mode,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the rule to filter the devices. For example, device.deviceAttribute2 -eq 'PrivilegedAccessWorkstation'.
	Rule nullable.Type[string] `json:"rule,omitempty"`
}
