package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerRosterMember{}

type PlannerRosterMember struct {
	// Additional roles associated with the PlannerRosterMember, which determines permissions of the member in the
	// plannerRoster. Currently there are no available roles to assign, and every member has full control over the contents
	// of the plannerRoster.
	Roles *[]string `json:"roles,omitempty"`

	// Identifier of the tenant the user belongs to. Currently only the users from the same tenant can be added to a
	// plannerRoster.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// Identifier of the user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s PlannerRosterMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerRosterMember{}

func (s PlannerRosterMember) MarshalJSON() ([]byte, error) {
	type wrapper PlannerRosterMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerRosterMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerRosterMember: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerRosterMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerRosterMember: %+v", err)
	}

	return encoded, nil
}
