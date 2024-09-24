package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCLaunchInfo struct {
	// The unique identifier of the Cloud PC.
	CloudPCId nullable.Type[string] `json:"cloudPcId,omitempty"`

	// The connect URL of the Cloud PC.
	CloudPCLaunchUrl nullable.Type[string] `json:"cloudPcLaunchUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the Cloud PC supports switch functionality. If the value is true, it supports switch functionality;
	// otherwise, false.
	Windows365SwitchCompatible nullable.Type[bool] `json:"windows365SwitchCompatible,omitempty"`

	// Indicates the reason the Cloud PC doesn't support switch. CPCOsVersionNotMeetRequirement indicates that the user
	// needs to update their Cloud PC operation system version. CPCHardwareNotMeetRequirement indicates that the Cloud PC
	// needs more CPU or RAM to support the functionality.
	Windows365SwitchNotCompatibleReason nullable.Type[string] `json:"windows365SwitchNotCompatibleReason,omitempty"`
}
