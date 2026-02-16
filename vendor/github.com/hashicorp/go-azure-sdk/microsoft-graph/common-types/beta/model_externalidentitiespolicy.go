package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyBase = ExternalIdentitiesPolicy{}

type ExternalIdentitiesPolicy struct {
	// Reserved for future use.
	AllowDeletedIdentitiesDataRemoval *bool `json:"allowDeletedIdentitiesDataRemoval,omitempty"`

	// Defines whether external users can leave the guest tenant. If set to false, self-service controls are disabled, and
	// the admin of the guest tenant must manually remove the external user from the guest tenant. When the external user
	// leaves the tenant, their data in the guest tenant is first soft-deleted then permanently deleted in 30 days.
	AllowExternalIdentitiesToLeave *bool `json:"allowExternalIdentitiesToLeave,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description string `json:"description"`

	// Display name for this policy. Required.
	DisplayName string `json:"displayName"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s ExternalIdentitiesPolicy) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExternalIdentitiesPolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExternalIdentitiesPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalIdentitiesPolicy{}

func (s ExternalIdentitiesPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ExternalIdentitiesPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalIdentitiesPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalIdentitiesPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalIdentitiesPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalIdentitiesPolicy: %+v", err)
	}

	return encoded, nil
}
