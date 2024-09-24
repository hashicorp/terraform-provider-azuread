package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantUserSyncInbound struct {
	// Defines whether user objects should be synchronized from the partner tenant. false causes any current user
	// synchronization from the source tenant to the target tenant to stop. This property has no impact on existing users
	// who have already been synchronized.
	IsSyncAllowed nullable.Type[bool] `json:"isSyncAllowed,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
