package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegationSettings{}

type DelegationSettings struct {
	// The allowed actions for the delegator or delegate.
	AllowedActions *DelegateAllowedActions `json:"allowedActions,omitempty"`

	// Date and time when the delegator or delegate entry was created. The DateTimeOffset type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates whether the delegator or delegate relationship is currently active.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

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

func (s DelegationSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegationSettings{}

func (s DelegationSettings) MarshalJSON() ([]byte, error) {
	type wrapper DelegationSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegationSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegationSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegationSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegationSettings: %+v", err)
	}

	return encoded, nil
}
