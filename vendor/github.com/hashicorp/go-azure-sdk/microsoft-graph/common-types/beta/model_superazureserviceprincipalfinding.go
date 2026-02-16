package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityFinding = SuperAzureServicePrincipalFinding{}

type SuperAzureServicePrincipalFinding struct {

	// Fields inherited from IdentityFinding

	ActionSummary *ActionSummary               `json:"actionSummary,omitempty"`
	Identity      *AuthorizationSystemIdentity `json:"identity,omitempty"`

	// An identity's information details.
	IdentityDetails *IdentityDetails `json:"identityDetails,omitempty"`

	PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`

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

func (s SuperAzureServicePrincipalFinding) IdentityFinding() BaseIdentityFindingImpl {
	return BaseIdentityFindingImpl{
		ActionSummary:         s.ActionSummary,
		Identity:              s.Identity,
		IdentityDetails:       s.IdentityDetails,
		PermissionsCreepIndex: s.PermissionsCreepIndex,
		CreatedDateTime:       s.CreatedDateTime,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s SuperAzureServicePrincipalFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s SuperAzureServicePrincipalFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SuperAzureServicePrincipalFinding{}

func (s SuperAzureServicePrincipalFinding) MarshalJSON() ([]byte, error) {
	type wrapper SuperAzureServicePrincipalFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SuperAzureServicePrincipalFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SuperAzureServicePrincipalFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.superAzureServicePrincipalFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SuperAzureServicePrincipalFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SuperAzureServicePrincipalFinding{}

func (s *SuperAzureServicePrincipalFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActionSummary         *ActionSummary         `json:"actionSummary,omitempty"`
		IdentityDetails       *IdentityDetails       `json:"identityDetails,omitempty"`
		PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
		CreatedDateTime       *string                `json:"createdDateTime,omitempty"`
		Id                    *string                `json:"id,omitempty"`
		ODataId               *string                `json:"@odata.id,omitempty"`
		ODataType             *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActionSummary = decoded.ActionSummary
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.IdentityDetails = decoded.IdentityDetails
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SuperAzureServicePrincipalFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'SuperAzureServicePrincipalFinding': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}
