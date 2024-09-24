package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftIncidentMetadata struct {
	// The name of the application the diagnostic is from. Example: com.microsoft.CompanyPortal
	Application nullable.Type[string] `json:"application,omitempty"`

	// The version of the application. Example: 5.2203.1
	ClientVersion nullable.Type[string] `json:"clientVersion,omitempty"`

	// The time the app diagnostic was created. Example: 2022-04-19T17:24:45.313Z
	CreatedAtDateTime nullable.Type[string] `json:"createdAtDateTime,omitempty"`

	// The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. Example: 8520467A
	EasyId nullable.Type[string] `json:"easyId,omitempty"`

	// A list of files that are associated with the diagnostic.
	FileNames *[]string `json:"fileNames,omitempty"`

	// The locale information of the application. Example: en-US
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The device's OS the diagnostic is from. Example: iOS
	Platform nullable.Type[string] `json:"platform,omitempty"`

	// The unique identifier of the app diagnostic. Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
	PowerliftId *string `json:"powerliftId,omitempty"`
}
