package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceDetail struct {
	// Indicates the browser information of the used for signing-in.
	Browser nullable.Type[string] `json:"browser,omitempty"`

	BrowserId nullable.Type[string] `json:"browserId,omitempty"`

	// Refers to the UniqueID of the device used for signing-in.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Refers to the name of the device used for signing-in.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the device is compliant or not.
	IsCompliant nullable.Type[bool] `json:"isCompliant,omitempty"`

	// Indicates if the device is managed or not.
	IsManaged nullable.Type[bool] `json:"isManaged,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the OS name and version used for signing-in.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// Indicates information on whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
	TrustType nullable.Type[string] `json:"trustType,omitempty"`
}
