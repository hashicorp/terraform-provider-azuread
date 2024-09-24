package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTrainingContentEventInfo struct {
	// Browser of the user from where the training event was generated.
	Browser nullable.Type[string] `json:"browser,omitempty"`

	// Date and time of the training content playback by the user.
	ContentDateTime nullable.Type[string] `json:"contentDateTime,omitempty"`

	// IP address of the user for the training event.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operating system, platform, and device details of the user for the training event.
	OsPlatformDeviceDetails nullable.Type[string] `json:"osPlatformDeviceDetails,omitempty"`
}
