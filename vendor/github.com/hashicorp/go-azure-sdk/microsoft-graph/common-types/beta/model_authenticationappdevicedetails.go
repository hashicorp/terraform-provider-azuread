package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppDeviceDetails struct {
	// The version of the client authentication app used during the authentication step.
	AppVersion nullable.Type[string] `json:"appVersion,omitempty"`

	// The name of the client authentication app used during the authentication step.
	ClientApp nullable.Type[string] `json:"clientApp,omitempty"`

	// ID of the device used during the authentication step.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operating system running on the device used for the authentication step.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`
}
