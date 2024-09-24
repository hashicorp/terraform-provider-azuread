package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleDeviceMode struct {
	// Windows 10 Device Mode type.
	DeviceMode *Windows10DeviceModeType `json:"deviceMode,omitempty"`

	// Name for object.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Supported Applicability rule types for Device Configuration
	RuleType *DeviceManagementApplicabilityRuleType `json:"ruleType,omitempty"`
}
