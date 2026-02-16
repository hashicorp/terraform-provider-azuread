package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodTarget = PasskeyAuthenticationMethodTarget{}

type PasskeyAuthenticationMethodTarget struct {

	// Fields inherited from AuthenticationMethodTarget

	// Determines if the user is enforced to register the authentication method.
	IsRegistrationRequired *bool `json:"isRegistrationRequired,omitempty"`

	TargetType *AuthenticationMethodTargetType `json:"targetType,omitempty"`

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

func (s PasskeyAuthenticationMethodTarget) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return BaseAuthenticationMethodTargetImpl{
		IsRegistrationRequired: s.IsRegistrationRequired,
		TargetType:             s.TargetType,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s PasskeyAuthenticationMethodTarget) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PasskeyAuthenticationMethodTarget{}

func (s PasskeyAuthenticationMethodTarget) MarshalJSON() ([]byte, error) {
	type wrapper PasskeyAuthenticationMethodTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PasskeyAuthenticationMethodTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PasskeyAuthenticationMethodTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.passkeyAuthenticationMethodTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PasskeyAuthenticationMethodTarget: %+v", err)
	}

	return encoded, nil
}
