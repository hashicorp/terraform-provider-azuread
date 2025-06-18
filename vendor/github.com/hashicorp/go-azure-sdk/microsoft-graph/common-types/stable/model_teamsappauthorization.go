package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppAuthorization struct {
	// The registration ID of the Microsoft Entra app ID associated with the teamsApp.
	ClientAppId nullable.Type[string] `json:"clientAppId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Set of permissions required by the teamsApp.
	RequiredPermissionSet *TeamsAppPermissionSet `json:"requiredPermissionSet,omitempty"`
}
