package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyInboundTrust struct {
	// Specifies whether compliant devices from external Microsoft Entra organizations are trusted.
	IsCompliantDeviceAccepted nullable.Type[bool] `json:"isCompliantDeviceAccepted,omitempty"`

	// Specifies whether Microsoft Entra hybrid joined devices from external Microsoft Entra organizations are trusted.
	IsHybridAzureADJoinedDeviceAccepted nullable.Type[bool] `json:"isHybridAzureADJoinedDeviceAccepted,omitempty"`

	// Specifies whether MFA from external Microsoft Entra organizations is trusted.
	IsMfaAccepted nullable.Type[bool] `json:"isMfaAccepted,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
