package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RbacApplication{}

type RbacApplication struct {
	ResourceNamespaces *[]UnifiedRbacResourceNamespace `json:"resourceNamespaces,omitempty"`

	// Instances for active role assignments.
	RoleAssignmentScheduleInstances *[]UnifiedRoleAssignmentScheduleInstance `json:"roleAssignmentScheduleInstances,omitempty"`

	// Requests for active role assignments to principals through PIM.
	RoleAssignmentScheduleRequests *[]UnifiedRoleAssignmentScheduleRequest `json:"roleAssignmentScheduleRequests,omitempty"`

	// Schedules for active role assignment operations.
	RoleAssignmentSchedules *[]UnifiedRoleAssignmentSchedule `json:"roleAssignmentSchedules,omitempty"`

	// Resource to grant access to users or groups.
	RoleAssignments *[]UnifiedRoleAssignment `json:"roleAssignments,omitempty"`

	// Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
	RoleDefinitions *[]UnifiedRoleDefinition `json:"roleDefinitions,omitempty"`

	// Instances for role eligibility requests.
	RoleEligibilityScheduleInstances *[]UnifiedRoleEligibilityScheduleInstance `json:"roleEligibilityScheduleInstances,omitempty"`

	// Requests for role eligibilities for principals through PIM.
	RoleEligibilityScheduleRequests *[]UnifiedRoleEligibilityScheduleRequest `json:"roleEligibilityScheduleRequests,omitempty"`

	// Schedules for role eligibility operations.
	RoleEligibilitySchedules *[]UnifiedRoleEligibilitySchedule `json:"roleEligibilitySchedules,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s RbacApplication) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RbacApplication{}

func (s RbacApplication) MarshalJSON() ([]byte, error) {
	type wrapper RbacApplication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RbacApplication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RbacApplication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.rbacApplication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RbacApplication: %+v", err)
	}

	return encoded, nil
}
