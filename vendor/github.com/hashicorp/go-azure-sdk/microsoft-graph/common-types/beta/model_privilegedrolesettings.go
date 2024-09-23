package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedRoleSettings{}

type PrivilegedRoleSettings struct {
	ApprovalOnElevation           nullable.Type[bool]   `json:"approvalOnElevation,omitempty"`
	ApproverIds                   *[]string             `json:"approverIds,omitempty"`
	ElevationDuration             nullable.Type[string] `json:"elevationDuration,omitempty"`
	IsMfaOnElevationConfigurable  nullable.Type[bool]   `json:"isMfaOnElevationConfigurable,omitempty"`
	LastGlobalAdmin               nullable.Type[bool]   `json:"lastGlobalAdmin,omitempty"`
	MaxElavationDuration          nullable.Type[string] `json:"maxElavationDuration,omitempty"`
	MfaOnElevation                nullable.Type[bool]   `json:"mfaOnElevation,omitempty"`
	MinElevationDuration          nullable.Type[string] `json:"minElevationDuration,omitempty"`
	NotificationToUserOnElevation nullable.Type[bool]   `json:"notificationToUserOnElevation,omitempty"`
	TicketingInfoOnElevation      nullable.Type[bool]   `json:"ticketingInfoOnElevation,omitempty"`

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

func (s PrivilegedRoleSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedRoleSettings{}

func (s PrivilegedRoleSettings) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedRoleSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedRoleSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedRoleSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedRoleSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedRoleSettings: %+v", err)
	}

	return encoded, nil
}
