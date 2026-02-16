package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedRoleSummary{}

type PrivilegedRoleSummary struct {
	ElevatedCount nullable.Type[int64] `json:"elevatedCount,omitempty"`
	ManagedCount  nullable.Type[int64] `json:"managedCount,omitempty"`
	MfaEnabled    nullable.Type[bool]  `json:"mfaEnabled,omitempty"`
	Status        *RoleSummaryStatus   `json:"status,omitempty"`
	UsersCount    nullable.Type[int64] `json:"usersCount,omitempty"`

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

func (s PrivilegedRoleSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedRoleSummary{}

func (s PrivilegedRoleSummary) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedRoleSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedRoleSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedRoleSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedRoleSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedRoleSummary: %+v", err)
	}

	return encoded, nil
}
