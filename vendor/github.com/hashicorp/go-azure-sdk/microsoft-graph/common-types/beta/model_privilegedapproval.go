package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedApproval{}

type PrivilegedApproval struct {
	ApprovalDuration nullable.Type[string]            `json:"approvalDuration,omitempty"`
	ApprovalState    *ApprovalState                   `json:"approvalState,omitempty"`
	ApprovalType     nullable.Type[string]            `json:"approvalType,omitempty"`
	ApproverReason   nullable.Type[string]            `json:"approverReason,omitempty"`
	EndDateTime      nullable.Type[string]            `json:"endDateTime,omitempty"`
	Request          *PrivilegedRoleAssignmentRequest `json:"request,omitempty"`
	RequestorReason  nullable.Type[string]            `json:"requestorReason,omitempty"`
	RoleId           nullable.Type[string]            `json:"roleId,omitempty"`
	RoleInfo         *PrivilegedRole                  `json:"roleInfo,omitempty"`
	StartDateTime    nullable.Type[string]            `json:"startDateTime,omitempty"`
	UserId           nullable.Type[string]            `json:"userId,omitempty"`

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

func (s PrivilegedApproval) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedApproval{}

func (s PrivilegedApproval) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedApproval
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedApproval: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedApproval: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedApproval"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedApproval: %+v", err)
	}

	return encoded, nil
}
