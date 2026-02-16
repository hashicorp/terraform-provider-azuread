package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyAutopatch struct {
	// The unique identifier (ID) of a Windows Autopatch group. An Autopatch group is a logical container or unit that
	// groups several Microsoft Entra groups and software update policies. Devices with the same Autopatch group ID share
	// unified software update management. The default value is null that indicates that no Autopatch group is associated
	// with the provisioning policy.
	AutopatchGroupId nullable.Type[string] `json:"autopatchGroupId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
