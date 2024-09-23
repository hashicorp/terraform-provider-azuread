package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardSettings struct {
	// List of safeguards to ignore per device.
	DisabledSafeguardProfiles *[]WindowsUpdatesSafeguardProfile `json:"disabledSafeguardProfiles,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
