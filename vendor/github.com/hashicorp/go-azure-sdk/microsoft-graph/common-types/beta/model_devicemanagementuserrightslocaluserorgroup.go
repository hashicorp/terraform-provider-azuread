package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementUserRightsLocalUserOrGroup struct {
	// Admin’s description of this local user or group.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of this local user or group.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The security identifier of this local user or group (e.g. S-1-5-32-544).
	SecurityIdentifier nullable.Type[string] `json:"securityIdentifier,omitempty"`
}
