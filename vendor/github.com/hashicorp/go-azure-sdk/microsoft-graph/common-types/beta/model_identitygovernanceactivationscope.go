package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceActivationScope interface {
	IdentityGovernanceActivationScope() BaseIdentityGovernanceActivationScopeImpl
}

var _ IdentityGovernanceActivationScope = BaseIdentityGovernanceActivationScopeImpl{}

type BaseIdentityGovernanceActivationScopeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentityGovernanceActivationScopeImpl) IdentityGovernanceActivationScope() BaseIdentityGovernanceActivationScopeImpl {
	return s
}

var _ IdentityGovernanceActivationScope = RawIdentityGovernanceActivationScopeImpl{}

// RawIdentityGovernanceActivationScopeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIdentityGovernanceActivationScopeImpl struct {
	identityGovernanceActivationScope BaseIdentityGovernanceActivationScopeImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawIdentityGovernanceActivationScopeImpl) IdentityGovernanceActivationScope() BaseIdentityGovernanceActivationScopeImpl {
	return s.identityGovernanceActivationScope
}

func (s RawIdentityGovernanceActivationScopeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIdentityGovernanceActivationScopeImplementation(input []byte) (IdentityGovernanceActivationScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceActivationScope into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.activateGroupScope") {
		var out IdentityGovernanceActivateGroupScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceActivateGroupScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.activateProcessingResultScope") {
		var out IdentityGovernanceActivateProcessingResultScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceActivateProcessingResultScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.activateRunScope") {
		var out IdentityGovernanceActivateRunScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceActivateRunScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.activateUserScope") {
		var out IdentityGovernanceActivateUserScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceActivateUserScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityGovernanceActivationScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityGovernanceActivationScopeImpl: %+v", err)
	}

	return RawIdentityGovernanceActivationScopeImpl{
		identityGovernanceActivationScope: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
