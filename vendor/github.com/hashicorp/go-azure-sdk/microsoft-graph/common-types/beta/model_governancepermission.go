package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernancePermission struct {
	// The access level. Valid values: None, UserRead, AdminRead, and AdminReadWrite.
	AccessLevel nullable.Type[string] `json:"accessLevel,omitempty"`

	// Indicate if the requestor has any active role assignment for the access level.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// Indicate if the requestor has any eligible role assignment for the access level.
	IsEligible nullable.Type[bool] `json:"isEligible,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
