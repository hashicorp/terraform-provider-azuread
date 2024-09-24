package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UnifiedRoleManagementAlertIncident = RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident{}

type RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident struct {
	// Display name of the subject that the incident applies to.
	AssigneeDisplayName nullable.Type[string] `json:"assigneeDisplayName,omitempty"`

	// The identifier of the subject that the incident applies to.
	AssigneeId nullable.Type[string] `json:"assigneeId,omitempty"`

	// User principal name of the subject that the incident applies to. Applies to user principals.
	AssigneeUserPrincipalName nullable.Type[string] `json:"assigneeUserPrincipalName,omitempty"`

	// Date and time of assignment creation.
	AssignmentCreatedDateTime nullable.Type[string] `json:"assignmentCreatedDateTime,omitempty"`

	// The identifier for the directory role definition that's in scope of this incident.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The display name for the directory role.
	RoleDisplayName nullable.Type[string] `json:"roleDisplayName,omitempty"`

	// The globally unique identifier for the directory role.
	RoleTemplateId nullable.Type[string] `json:"roleTemplateId,omitempty"`

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

func (s RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) UnifiedRoleManagementAlertIncident() BaseUnifiedRoleManagementAlertIncidentImpl {
	return BaseUnifiedRoleManagementAlertIncidentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident{}

func (s RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident) MarshalJSON() ([]byte, error) {
	type wrapper RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.rolesAssignedOutsidePrivilegedIdentityManagementAlertIncident"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RolesAssignedOutsidePrivilegedIdentityManagementAlertIncident: %+v", err)
	}

	return encoded, nil
}
