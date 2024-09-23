package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiInformation struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the package type of an MSI Win32LobApp.
	PackageType *Win32LobAppMsiPackageType `json:"packageType,omitempty"`

	// The MSI product code.
	ProductCode nullable.Type[string] `json:"productCode,omitempty"`

	// The MSI product name.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The MSI product version.
	ProductVersion nullable.Type[string] `json:"productVersion,omitempty"`

	// The MSI publisher.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Whether the MSI app requires the machine to reboot to complete installation.
	RequiresReboot *bool `json:"requiresReboot,omitempty"`

	// The MSI upgrade code.
	UpgradeCode nullable.Type[string] `json:"upgradeCode,omitempty"`
}
