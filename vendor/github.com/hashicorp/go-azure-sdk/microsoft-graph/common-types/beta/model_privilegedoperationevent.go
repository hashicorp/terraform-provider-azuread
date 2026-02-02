package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrivilegedOperationEvent{}

type PrivilegedOperationEvent struct {
	AdditionalInformation nullable.Type[string] `json:"additionalInformation,omitempty"`
	CreationDateTime      nullable.Type[string] `json:"creationDateTime,omitempty"`
	ExpirationDateTime    nullable.Type[string] `json:"expirationDateTime,omitempty"`
	ReferenceKey          nullable.Type[string] `json:"referenceKey,omitempty"`
	ReferenceSystem       nullable.Type[string] `json:"referenceSystem,omitempty"`
	RequestType           nullable.Type[string] `json:"requestType,omitempty"`
	RequestorId           nullable.Type[string] `json:"requestorId,omitempty"`
	RequestorName         nullable.Type[string] `json:"requestorName,omitempty"`
	RoleId                nullable.Type[string] `json:"roleId,omitempty"`
	RoleName              nullable.Type[string] `json:"roleName,omitempty"`
	TenantId              nullable.Type[string] `json:"tenantId,omitempty"`
	UserId                nullable.Type[string] `json:"userId,omitempty"`
	UserMail              nullable.Type[string] `json:"userMail,omitempty"`
	UserName              nullable.Type[string] `json:"userName,omitempty"`

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

func (s PrivilegedOperationEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrivilegedOperationEvent{}

func (s PrivilegedOperationEvent) MarshalJSON() ([]byte, error) {
	type wrapper PrivilegedOperationEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrivilegedOperationEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegedOperationEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegedOperationEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrivilegedOperationEvent: %+v", err)
	}

	return encoded, nil
}
