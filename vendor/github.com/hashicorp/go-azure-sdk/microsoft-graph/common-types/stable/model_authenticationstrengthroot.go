package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationStrengthRoot{}

type AuthenticationStrengthRoot struct {
	// Names and descriptions of all valid authentication method modes in the system.
	AuthenticationMethodModes *[]AuthenticationMethodModeDetail `json:"authenticationMethodModes,omitempty"`

	Combinations *[]AuthenticationMethodModes `json:"combinations,omitempty"`

	// A collection of authentication strength policies that exist for this tenant, including both built-in and custom
	// policies.
	Policies *[]AuthenticationStrengthPolicy `json:"policies,omitempty"`

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

func (s AuthenticationStrengthRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationStrengthRoot{}

func (s AuthenticationStrengthRoot) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationStrengthRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationStrengthRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationStrengthRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationStrengthRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationStrengthRoot: %+v", err)
	}

	return encoded, nil
}
