package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecurityToolAdministrationFinding interface {
	Entity
	Finding
	AwsSecurityToolAdministrationFinding() BaseAwsSecurityToolAdministrationFindingImpl
}

var _ AwsSecurityToolAdministrationFinding = BaseAwsSecurityToolAdministrationFindingImpl{}

type BaseAwsSecurityToolAdministrationFindingImpl struct {
	Identity              *AuthorizationSystemIdentity `json:"identity,omitempty"`
	IdentityDetails       *IdentityDetails             `json:"identityDetails,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex       `json:"permissionsCreepIndex,omitempty"`
	SecurityTools         *AwsSecurityToolWebServices  `json:"securityTools,omitempty"`

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

func (s BaseAwsSecurityToolAdministrationFindingImpl) AwsSecurityToolAdministrationFinding() BaseAwsSecurityToolAdministrationFindingImpl {
	return s
}

func (s BaseAwsSecurityToolAdministrationFindingImpl) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseAwsSecurityToolAdministrationFindingImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AwsSecurityToolAdministrationFinding = RawAwsSecurityToolAdministrationFindingImpl{}

// RawAwsSecurityToolAdministrationFindingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAwsSecurityToolAdministrationFindingImpl struct {
	awsSecurityToolAdministrationFinding BaseAwsSecurityToolAdministrationFindingImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawAwsSecurityToolAdministrationFindingImpl) AwsSecurityToolAdministrationFinding() BaseAwsSecurityToolAdministrationFindingImpl {
	return s.awsSecurityToolAdministrationFinding
}

func (s RawAwsSecurityToolAdministrationFindingImpl) Finding() BaseFindingImpl {
	return s.awsSecurityToolAdministrationFinding.Finding()
}

func (s RawAwsSecurityToolAdministrationFindingImpl) Entity() BaseEntityImpl {
	return s.awsSecurityToolAdministrationFinding.Entity()
}

var _ json.Marshaler = BaseAwsSecurityToolAdministrationFindingImpl{}

func (s BaseAwsSecurityToolAdministrationFindingImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAwsSecurityToolAdministrationFindingImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAwsSecurityToolAdministrationFindingImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAwsSecurityToolAdministrationFindingImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsSecurityToolAdministrationFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAwsSecurityToolAdministrationFindingImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAwsSecurityToolAdministrationFindingImpl{}

func (s *BaseAwsSecurityToolAdministrationFindingImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IdentityDetails       *IdentityDetails            `json:"identityDetails,omitempty"`
		PermissionsCreepIndex *PermissionsCreepIndex      `json:"permissionsCreepIndex,omitempty"`
		SecurityTools         *AwsSecurityToolWebServices `json:"securityTools,omitempty"`
		CreatedDateTime       *string                     `json:"createdDateTime,omitempty"`
		Id                    *string                     `json:"id,omitempty"`
		ODataId               *string                     `json:"@odata.id,omitempty"`
		ODataType             *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IdentityDetails = decoded.IdentityDetails
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.SecurityTools = decoded.SecurityTools
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAwsSecurityToolAdministrationFindingImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BaseAwsSecurityToolAdministrationFindingImpl': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}

func UnmarshalAwsSecurityToolAdministrationFindingImplementation(input []byte) (AwsSecurityToolAdministrationFinding, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsSecurityToolAdministrationFinding into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsResourceAdministratorFinding") {
		var out SecurityToolAwsResourceAdministratorFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsResourceAdministratorFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsRoleAdministratorFinding") {
		var out SecurityToolAwsRoleAdministratorFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsRoleAdministratorFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsServerlessFunctionAdministratorFinding") {
		var out SecurityToolAwsServerlessFunctionAdministratorFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsServerlessFunctionAdministratorFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.securityToolAwsUserAdministratorFinding") {
		var out SecurityToolAwsUserAdministratorFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityToolAwsUserAdministratorFinding: %+v", err)
		}
		return out, nil
	}

	var parent BaseAwsSecurityToolAdministrationFindingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAwsSecurityToolAdministrationFindingImpl: %+v", err)
	}

	return RawAwsSecurityToolAdministrationFindingImpl{
		awsSecurityToolAdministrationFinding: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
