package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupLifecyclePolicy{}

type GroupLifecyclePolicy struct {
	// List of email address to send notifications for groups without owners. Multiple email address can be defined by
	// separating email address with a semicolon.
	AlternateNotificationEmails nullable.Type[string] `json:"alternateNotificationEmails,omitempty"`

	// Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the
	// number of days defined.
	GroupLifetimeInDays nullable.Type[int64] `json:"groupLifetimeInDays,omitempty"`

	// The group type for which the expiration policy applies. Possible values are All, Selected or None.
	ManagedGroupTypes nullable.Type[string] `json:"managedGroupTypes,omitempty"`

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

func (s GroupLifecyclePolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupLifecyclePolicy{}

func (s GroupLifecyclePolicy) MarshalJSON() ([]byte, error) {
	type wrapper GroupLifecyclePolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupLifecyclePolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupLifecyclePolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupLifecyclePolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupLifecyclePolicy: %+v", err)
	}

	return encoded, nil
}
