package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosAvailableUpdateVersion struct {
	// The expiration date of the update.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The posting date of the update.
	PostingDateTime *string `json:"postingDateTime,omitempty"`

	// The version of the update.
	ProductVersion *string `json:"productVersion,omitempty"`

	// List of supported devices for the update.
	SupportedDevices *[]string `json:"supportedDevices,omitempty"`
}
