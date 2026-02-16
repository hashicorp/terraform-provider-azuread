package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// When TRUE, indicates OS X 10.10 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1010 *bool `json:"v10_10,omitempty"`

	// When TRUE, indicates OS X 10.11 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1011 *bool `json:"v10_11,omitempty"`

	// When TRUE, indicates macOS 10.12 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1012 *bool `json:"v10_12,omitempty"`

	// When TRUE, indicates macOS 10.13 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1013 *bool `json:"v10_13,omitempty"`

	// When TRUE, indicates macOS 10.14 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1014 *bool `json:"v10_14,omitempty"`

	// When TRUE, indicates macOS 10.15 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V1015 *bool `json:"v10_15,omitempty"`

	// When TRUE, indicates Mac OS X 10.7 or later is required to install the app. When FALSE, indicates some other OS
	// version is the minimum OS to install the app. Default value is FALSE.
	V107 *bool `json:"v10_7,omitempty"`

	// When TRUE, indicates OS X 10.8 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V108 *bool `json:"v10_8,omitempty"`

	// When TRUE, indicates OS X 10.9 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V109 *bool `json:"v10_9,omitempty"`

	// When TRUE, indicates macOS 11.0 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V110 *bool `json:"v11_0,omitempty"`

	// When TRUE, indicates macOS 12.0 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V120 *bool `json:"v12_0,omitempty"`

	// When TRUE, indicates macOS 13.0 or later is required to install the app. When FALSE, indicates some other OS version
	// is the minimum OS to install the app. Default value is FALSE.
	V130 *bool `json:"v13_0,omitempty"`
}
