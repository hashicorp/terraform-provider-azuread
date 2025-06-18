package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestActivity struct {
	Action         nullable.Type[string] `json:"action,omitempty"`
	ActionDateTime nullable.Type[string] `json:"actionDateTime,omitempty"`
	Detail         nullable.Type[string] `json:"detail,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ScheduledDateTime nullable.Type[string] `json:"scheduledDateTime,omitempty"`
	UserDisplayName   nullable.Type[string] `json:"userDisplayName,omitempty"`
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
