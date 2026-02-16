package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignedRoleDetails struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Role Assignment IDs for the specifc Role Assignments assigned to a user. This property is read-only.
	RoleAssignmentIds *[]string `json:"roleAssignmentIds,omitempty"`

	// Role Definition IDs for the specifc Role Definitions assigned to a user. This property is read-only.
	RoleDefinitionIds *[]string `json:"roleDefinitionIds,omitempty"`
}

var _ json.Marshaler = DeviceAndAppManagementAssignedRoleDetails{}

func (s DeviceAndAppManagementAssignedRoleDetails) MarshalJSON() ([]byte, error) {
	type wrapper DeviceAndAppManagementAssignedRoleDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceAndAppManagementAssignedRoleDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceAndAppManagementAssignedRoleDetails: %+v", err)
	}

	delete(decoded, "roleAssignmentIds")
	delete(decoded, "roleDefinitionIds")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceAndAppManagementAssignedRoleDetails: %+v", err)
	}

	return encoded, nil
}
