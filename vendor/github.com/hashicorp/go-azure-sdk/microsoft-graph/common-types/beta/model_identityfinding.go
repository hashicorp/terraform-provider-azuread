package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityFinding interface {
	Entity
	Finding
	IdentityFinding() BaseIdentityFindingImpl
}

var _ IdentityFinding = BaseIdentityFindingImpl{}

type BaseIdentityFindingImpl struct {
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

func (s BaseIdentityFindingImpl) IdentityFinding() BaseIdentityFindingImpl {
	return s
}

func (s BaseIdentityFindingImpl) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s BaseIdentityFindingImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IdentityFinding = RawIdentityFindingImpl{}

// RawIdentityFindingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityFindingImpl struct {
	identityFinding BaseIdentityFindingImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawIdentityFindingImpl) IdentityFinding() BaseIdentityFindingImpl {
	return s.identityFinding
}

func (s RawIdentityFindingImpl) Finding() BaseFindingImpl {
	return s.identityFinding.Finding()
}

func (s RawIdentityFindingImpl) Entity() BaseEntityImpl {
	return s.identityFinding.Entity()
}

var _ json.Marshaler = BaseIdentityFindingImpl{}

func (s BaseIdentityFindingImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIdentityFindingImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIdentityFindingImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIdentityFindingImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIdentityFindingImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseIdentityFindingImpl{}

func (s *BaseIdentityFindingImpl) UnmarshalJSON(bytes []byte) error {
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
	s.IdentityDetails = decoded.IdentityDetails
	s.PermissionsCreepIndex = decoded.PermissionsCreepIndex
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseIdentityFindingImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalAuthorizationSystemIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BaseIdentityFindingImpl': %+v", err)
		}
		s.Identity = &impl
	}

	return nil
}

func UnmarshalIdentityFindingImplementation(input []byte) (IdentityFinding, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityFinding into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveAwsResourceFinding") {
		var out InactiveAwsResourceFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAwsResourceFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveAwsRoleFinding") {
		var out InactiveAwsRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAwsRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveAzureServicePrincipalFinding") {
		var out InactiveAzureServicePrincipalFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveAzureServicePrincipalFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveGcpServiceAccountFinding") {
		var out InactiveGcpServiceAccountFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveGcpServiceAccountFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveServerlessFunctionFinding") {
		var out InactiveServerlessFunctionFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveServerlessFunctionFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.inactiveUserFinding") {
		var out InactiveUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InactiveUserFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAwsResourceFinding") {
		var out OverprovisionedAwsResourceFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAwsResourceFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAwsRoleFinding") {
		var out OverprovisionedAwsRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAwsRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedAzureServicePrincipalFinding") {
		var out OverprovisionedAzureServicePrincipalFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedAzureServicePrincipalFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedGcpServiceAccountFinding") {
		var out OverprovisionedGcpServiceAccountFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedGcpServiceAccountFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedServerlessFunctionFinding") {
		var out OverprovisionedServerlessFunctionFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedServerlessFunctionFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.overprovisionedUserFinding") {
		var out OverprovisionedUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OverprovisionedUserFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superAwsResourceFinding") {
		var out SuperAwsResourceFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAwsResourceFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superAwsRoleFinding") {
		var out SuperAwsRoleFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAwsRoleFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superAzureServicePrincipalFinding") {
		var out SuperAzureServicePrincipalFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperAzureServicePrincipalFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superGcpServiceAccountFinding") {
		var out SuperGcpServiceAccountFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperGcpServiceAccountFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superServerlessFunctionFinding") {
		var out SuperServerlessFunctionFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperServerlessFunctionFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.superUserFinding") {
		var out SuperUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SuperUserFinding: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.unenforcedMfaAwsUserFinding") {
		var out UnenforcedMfaAwsUserFinding
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UnenforcedMfaAwsUserFinding: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityFindingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityFindingImpl: %+v", err)
	}

	return RawIdentityFindingImpl{
		identityFinding: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
