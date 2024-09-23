package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreAuthorizedApplication struct {
	// The unique identifier for the application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The unique identifier for the oauth2PermissionScopes the application requires.
	DelegatedPermissionIds *[]string `json:"delegatedPermissionIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
