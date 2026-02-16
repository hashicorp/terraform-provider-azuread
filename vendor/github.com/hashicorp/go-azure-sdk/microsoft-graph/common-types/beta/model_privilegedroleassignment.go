package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedRoleAssignment{}

type PrivilegedRoleAssignment struct {
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`
	IsElevated         nullable.Type[bool]   `json:"isElevated,omitempty"`
	ResultMessage      nullable.Type[string] `json:"resultMessage,omitempty"`
	RoleId             *string               `json:"roleId,omitempty"`
	RoleInfo           *PrivilegedRole       `json:"roleInfo,omitempty"`
	UserId             *string               `json:"userId,omitempty"`

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

func (s PrivilegedRoleAssignment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedRoleAssignment{}

func (s PrivilegedRoleAssignment) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedRoleAssignment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedRoleAssignment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedRoleAssignment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedRoleAssignment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedRoleAssignment: %+v", err)
	}

	return encoded, nil
}
