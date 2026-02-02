package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethod interface {
	Entity
	AuthenticationMethod() BaseAuthenticationMethodImpl
}

var _ AuthenticationMethod = BaseAuthenticationMethodImpl{}

type BaseAuthenticationMethodImpl struct {
	// The date and time the authentication method was registered to the user. Read-only. Optional. This optional value is
	// null if the authentication method doesn't populate it. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

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

func (s BaseAuthenticationMethodImpl) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return s
}

func (s BaseAuthenticationMethodImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationMethod = RawAuthenticationMethodImpl{}

// RawAuthenticationMethodImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationMethodImpl struct {
	authenticationMethod BaseAuthenticationMethodImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawAuthenticationMethodImpl) AuthenticationMethod() BaseAuthenticationMethodImpl {
	return s.authenticationMethod
}

func (s RawAuthenticationMethodImpl) Entity() BaseEntityImpl {
	return s.authenticationMethod.Entity()
}

var _ json.Marshaler = BaseAuthenticationMethodImpl{}

func (s BaseAuthenticationMethodImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationMethodImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationMethodImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationMethodImpl: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationMethod"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationMethodImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationMethodImplementation(input []byte) (AuthenticationMethod, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationMethod into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.emailAuthenticationMethod") {
		var out EmailAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EmailAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2AuthenticationMethod") {
		var out Fido2AuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2AuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.hardwareOathAuthenticationMethod") {
		var out HardwareOathAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HardwareOathAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod") {
		var out MicrosoftAuthenticatorAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftAuthenticatorAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordAuthenticationMethod") {
		var out PasswordAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.passwordlessMicrosoftAuthenticatorAuthenticationMethod") {
		var out PasswordlessMicrosoftAuthenticatorAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PasswordlessMicrosoftAuthenticatorAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.phoneAuthenticationMethod") {
		var out PhoneAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PhoneAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.platformCredentialAuthenticationMethod") {
		var out PlatformCredentialAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlatformCredentialAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.qrCodePinAuthenticationMethod") {
		var out QrCodePinAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into QrCodePinAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.softwareOathAuthenticationMethod") {
		var out SoftwareOathAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SoftwareOathAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.temporaryAccessPassAuthenticationMethod") {
		var out TemporaryAccessPassAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TemporaryAccessPassAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod") {
		var out WindowsHelloForBusinessAuthenticationMethod
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsHelloForBusinessAuthenticationMethod: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationMethodImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationMethodImpl: %+v", err)
	}

	return RawAuthenticationMethodImpl{
		authenticationMethod: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
