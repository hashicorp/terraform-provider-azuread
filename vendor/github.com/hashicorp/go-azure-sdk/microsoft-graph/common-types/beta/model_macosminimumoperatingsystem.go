package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.10 or
	// later is required to install the app. If 'False', OS X Version 10.10 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1010 *bool `json:"v10_10,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.11 or
	// later is required to install the app. If 'False', OS X Version 10.11 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1011 *bool `json:"v10_11,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.12 or
	// later is required to install the app. If 'False', OS X Version 10.12 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1012 *bool `json:"v10_12,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.13 or
	// later is required to install the app. If 'False', OS X Version 10.13 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1013 *bool `json:"v10_13,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.14 or
	// later is required to install the app. If 'False', OS X Version 10.14 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1014 *bool `json:"v10_14,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.15 or
	// later is required to install the app. If 'False', OS X Version 10.15 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V1015 *bool `json:"v10_15,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.7 or
	// later is required to install the app. If 'False', OS X Version 10.7 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V107 *bool `json:"v10_7,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.8 or
	// later is required to install the app. If 'False', OS X Version 10.8 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V108 *bool `json:"v10_8,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 10.9 or
	// later is required to install the app. If 'False', OS X Version 10.9 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V109 *bool `json:"v10_9,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 11.0 or
	// later is required to install the app. If 'False', OS X Version 11.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V110 *bool `json:"v11_0,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 12.0 or
	// later is required to install the app. If 'False', OS X Version 12.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V120 *bool `json:"v12_0,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 13.0 or
	// later is required to install the app. If 'False', OS X Version 13.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V130 *bool `json:"v13_0,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 14.0 or
	// later is required to install the app. If 'False', OS X Version 14.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V140 *bool `json:"v14_0,omitempty"`

	// Indicates the minimum OS X version support required for the managed device. When 'True', macOS with OS X 15.0 or
	// later is required to install the app. If 'False', OS X Version 15.0 is not the minimum version. Default value is
	// False. Exactly one of the minimum operating system boolean values will be TRUE.
	V150 *bool `json:"v15_0,omitempty"`
}
