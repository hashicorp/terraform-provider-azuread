package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Authentication{}

type Authentication struct {
	// The email address registered to a user for authentication.
	EmailMethods *[]EmailAuthenticationMethod `json:"emailMethods,omitempty"`

	// Represents the FIDO2 security keys registered to a user for authentication.
	Fido2Methods *[]Fido2AuthenticationMethod `json:"fido2Methods,omitempty"`

	// Represents all authentication methods registered to a user.
	Methods *[]AuthenticationMethod `json:"methods,omitempty"`

	// The details of the Microsoft Authenticator app registered to a user for authentication.
	MicrosoftAuthenticatorMethods *[]MicrosoftAuthenticatorAuthenticationMethod `json:"microsoftAuthenticatorMethods,omitempty"`

	// Represents the status of a long-running operation, such as a password reset operation.
	Operations *[]LongRunningOperation `json:"operations,omitempty"`

	// Represents the password registered to a user for authentication. For security, the password itself is never returned
	// in the object, but action can be taken to reset a password.
	PasswordMethods *[]PasswordAuthenticationMethod `json:"passwordMethods,omitempty"`

	// The phone numbers registered to a user for authentication.
	PhoneMethods *[]PhoneAuthenticationMethod `json:"phoneMethods,omitempty"`

	// Represents a platform credential instance registered to a user on Mac OS.
	PlatformCredentialMethods *[]PlatformCredentialAuthenticationMethod `json:"platformCredentialMethods,omitempty"`

	// The software OATH time-based one-time password (TOTP) applications registered to a user for authentication.
	SoftwareOathMethods *[]SoftwareOathAuthenticationMethod `json:"softwareOathMethods,omitempty"`

	// Represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
	TemporaryAccessPassMethods *[]TemporaryAccessPassAuthenticationMethod `json:"temporaryAccessPassMethods,omitempty"`

	// Represents the Windows Hello for Business authentication method registered to a user for authentication.
	WindowsHelloForBusinessMethods *[]WindowsHelloForBusinessAuthenticationMethod `json:"windowsHelloForBusinessMethods,omitempty"`

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

func (s Authentication) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Authentication{}

func (s Authentication) MarshalJSON() ([]byte, error) {
	type wrapper Authentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Authentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Authentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Authentication: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Authentication{}

func (s *Authentication) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EmailMethods                   *[]EmailAuthenticationMethod                   `json:"emailMethods,omitempty"`
		Fido2Methods                   *[]Fido2AuthenticationMethod                   `json:"fido2Methods,omitempty"`
		MicrosoftAuthenticatorMethods  *[]MicrosoftAuthenticatorAuthenticationMethod  `json:"microsoftAuthenticatorMethods,omitempty"`
		PasswordMethods                *[]PasswordAuthenticationMethod                `json:"passwordMethods,omitempty"`
		PhoneMethods                   *[]PhoneAuthenticationMethod                   `json:"phoneMethods,omitempty"`
		PlatformCredentialMethods      *[]PlatformCredentialAuthenticationMethod      `json:"platformCredentialMethods,omitempty"`
		SoftwareOathMethods            *[]SoftwareOathAuthenticationMethod            `json:"softwareOathMethods,omitempty"`
		TemporaryAccessPassMethods     *[]TemporaryAccessPassAuthenticationMethod     `json:"temporaryAccessPassMethods,omitempty"`
		WindowsHelloForBusinessMethods *[]WindowsHelloForBusinessAuthenticationMethod `json:"windowsHelloForBusinessMethods,omitempty"`
		Id                             *string                                        `json:"id,omitempty"`
		ODataId                        *string                                        `json:"@odata.id,omitempty"`
		ODataType                      *string                                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EmailMethods = decoded.EmailMethods
	s.Fido2Methods = decoded.Fido2Methods
	s.MicrosoftAuthenticatorMethods = decoded.MicrosoftAuthenticatorMethods
	s.PasswordMethods = decoded.PasswordMethods
	s.PhoneMethods = decoded.PhoneMethods
	s.PlatformCredentialMethods = decoded.PlatformCredentialMethods
	s.SoftwareOathMethods = decoded.SoftwareOathMethods
	s.TemporaryAccessPassMethods = decoded.TemporaryAccessPassMethods
	s.WindowsHelloForBusinessMethods = decoded.WindowsHelloForBusinessMethods
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Authentication into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["methods"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Methods into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethod, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Methods' for 'Authentication': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Methods = &output
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]LongRunningOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLongRunningOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'Authentication': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
