package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceActivationScope = IdentityGovernanceActivateGroupScope{}

type IdentityGovernanceActivateGroupScope struct {
	Group *Group `json:"group,omitempty"`

	// Fields inherited from IdentityGovernanceActivationScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceActivateGroupScope) IdentityGovernanceActivationScope() BaseIdentityGovernanceActivationScopeImpl {
	return BaseIdentityGovernanceActivationScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceActivateGroupScope{}

func (s IdentityGovernanceActivateGroupScope) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceActivateGroupScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceActivateGroupScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceActivateGroupScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.activateGroupScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceActivateGroupScope: %+v", err)
	}

	return encoded, nil
}
