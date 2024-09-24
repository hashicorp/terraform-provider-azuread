package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperience struct {
	// Indicates the type of restart action.
	DeviceRestartBehavior *Win32LobAppRestartBehavior `json:"deviceRestartBehavior,omitempty"`

	// The number of minutes the system will wait for install program to finish. Default value is 60 minutes.
	MaxRunTimeInMinutes nullable.Type[int64] `json:"maxRunTimeInMinutes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
}
