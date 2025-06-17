package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftIncidentDetail struct {
	// TThe name of the application for which the diagnostic is collected. Example: com.microsoft.CompanyPortal
	ApplicationName nullable.Type[string] `json:"applicationName,omitempty"`

	// The version of the application for which the diagnostic is collected. Example: 5.2203.1
	ClientApplicationVersion nullable.Type[string] `json:"clientApplicationVersion,omitempty"`

	// The time the app diagnostic was created. The value cannot be modified and is automatically populated when the
	// diagnostic is uploaded. The Timestamp type represents date and time information using ISO 8601 format and is always
	// in UTC time.Example: 2022-04-19T17:24:45.313Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The unique app diagnostic identifier as a user friendly 8 character hexadecimal string. This id is smaller compared
	// to the powerliftId. Th Example: 8520467A
	EasyId nullable.Type[string] `json:"easyId,omitempty"`

	// A list of files that are associated with the diagnostic.
	FileNames *[]string `json:"fileNames,omitempty"`

	// The locale information of the application for which the diagnostic is collected. Example: en-US
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operating system of the device from which diagnostics are collected. Example: iOS
	PlatformDisplayName nullable.Type[string] `json:"platformDisplayName,omitempty"`

	// The unique identifier of the app diagnostic. This id is assigned to a diagnostic when it is uploaded to Powerlift.
	// Example: 8520467a-49a9-44a4-8447-8dfb8bec6726
	PowerliftId nullable.Type[string] `json:"powerliftId,omitempty"`
}
