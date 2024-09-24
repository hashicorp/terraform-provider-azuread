package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsSecretInformationAccessFinding = SecretInformationAccessAwsServerlessFunctionFinding{}

type SecretInformationAccessAwsServerlessFunctionFinding struct {

	// Fields inherited from AwsSecretInformationAccessFinding

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

func (s SecretInformationAccessAwsServerlessFunctionFinding) AwsSecretInformationAccessFinding() BaseAwsSecretInformationAccessFindingImpl {
	return BaseAwsSecretInformationAccessFindingImpl{
		Identity:                     s.Identity,
		IdentityDetails:              s.IdentityDetails,
		PermissionsCreepIndex:        s.PermissionsCreepIndex,
		SecretInformationWebServices: s.SecretInformationWebServices,
		CreatedDateTime:              s.CreatedDateTime,
		Id:                           s.Id,
		ODataId:                      s.ODataId,
		ODataType:                    s.ODataType,
	}
}

func (s SecretInformationAccessAwsServerlessFunctionFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SecretInformationAccessAwsServerlessFunctionFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecretInformationAccessAwsServerlessFunctionFinding{}

func (s SecretInformationAccessAwsServerlessFunctionFinding) MarshalJSON() ([]byte, error) {
	type wrapper SecretInformationAccessAwsServerlessFunctionFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecretInformationAccessAwsServerlessFunctionFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecretInformationAccessAwsServerlessFunctionFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.secretInformationAccessAwsServerlessFunctionFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecretInformationAccessAwsServerlessFunctionFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecretInformationAccessAwsServerlessFunctionFinding{}

func (s *SecretInformationAccessAwsServerlessFunctionFinding) UnmarshalJSON(bytes []byte) error {
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

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.IdentityDetails = decoded.IdentityDetails
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.SecretInformationWebServices = decoded.SecretInformationWebServices

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecretInformationAccessAwsServerlessFunctionFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'SecretInformationAccessAwsServerlessFunctionFinding': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}
