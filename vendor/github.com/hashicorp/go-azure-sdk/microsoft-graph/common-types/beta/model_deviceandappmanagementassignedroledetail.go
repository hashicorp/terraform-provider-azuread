package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignedRoleDetail struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The list of permissions assigned to a specific user based on their associated role definitions. Each permission
	// defines the specific actions the user can perform on Intune resources, such as managing devices, applications, or
	// configurations. Some possible values are: Microsoft.Intune/MobileApps/Read,
	// Microsoft.Intune/DeviceConfigurations/Write, Microsoft.Intune/ManagedDevices/Retire, and
	// Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions property provides a comprehensive view of the
	// user's effective access rights, ensuring that they can only perform actions relevant to their assigned roles. This
	// property is read-only.
	Permissions *[]string `json:"permissions,omitempty"`

	// A collection of RoleDefinitions represents the various administrative roles that define permissions and access levels
	// within Microsoft Intune. Each RoleDefinition outlines a set of permissions that determine the actions an admin or
	// user can perform in the Intune environment. These permissions can include actions like reading or writing to specific
	// resources, managing device configurations, deploying policies, or handling user data. RoleDefinitions are critical
	// for enforcing role-based access control (RBAC), ensuring that administrators can only interact with the features and
	// data relevant to their responsibilities. RoleDefinitions in Intune can either be built-in roles provided by Microsoft
	// or custom roles created by an organization to tailor access based on specific needs. These definitions are referenced
	// when assigning roles to users or groups, effectively controlling the scope of their administrative privileges. The
	// collection of RoleDefinitions is managed through the Intune console or the Graph API, allowing for scalable role
	// management across large environments. This property is read-only.
	RoleDefinitions *[]DeviceAndAppManagementAssignedRoleDefinition `json:"roleDefinitions,omitempty"`
}

var _ json.Marshaler = DeviceAndAppManagementAssignedRoleDetail{}

func (s DeviceAndAppManagementAssignedRoleDetail) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementAssignedRoleDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementAssignedRoleDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignedRoleDetail: %+v", err)
	}

	delete(decoded, "permissions")
	delete(decoded, "roleDefinitions")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementAssignedRoleDetail: %+v", err)
	}

	return encoded, nil
}
