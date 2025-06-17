package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignedRoleDefinition struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A list of permissions based on its associated role. Each permission defines the specific actions the user can perform
	// on Intune resources, such as managing devices, applications, or configurations. Some possible values are:
	// Microsoft.Intune/MobileApps/Read, Microsoft.Intune/DeviceConfigurations/Write,
	// Microsoft.Intune/ManagedDevices/Retire, and Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions
	// property offers a comprehensive view of the user's effective access rights, ensuring that they can only perform
	// actions relevant to their assigned roles. This property is read-only.
	Permissions *[]string `json:"permissions,omitempty"`

	// The RoleDefinitionDisplayName property represents the human-readable name of a specific role definition in Microsoft
	// Intune. This property provides a clear and descriptive name that indicates the purpose or scope of the role, helping
	// administrators identify and assign appropriate roles to users or groups.Some example values for
	// RoleDefinitionDisplayName might include: 'Helpdesk Operator,' 'Application Manager,' or 'Policy Administrator.' This
	// display name is primarily used in the Intune console or Graph API to present roles in a user-friendly manner, making
	// it easier for administrators to manage role-based access control (RBAC) efficiently. This property is read-only.
	RoleDefinitionDisplayName nullable.Type[string] `json:"roleDefinitionDisplayName,omitempty"`
}

var _ json.Marshaler = DeviceAndAppManagementAssignedRoleDefinition{}

func (s DeviceAndAppManagementAssignedRoleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementAssignedRoleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementAssignedRoleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignedRoleDefinition: %+v", err)
	}

	delete(decoded, "permissions")
	delete(decoded, "roleDefinitionDisplayName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementAssignedRoleDefinition: %+v", err)
	}

	return encoded, nil
}
