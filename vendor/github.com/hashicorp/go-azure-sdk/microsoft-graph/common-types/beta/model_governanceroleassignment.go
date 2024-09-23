package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceRoleAssignment{}

type GovernanceRoleAssignment struct {
	// The state of the assignment. The value can be Eligible for eligible assignment or Active if it's directly assigned
	// Active by administrators, or activated on an eligible assignment by the users.
	AssignmentState *string `json:"assignmentState,omitempty"`

	// For a non-permanent role assignment, this is the time when the role assignment is expired. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The external ID the resource that is used to identify the role assignment in the provider.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Read-only. If this is an active assignment and created due to activation on an eligible assignment, it represents the
	// object of that eligible assignment; Otherwise, the value is null.
	LinkedEligibleRoleAssignment *GovernanceRoleAssignment `json:"linkedEligibleRoleAssignment,omitempty"`

	// If this is an active assignment and created due to activation on an eligible assignment, it represents the ID of that
	// eligible assignment; Otherwise, the value is null.
	LinkedEligibleRoleAssignmentId nullable.Type[string] `json:"linkedEligibleRoleAssignmentId,omitempty"`

	// The type of member. The value can be: Inherited (if the role assignment is inherited from a parent resource scope),
	// Group (if the role assignment isn't inherited, but comes from the membership of a group assignment), or User (if the
	// role assignment isn't inherited or from a group assignment).
	MemberType *string `json:"memberType,omitempty"`

	// Read-only. The resource associated with the role assignment.
	Resource *GovernanceResource `json:"resource,omitempty"`

	// Required. The ID of the resource that the role assignment is associated with.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Read-only. The role definition associated with the role assignment.
	RoleDefinition *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`

	// Required. The ID of the role definition that the role assignment is associated with.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The start time of the role assignment. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	Status *string `json:"status,omitempty"`

	// Read-only. The subject associated with the role assignment.
	Subject *GovernanceSubject `json:"subject,omitempty"`

	// Required. The ID of the subject that the role assignment is associated with.
	SubjectId nullable.Type[string] `json:"subjectId,omitempty"`

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

func (s GovernanceRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceRoleAssignment{}

func (s GovernanceRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceRoleAssignment: %+v", err)
	}

	delete(decoded, "linkedEligibleRoleAssignment")
	delete(decoded, "resource")
	delete(decoded, "roleDefinition")
	delete(decoded, "subject")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceRoleAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceRoleAssignment: %+v", err)
	}

	return encoded, nil
}
