package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = AwsExternalSystemAccessFinding{}

type AwsExternalSystemAccessFinding struct {
	AccessMethods    *ExternalSystemAccessMethods `json:"accessMethods,omitempty"`
	AffectedSystem   *AuthorizationSystem         `json:"affectedSystem,omitempty"`
	SystemWithAccess *AuthorizationSystemInfo     `json:"systemWithAccess,omitempty"`

	// The number of identities in the external system that are trusted, if not all. Supports $orderby.
	TrustedIdentityCount nullable.Type[int64] `json:"trustedIdentityCount,omitempty"`

	// Flag that determines if all identities in the external system are trusted, or only a subset.
	TrustsAllIdentities *bool `json:"trustsAllIdentities,omitempty"`

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

func (s AwsExternalSystemAccessFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AwsExternalSystemAccessFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsExternalSystemAccessFinding{}

func (s AwsExternalSystemAccessFinding) MarshalJSON() ([]byte, error) {
	type wrapper AwsExternalSystemAccessFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsExternalSystemAccessFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsExternalSystemAccessFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsExternalSystemAccessFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsExternalSystemAccessFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AwsExternalSystemAccessFinding{}

func (s *AwsExternalSystemAccessFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessMethods        *ExternalSystemAccessMethods `json:"accessMethods,omitempty"`
		SystemWithAccess     *AuthorizationSystemInfo     `json:"systemWithAccess,omitempty"`
		TrustedIdentityCount nullable.Type[int64]         `json:"trustedIdentityCount,omitempty"`
		TrustsAllIdentities  *bool                        `json:"trustsAllIdentities,omitempty"`
		CreatedDateTime      *string                      `json:"createdDateTime,omitempty"`
		Id                   *string                      `json:"id,omitempty"`
		ODataId              *string                      `json:"@odata.id,omitempty"`
		ODataType            *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessMethods = decoded.AccessMethods
	s.SystemWithAccess = decoded.SystemWithAccess
	s.TrustedIdentityCount = decoded.TrustedIdentityCount
	s.TrustsAllIdentities = decoded.TrustsAllIdentities
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AwsExternalSystemAccessFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["affectedSystem"]; ok {
		impl, err := UnmarshalAuthorizationSystemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AffectedSystem' for 'AwsExternalSystemAccessFinding': %+v", err)
		}
		s.AffectedSystem = &impl
	}

	return nil
}
