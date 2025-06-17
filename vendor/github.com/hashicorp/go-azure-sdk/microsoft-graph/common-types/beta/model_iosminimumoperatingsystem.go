package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 10.0 or
	// later is required to install the app. If 'False', iOS Version 10.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V100 *bool `json:"v10_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 11.0 or
	// later is required to install the app. If 'False', iOS Version 11.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V110 *bool `json:"v11_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 12.0 or
	// later is required to install the app. If 'False', iOS Version 12.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V120 *bool `json:"v12_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 13.0 or
	// later is required to install the app. If 'False', iOS Version 13.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V130 *bool `json:"v13_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 14.0 or
	// later is required to install the app. If 'False', iOS Version 14.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V140 *bool `json:"v14_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 15.0 or
	// later is required to install the app. If 'False', iOS Version 15.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V150 *bool `json:"v15_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 16.0 or
	// later is required to install the app. If 'False', iOS Version 16.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V160 *bool `json:"v16_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 17.0 or
	// later is required to install the app. If 'False', iOS Version 17.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V170 *bool `json:"v17_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 18.0 or
	// later is required to install the app. If 'False', iOS Version 18.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V180 *bool `json:"v18_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 8.0 or
	// later is required to install the app. If 'False', iOS Version 8.0 is not the minimum version. Default value is False.
	// Exactly one of the minimum operating system boolean values will be TRUE.
	V80 *bool `json:"v8_0,omitempty"`

	// Indicates the minimum iOS version support required for the managed device. When 'True', iOS with OS Version 9.0 or
	// later is required to install the app. If 'False', iOS Version 9.0 is not the minimum version. Default value is False.
	// Exactly one of the minimum operating system boolean values will be TRUE.
	V90 *bool `json:"v9_0,omitempty"`
}
