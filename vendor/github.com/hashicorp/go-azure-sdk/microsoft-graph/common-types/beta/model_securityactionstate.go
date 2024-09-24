package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionState struct {
	// The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be
	// extracted from the auth token and not entered manually by the calling application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
	Status *OperationStatus `json:"status,omitempty"`

	// Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	UpdatedDateTime nullable.Type[string] `json:"updatedDateTime,omitempty"`

	// The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be
	// extracted from the auth token and not entered manually by the calling application.
	User nullable.Type[string] `json:"user,omitempty"`
}
