package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoSettings struct {
	// Controls whether users can join lists from users external to your organization.
	IsExternalJoinEnabled nullable.Type[bool] `json:"isExternalJoinEnabled,omitempty"`

	// Controls whether users can share lists with external users.
	IsExternalShareEnabled nullable.Type[bool] `json:"isExternalShareEnabled,omitempty"`

	// Controls whether push notifications are enabled for your users.
	IsPushNotificationEnabled nullable.Type[bool] `json:"isPushNotificationEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
