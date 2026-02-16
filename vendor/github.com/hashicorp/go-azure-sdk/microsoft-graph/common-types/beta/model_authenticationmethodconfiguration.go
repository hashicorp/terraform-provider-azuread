package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodConfiguration interface {
	Entity
	AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl
}

var _ AuthenticationMethodConfiguration = BaseAuthenticationMethodConfigurationImpl{}

type BaseAuthenticationMethodConfigurationImpl struct {
	// Groups of users that are excluded from a policy.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// The state of the policy. Possible values are: enabled, disabled.
	State *AuthenticationMethodState `json:"state,omitempty"`

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

func (s BaseAuthenticationMethodConfigurationImpl) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return s
}

func (s BaseAuthenticationMethodConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationMethodConfiguration = RawAuthenticationMethodConfigurationImpl{}

// RawAuthenticationMethodConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationMethodConfigurationImpl struct {
	authenticationMethodConfiguration BaseAuthenticationMethodConfigurationImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawAuthenticationMethodConfigurationImpl) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return s.authenticationMethodConfiguration
}

func (s RawAuthenticationMethodConfigurationImpl) Entity() BaseEntityImpl {
	return s.authenticationMethodConfiguration.Entity()
}

var _ json.Marshaler = BaseAuthenticationMethodConfigurationImpl{}

func (s BaseAuthenticationMethodConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationMethodConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationMethodConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationMethodConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationMethodConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationMethodConfigurationImplementation(input []byte) (AuthenticationMethodConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethodConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.emailAuthenticationMethodConfiguration") {
		var out EmailAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalAuthenticationMethodConfiguration") {
		var out ExternalAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2AuthenticationMethodConfiguration") {
		var out Fido2AuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2AuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathAuthenticationMethodConfiguration") {
		var out HardwareOathAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethodConfiguration") {
		var out MicrosoftAuthenticatorAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.qrCodePinAuthenticationMethodConfiguration") {
		var out QrCodePinAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into QrCodePinAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.smsAuthenticationMethodConfiguration") {
		var out SmsAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SmsAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.softwareOathAuthenticationMethodConfiguration") {
		var out SoftwareOathAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SoftwareOathAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.temporaryAccessPassAuthenticationMethodConfiguration") {
		var out TemporaryAccessPassAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TemporaryAccessPassAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.voiceAuthenticationMethodConfiguration") {
		var out VoiceAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VoiceAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateAuthenticationMethodConfiguration") {
		var out X509CertificateAuthenticationMethodConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateAuthenticationMethodConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationMethodConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationMethodConfigurationImpl: %+v", err)
	}

	return RawAuthenticationMethodConfigurationImpl{
		authenticationMethodConfiguration: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
