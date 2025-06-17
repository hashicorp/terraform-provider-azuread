package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceActivationScope = IdentityGovernanceActivateUserScope{}

type IdentityGovernanceActivateUserScope struct {
	Users *[]User `json:"users,omitempty"`

	// Fields inherited from IdentityGovernanceActivationScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceActivateUserScope) IdentityGovernanceActivationScope() BaseIdentityGovernanceActivationScopeImpl {
	return BaseIdentityGovernanceActivationScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceActivateUserScope{}

func (s IdentityGovernanceActivateUserScope) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceActivateUserScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceActivateUserScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceActivateUserScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.activateUserScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceActivateUserScope: %+v", err)
	}

	return encoded, nil
}
