package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyBase = IdentitySecurityDefaultsEnforcementPolicy{}

type IdentitySecurityDefaultsEnforcementPolicy struct {
	// If set to true, Microsoft Entra security defaults are enabled for the tenant.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// Fields inherited from PolicyBase

	// Description for this policy. Required.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name for this policy. Required.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

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

func (s IdentitySecurityDefaultsEnforcementPolicy) PolicyBase() BasePolicyBaseImpl {
	return BasePolicyBaseImpl{
		Description:     s.Description,
		DisplayName:     s.DisplayName,
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s IdentitySecurityDefaultsEnforcementPolicy) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s IdentitySecurityDefaultsEnforcementPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentitySecurityDefaultsEnforcementPolicy{}

func (s IdentitySecurityDefaultsEnforcementPolicy) MarshalJSON() ([]byte, error) {
	type wrapper IdentitySecurityDefaultsEnforcementPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentitySecurityDefaultsEnforcementPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentitySecurityDefaultsEnforcementPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identitySecurityDefaultsEnforcementPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentitySecurityDefaultsEnforcementPolicy: %+v", err)
	}

	return encoded, nil
}
