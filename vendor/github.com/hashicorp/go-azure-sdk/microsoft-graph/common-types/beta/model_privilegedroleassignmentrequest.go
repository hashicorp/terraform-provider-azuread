package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedRoleAssignmentRequest{}

type PrivilegedRoleAssignmentRequest struct {
	AssignmentState   nullable.Type[string] `json:"assignmentState,omitempty"`
	Duration          nullable.Type[string] `json:"duration,omitempty"`
	Reason            nullable.Type[string] `json:"reason,omitempty"`
	RequestedDateTime nullable.Type[string] `json:"requestedDateTime,omitempty"`
	RoleId            nullable.Type[string] `json:"roleId,omitempty"`
	RoleInfo          *PrivilegedRole       `json:"roleInfo,omitempty"`
	Schedule          *GovernanceSchedule   `json:"schedule,omitempty"`
	Status            nullable.Type[string] `json:"status,omitempty"`
	TicketNumber      nullable.Type[string] `json:"ticketNumber,omitempty"`
	TicketSystem      nullable.Type[string] `json:"ticketSystem,omitempty"`
	Type              nullable.Type[string] `json:"type,omitempty"`
	UserId            nullable.Type[string] `json:"userId,omitempty"`

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

func (s PrivilegedRoleAssignmentRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedRoleAssignmentRequest{}

func (s PrivilegedRoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedRoleAssignmentRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedRoleAssignmentRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedRoleAssignmentRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedRoleAssignmentRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedRoleAssignmentRequest: %+v", err)
	}

	return encoded, nil
}
