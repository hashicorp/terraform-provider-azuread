package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationSettingState struct {
	// Current value of setting on device
	CurrentValue nullable.Type[string] `json:"currentValue,omitempty"`

	// Error code for the setting
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// Error description
	ErrorDescription nullable.Type[string] `json:"errorDescription,omitempty"`

	// Name of setting instance that is being reported.
	InstanceDisplayName nullable.Type[string] `json:"instanceDisplayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The setting that is being reported
	Setting nullable.Type[string] `json:"setting,omitempty"`

	// SettingInstanceId
	SettingInstanceId nullable.Type[string] `json:"settingInstanceId,omitempty"`

	// Localized/user friendly setting name that is being reported
	SettingName nullable.Type[string] `json:"settingName,omitempty"`

	// Contributing policies
	Sources *[]SettingSource `json:"sources,omitempty"`

	State *ComplianceStatus `json:"state,omitempty"`

	// UserEmail
	UserEmail nullable.Type[string] `json:"userEmail,omitempty"`

	// UserId
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// UserName
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// UserPrincipalName.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`
}
