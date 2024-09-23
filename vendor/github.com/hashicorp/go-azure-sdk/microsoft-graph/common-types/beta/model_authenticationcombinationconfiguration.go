package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationCombinationConfiguration interface {
	Entity
	AuthenticationCombinationConfiguration() BaseAuthenticationCombinationConfigurationImpl
}

var _ AuthenticationCombinationConfiguration = BaseAuthenticationCombinationConfigurationImpl{}

type BaseAuthenticationCombinationConfigurationImpl struct {
	// Which authentication method combinations this configuration applies to. Must be an allowedCombinations object defined
	// for the authenticationStrengthPolicy. For fido2combinationConfigurations use 'fido2', for
	// x509certificatecombinationconfiguration use 'x509CertificateSingleFactor' or 'x509CertificateMultiFactor'.
	AppliesToCombinations *[]AuthenticationMethodModes `json:"appliesToCombinations,omitempty"`

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

func (s BaseAuthenticationCombinationConfigurationImpl) AuthenticationCombinationConfiguration() BaseAuthenticationCombinationConfigurationImpl {
	return s
}

func (s BaseAuthenticationCombinationConfigurationImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationCombinationConfiguration = RawAuthenticationCombinationConfigurationImpl{}

// RawAuthenticationCombinationConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationCombinationConfigurationImpl struct {
	authenticationCombinationConfiguration BaseAuthenticationCombinationConfigurationImpl
	Type                                   string
	Values                                 map[string]interface{}
}

func (s RawAuthenticationCombinationConfigurationImpl) AuthenticationCombinationConfiguration() BaseAuthenticationCombinationConfigurationImpl {
	return s.authenticationCombinationConfiguration
}

func (s RawAuthenticationCombinationConfigurationImpl) Entity() BaseEntityImpl {
	return s.authenticationCombinationConfiguration.Entity()
}

var _ json.Marshaler = BaseAuthenticationCombinationConfigurationImpl{}

func (s BaseAuthenticationCombinationConfigurationImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationCombinationConfigurationImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationCombinationConfigurationImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationCombinationConfigurationImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationCombinationConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationCombinationConfigurationImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationCombinationConfigurationImplementation(input []byte) (AuthenticationCombinationConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationCombinationConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.fido2CombinationConfiguration") {
		var out Fido2CombinationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Fido2CombinationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.x509CertificateCombinationConfiguration") {
		var out X509CertificateCombinationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into X509CertificateCombinationConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationCombinationConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationCombinationConfigurationImpl: %+v", err)
	}

	return RawAuthenticationCombinationConfigurationImpl{
		authenticationCombinationConfiguration: parent,
		Type:                                   value,
		Values:                                 temp,
	}, nil

}
