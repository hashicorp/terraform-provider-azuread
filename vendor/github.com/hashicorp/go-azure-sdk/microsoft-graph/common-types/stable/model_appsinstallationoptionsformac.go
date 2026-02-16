package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsInstallationOptionsForMac struct {
	// Specifies whether users can install Microsoft 365 apps on their MAC devices. The default value is true.
	IsMicrosoft365AppsEnabled *bool `json:"isMicrosoft365AppsEnabled,omitempty"`

	// Specifies whether users can install Skype for Business on their MAC devices running OS X El Capitan 10.11 or later.
	// The default value is true.
	IsSkypeForBusinessEnabled *bool `json:"isSkypeForBusinessEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
