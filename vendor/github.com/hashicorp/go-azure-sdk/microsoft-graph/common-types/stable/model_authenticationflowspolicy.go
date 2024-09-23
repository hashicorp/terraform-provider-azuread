package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationFlowsPolicy{}

type AuthenticationFlowsPolicy struct {
	// Inherited property. A description of the policy. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Inherited property. The human-readable name of the policy. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is
	// enabled or disabled. Optional. Read-only.
	SelfServiceSignUp *SelfServiceSignUpAuthenticationFlowConfiguration `json:"selfServiceSignUp,omitempty"`

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

func (s AuthenticationFlowsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationFlowsPolicy{}

func (s AuthenticationFlowsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationFlowsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationFlowsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationFlowsPolicy: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "selfServiceSignUp")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationFlowsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationFlowsPolicy: %+v", err)
	}

	return encoded, nil
}
