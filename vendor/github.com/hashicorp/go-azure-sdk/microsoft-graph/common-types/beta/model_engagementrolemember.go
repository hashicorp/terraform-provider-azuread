package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EngagementRoleMember{}

type EngagementRoleMember struct {
	// The timestamp when the role was assigned to the user.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// User entity of the member who has been assigned the role.
	User *User `json:"user,omitempty"`

	// The Microsoft Entra ID of the user who has the role assigned.
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

func (s EngagementRoleMember) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EngagementRoleMember{}

func (s EngagementRoleMember) MarshalJSON() ([]byte, error) {
	type wrapper EngagementRoleMember
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EngagementRoleMember: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EngagementRoleMember: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "userId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.engagementRoleMember"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EngagementRoleMember: %+v", err)
	}

	return encoded, nil
}
