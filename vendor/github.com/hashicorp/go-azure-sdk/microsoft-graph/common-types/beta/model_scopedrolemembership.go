package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ScopedRoleMembership{}

type ScopedRoleMembership struct {
	// Unique identifier for the administrative unit that the directory role is scoped to
	AdministrativeUnitId *string `json:"administrativeUnitId,omitempty"`

	// Unique identifier for the directory role that the member is in.
	RoleId *string `json:"roleId,omitempty"`

	RoleMemberInfo Identity `json:"roleMemberInfo"`

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

func (s ScopedRoleMembership) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ScopedRoleMembership{}

func (s ScopedRoleMembership) MarshalJSON() ([]byte, error) {
	type wrapper ScopedRoleMembership
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScopedRoleMembership: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScopedRoleMembership: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.scopedRoleMembership"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScopedRoleMembership: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ScopedRoleMembership{}

func (s *ScopedRoleMembership) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdministrativeUnitId *string `json:"administrativeUnitId,omitempty"`
		RoleId               *string `json:"roleId,omitempty"`
		Id                   *string `json:"id,omitempty"`
		ODataId              *string `json:"@odata.id,omitempty"`
		ODataType            *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdministrativeUnitId = decoded.AdministrativeUnitId
	s.RoleId = decoded.RoleId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ScopedRoleMembership into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["roleMemberInfo"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RoleMemberInfo' for 'ScopedRoleMembership': %+v", err)
		}
		s.RoleMemberInfo = impl
	}

	return nil
}
