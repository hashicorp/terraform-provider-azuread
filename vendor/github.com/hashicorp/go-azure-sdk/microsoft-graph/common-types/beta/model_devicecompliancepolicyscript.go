package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyScript struct {
	// Device compliance script Id.
	DeviceComplianceScriptId nullable.Type[string] `json:"deviceComplianceScriptId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Json of the rules.
	RulesContent nullable.Type[string] `json:"rulesContent,omitempty"`
}
