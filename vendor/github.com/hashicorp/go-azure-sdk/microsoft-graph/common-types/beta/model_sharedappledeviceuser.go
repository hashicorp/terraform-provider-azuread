package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedAppleDeviceUser struct {
	// Data quota
	DataQuota nullable.Type[int64] `json:"dataQuota,omitempty"`

	// Data to sync
	DataToSync *bool `json:"dataToSync,omitempty"`

	// Data quota
	DataUsed *int64 `json:"dataUsed,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// User name
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
