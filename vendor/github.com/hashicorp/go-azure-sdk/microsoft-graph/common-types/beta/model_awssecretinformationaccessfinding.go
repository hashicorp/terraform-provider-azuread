package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecretInformationAccessFinding interface {
	Entity
	Finding
	AwsSecretInformationAccessFinding() BaseAwsSecretInformationAccessFindingImpl
}

var _ AwsSecretInformationAccessFinding = BaseAwsSecretInformationAccessFindingImpl{}

type BaseAwsSecretInformationAccessFindingImpl struct {
	Identity                     *AuthorizationSystemIdentity     `json:"identity,omitempty"`
	IdentityDetails              *IdentityDetails                 `json:"identityDetails,omitempty"`
	PermissionsCreepIndex        *PermissionsCreepIndex           `json:"permissionsCreepIndex,omitempty"`
	SecretInformationWebServices *AwsSecretInformationWebServices `json:"secretInformationWebServices,omitempty"`

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

func (s BaseAwsSecretInformationAccessFindingImpl) AwsSecretInformationAccessFinding() BaseAwsSecretInformationAccessFindingImpl {
	return s
}

func (s BaseAwsSecretInformationAccessFindingImpl) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseAwsSecretInformationAccessFindingImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AwsSecretInformationAccessFinding = RawAwsSecretInformationAccessFindingImpl{}

// RawAwsSecretInformationAccessFindingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAwsSecretInformationAccessFindingImpl struct {
	awsSecretInformationAccessFinding BaseAwsSecretInformationAccessFindingImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawAwsSecretInformationAccessFindingImpl) AwsSecretInformationAccessFinding() BaseAwsSecretInformationAccessFindingImpl {
	return s.awsSecretInformationAccessFinding
}

func (s RawAwsSecretInformationAccessFindingImpl) Finding() BaseFindingImpl {
	return s.awsSecretInformationAccessFinding.Finding()
}

func (s RawAwsSecretInformationAccessFindingImpl) Entity() BaseEntityImpl {
	return s.awsSecretInformationAccessFinding.Entity()
}

var _ json.Marshaler = BaseAwsSecretInformationAccessFindingImpl{}

func (s BaseAwsSecretInformationAccessFindingImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAwsSecretInformationAccessFindingImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAwsSecretInformationAccessFindingImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAwsSecretInformationAccessFindingImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsSecretInformationAccessFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAwsSecretInformationAccessFindingImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAwsSecretInformationAccessFindingImpl{}

func (s *BaseAwsSecretInformationAccessFindingImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IdentityDetails              *IdentityDetails                 `json:"identityDetails,omitempty"`
		PermissionsCreepIndex        *PermissionsCreepIndex           `json:"permissionsCreepIndex,omitempty"`
		SecretInformationWebServices *AwsSecretInformationWebServices `json:"secretInformationWebServices,omitempty"`
		CreatedDateTime              *string                          `json:"createdDateTime,omitempty"`
		Id                           *string                          `json:"id,omitempty"`
		ODataId                      *string                          `json:"@odata.id,omitempty"`
		ODataType                    *string                          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IdentityDetails = decoded.IdentityDetails
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.SecretInformationWebServices = decoded.SecretInformationWebServices
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAwsSecretInformationAccessFindingImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BaseAwsSecretInformationAccessFindingImpl': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}

func UnmarshalAwsSecretInformationAccessFindingImplementation(input []byte) (AwsSecretInformationAccessFinding, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsSecretInformationAccessFinding into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsResourceFinding") {
		var out SecretInformationAccessAwsResourceFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsResourceFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsRoleFinding") {
		var out SecretInformationAccessAwsRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsServerlessFunctionFinding") {
		var out SecretInformationAccessAwsServerlessFunctionFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsServerlessFunctionFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secretInformationAccessAwsUserFinding") {
		var out SecretInformationAccessAwsUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecretInformationAccessAwsUserFinding: %+v", err)
		}
		return out, nil
	}

	var parent BaseAwsSecretInformationAccessFindingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAwsSecretInformationAccessFindingImpl: %+v", err)
	}

	return RawAwsSecretInformationAccessFindingImpl{
		awsSecretInformationAccessFinding: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
