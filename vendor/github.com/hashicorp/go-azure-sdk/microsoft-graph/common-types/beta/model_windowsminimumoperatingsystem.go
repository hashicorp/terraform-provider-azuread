package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Windows version 10.0 or later.
	V100 *bool `json:"v10_0,omitempty"`

	// Windows 10 1607 or later.
	V101607 *bool `json:"v10_1607,omitempty"`

	// Windows 10 1703 or later.
	V101703 *bool `json:"v10_1703,omitempty"`

	// Windows 10 1709 or later.
	V101709 *bool `json:"v10_1709,omitempty"`

	// Windows 10 1803 or later.
	V101803 *bool `json:"v10_1803,omitempty"`

	// Windows 10 1809 or later.
	V101809 *bool `json:"v10_1809,omitempty"`

	// Windows 10 1903 or later.
	V101903 *bool `json:"v10_1903,omitempty"`

	// Windows 10 1909 or later.
	V101909 *bool `json:"v10_1909,omitempty"`

	// Windows 10 2004 or later.
	V102004 *bool `json:"v10_2004,omitempty"`

	// Windows 10 21H1 or later.
	V1021H1 *bool `json:"v10_21H1,omitempty"`

	// Windows 10 2H20 or later.
	V102H20 *bool `json:"v10_2H20,omitempty"`

	// Windows version 8.0 or later.
	V80 *bool `json:"v8_0,omitempty"`

	// Windows version 8.1 or later.
	V81 *bool `json:"v8_1,omitempty"`
}
