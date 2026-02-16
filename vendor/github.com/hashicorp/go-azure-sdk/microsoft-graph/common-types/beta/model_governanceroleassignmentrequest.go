package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GovernanceRoleAssignmentRequest{}

type GovernanceRoleAssignmentRequest struct {
	// Required. The state of the assignment. The possible values are: Eligible (for eligible assignment), Active (if it is
	// directly assigned), Active (by administrators, or activated on an eligible assignment by the users).
	AssignmentState string `json:"assignmentState"`

	// If this is a request for role activation, it represents the id of the eligible assignment being referred; Otherwise,
	// the value is null.
	LinkedEligibleRoleAssignmentId nullable.Type[string] `json:"linkedEligibleRoleAssignmentId,omitempty"`

	// A message provided by users and administrators when create the request about why it is needed.
	Reason nullable.Type[string] `json:"reason,omitempty"`

	// Read-only. The request create time. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	RequestedDateTime nullable.Type[string] `json:"requestedDateTime,omitempty"`

	// Read-only. The resource that the request aims to.
	Resource *GovernanceResource `json:"resource,omitempty"`

	// Required. The unique identifier of the Azure resource that is associated with the role assignment request. Azure
	// resources can include subscriptions, resource groups, virtual machines, and SQL databases.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Read-only. The role definition that the request aims to.
	RoleDefinition *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`

	// Required. The identifier of the Azure role definition that the role assignment request is associated with.
	RoleDefinitionId nullable.Type[string] `json:"roleDefinitionId,omitempty"`

	// The schedule object of the role assignment request.
	Schedule *GovernanceSchedule `json:"schedule,omitempty"`

	// The status of the role assignment request.
	Status *GovernanceRoleAssignmentRequestStatus `json:"status,omitempty"`

	// Read-only. The user/group principal.
	Subject *GovernanceSubject `json:"subject,omitempty"`

	// Required. The unique identifier of the principal or subject that the role assignment request is associated with.
	// Principals can be users, groups, or service principals.
	SubjectId nullable.Type[string] `json:"subjectId,omitempty"`

	// Required. Representing the type of the operation on the role assignment. The possible values are: AdminAdd , UserAdd
	// , AdminUpdate , AdminRemove , UserRemove , UserExtend , AdminExtend , UserRenew , AdminRenew.
	Type string `json:"type"`

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

func (s GovernanceRoleAssignmentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GovernanceRoleAssignmentRequest{}

func (s GovernanceRoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	type wrapper GovernanceRoleAssignmentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GovernanceRoleAssignmentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GovernanceRoleAssignmentRequest: %+v", err)
	}

	delete(decoded, "requestedDateTime")
	delete(decoded, "resource")
	delete(decoded, "roleDefinition")
	delete(decoded, "subject")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.governanceRoleAssignmentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GovernanceRoleAssignmentRequest: %+v", err)
	}

	return encoded, nil
}
