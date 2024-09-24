package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementUserRightsSetting struct {
	// Representing a collection of local users or groups which will be set on device if the state of this setting is
	// Allowed. This collection can contain a maximum of 500 elements.
	LocalUsersOrGroups *[]DeviceManagementUserRightsLocalUserOrGroup `json:"localUsersOrGroups,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// State Management Setting.
	State *StateManagementSetting `json:"state,omitempty"`
}
