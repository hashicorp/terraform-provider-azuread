package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationTargetedUserAndDevice struct {
	// The id of the device in the checkin.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The name of the device in the checkin.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Last checkin time for this user/device pair.
	LastCheckinDateTime *string `json:"lastCheckinDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The display name of the user in the checkin
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The id of the user in the checkin.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The UPN of the user in the checkin.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
