package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodTarget interface {
	Entity
	AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl
}

var _ AuthenticationMethodTarget = BaseAuthenticationMethodTargetImpl{}

type BaseAuthenticationMethodTargetImpl struct {
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

func (s BaseAuthenticationMethodTargetImpl) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return s
}

func (s BaseAuthenticationMethodTargetImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationMethodTarget = RawAuthenticationMethodTargetImpl{}

// RawAuthenticationMethodTargetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationMethodTargetImpl struct {
	authenticationMethodTarget BaseAuthenticationMethodTargetImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawAuthenticationMethodTargetImpl) AuthenticationMethodTarget() BaseAuthenticationMethodTargetImpl {
	return s.authenticationMethodTarget
}

func (s RawAuthenticationMethodTargetImpl) Entity() BaseEntityImpl {
	return s.authenticationMethodTarget.Entity()
}

var _ json.Marshaler = BaseAuthenticationMethodTargetImpl{}

func (s BaseAuthenticationMethodTargetImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationMethodTargetImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationMethodTargetImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationMethodTargetImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationMethodTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationMethodTargetImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationMethodTargetImplementation(input []byte) (AuthenticationMethodTarget, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethodTarget into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodTarget") {
		var out MicrosoftAuthenticatorAuthenticationMethodTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethodTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passkeyAuthenticationMethodTarget") {
		var out PasskeyAuthenticationMethodTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasskeyAuthenticationMethodTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.smsAuthenticationMethodTarget") {
		var out SmsAuthenticationMethodTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SmsAuthenticationMethodTarget: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.voiceAuthenticationMethodTarget") {
		var out VoiceAuthenticationMethodTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VoiceAuthenticationMethodTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationMethodTargetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationMethodTargetImpl: %+v", err)
	}

	return RawAuthenticationMethodTargetImpl{
		authenticationMethodTarget: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
