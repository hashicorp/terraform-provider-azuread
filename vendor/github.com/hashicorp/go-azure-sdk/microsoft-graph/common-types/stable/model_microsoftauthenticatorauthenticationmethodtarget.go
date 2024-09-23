package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodTarget = MicrosoftAuthenticatorAuthenticationMethodTarget{}

type MicrosoftAuthenticatorAuthenticationMethodTarget struct {
	AuthenticationMode *MicrosoftAuthenticatorAuthenticationMode `json:"authenticationMode,omitempty"`

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

func (s MicrosoftAuthenticatorAuthenticationMethodTarget) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return BaseAuthenticationMethodTargetImpl{
		IsRegistrationRequired: s.IsRegistrationRequired,
		TargetType:             s.TargetType,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s MicrosoftAuthenticatorAuthenticationMethodTarget) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftAuthenticatorAuthenticationMethodTarget{}

func (s MicrosoftAuthenticatorAuthenticationMethodTarget) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftAuthenticatorAuthenticationMethodTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftAuthenticatorAuthenticationMethodTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftAuthenticatorAuthenticationMethodTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftAuthenticatorAuthenticationMethodTarget: %+v", err)
	}

	return encoded, nil
}
