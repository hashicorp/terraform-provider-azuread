package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeEscalationFinding interface {
	Entity
	Finding
	PrivilegeEscalationFinding() BasePrivilegeEscalationFindingImpl
}

var _ PrivilegeEscalationFinding = BasePrivilegeEscalationFindingImpl{}

type BasePrivilegeEscalationFindingImpl struct {
	Identity *AuthorizationSystemIdentity `json:"identity,omitempty"`

	// An identity's information details. Inherited from finding.
	IdentityDetails *IdentityDetails `json:"identityDetails,omitempty"`

	PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`

	// The list of escalations that the identity is capable of performing.
	PrivilegeEscalationDetails *[]PrivilegeEscalation `json:"privilegeEscalationDetails,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

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

func (s BasePrivilegeEscalationFindingImpl) PrivilegeEscalationFinding() BasePrivilegeEscalationFindingImpl {
	return s
}

func (s BasePrivilegeEscalationFindingImpl) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BasePrivilegeEscalationFindingImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ PrivilegeEscalationFinding = RawPrivilegeEscalationFindingImpl{}

// RawPrivilegeEscalationFindingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPrivilegeEscalationFindingImpl struct {
	privilegeEscalationFinding BasePrivilegeEscalationFindingImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawPrivilegeEscalationFindingImpl) PrivilegeEscalationFinding() BasePrivilegeEscalationFindingImpl {
	return s.privilegeEscalationFinding
}

func (s RawPrivilegeEscalationFindingImpl) Finding() BaseFindingImpl {
	return s.privilegeEscalationFinding.Finding()
}

func (s RawPrivilegeEscalationFindingImpl) Entity() BaseEntityImpl {
	return s.privilegeEscalationFinding.Entity()
}

var _ json.Marshaler = BasePrivilegeEscalationFindingImpl{}

func (s BasePrivilegeEscalationFindingImpl) MarshalJSON() ([]byte, error) {
	type wrapper BasePrivilegeEscalationFindingImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasePrivilegeEscalationFindingImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasePrivilegeEscalationFindingImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.privilegeEscalationFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasePrivilegeEscalationFindingImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BasePrivilegeEscalationFindingImpl{}

func (s *BasePrivilegeEscalationFindingImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IdentityDetails            *IdentityDetails       `json:"identityDetails,omitempty"`
		PermissionsCreepIndex      *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
		PrivilegeEscalationDetails *[]PrivilegeEscalation `json:"privilegeEscalationDetails,omitempty"`
		CreatedDateTime            *string                `json:"createdDateTime,omitempty"`
		Id                         *string                `json:"id,omitempty"`
		ODataId                    *string                `json:"@odata.id,omitempty"`
		ODataType                  *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IdentityDetails = decoded.IdentityDetails
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.PrivilegeEscalationDetails = decoded.PrivilegeEscalationDetails
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BasePrivilegeEscalationFindingImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BasePrivilegeEscalationFindingImpl': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}

func UnmarshalPrivilegeEscalationFindingImplementation(input []byte) (PrivilegeEscalationFinding, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PrivilegeEscalationFinding into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationAwsResourceFinding") {
		var out PrivilegeEscalationAwsResourceFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationAwsResourceFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationAwsRoleFinding") {
		var out PrivilegeEscalationAwsRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationAwsRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationGcpServiceAccountFinding") {
		var out PrivilegeEscalationGcpServiceAccountFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationGcpServiceAccountFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.privilegeEscalationUserFinding") {
		var out PrivilegeEscalationUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PrivilegeEscalationUserFinding: %+v", err)
		}
		return out, nil
	}

	var parent BasePrivilegeEscalationFindingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePrivilegeEscalationFindingImpl: %+v", err)
	}

	return RawPrivilegeEscalationFindingImpl{
		privilegeEscalationFinding: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
