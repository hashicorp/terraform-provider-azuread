package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceDetail struct {
	// Indicates the browser information of the used in the sign-in. Populated for devices registered in Microsoft Entra.
	Browser nullable.Type[string] `json:"browser,omitempty"`

	// Refers to the unique ID of the device used in the sign-in. Populated for devices registered in Microsoft Entra.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Refers to the name of the device used in the sign-in. Populated for devices registered in Microsoft Entra.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the device is compliant or not.
	IsCompliant nullable.Type[bool] `json:"isCompliant,omitempty"`

	// Indicates if the device is managed or not.
	IsManaged nullable.Type[bool] `json:"isManaged,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the OS name and version used in the sign-in.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// Indicates information on whether the device used in the sign-in is workplace-joined, Microsoft Entra-joined,
	// domain-joined.
	TrustType nullable.Type[string] `json:"trustType,omitempty"`
}
