package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultUserRolePermissions struct {
	// Indicates whether the default user role can create applications. This setting corresponds to the Users can register
	// applications setting in the User settings menu in the Microsoft Entra admin center.
	AllowedToCreateApps *bool `json:"allowedToCreateApps,omitempty"`

	// Indicates whether the default user role can create security groups. This setting corresponds to the following menus
	// in the Microsoft Entra admin center: The Users can create security groups in Microsoft Entra admin centers, API or
	// PowerShell setting in the Group settings menu. Users can create security groups setting in the User settings menu.
	AllowedToCreateSecurityGroups *bool `json:"allowedToCreateSecurityGroups,omitempty"`

	// Indicates whether the default user role can create tenants. This setting corresponds to the Restrict non-admin users
	// from creating tenants setting in the User settings menu in the Microsoft Entra admin center. When this setting is
	// false, users assigned the Tenant Creator role can still create tenants.
	AllowedToCreateTenants nullable.Type[bool] `json:"allowedToCreateTenants,omitempty"`

	// Indicates whether the registered owners of a device can read their own BitLocker recovery keys with default user
	// role.
	AllowedToReadBitlockerKeysForOwnedDevice nullable.Type[bool] `json:"allowedToReadBitlockerKeysForOwnedDevice,omitempty"`

	// Indicates whether the default user role can read other users. DO NOT SET THIS VALUE TO false.
	AllowedToReadOtherUsers *bool `json:"allowedToReadOtherUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
