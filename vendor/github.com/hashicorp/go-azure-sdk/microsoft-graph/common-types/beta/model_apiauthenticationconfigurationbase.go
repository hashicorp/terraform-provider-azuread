package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiAuthenticationConfigurationBase interface {
	ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl
}

var _ ApiAuthenticationConfigurationBase = BaseApiAuthenticationConfigurationBaseImpl{}

type BaseApiAuthenticationConfigurationBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseApiAuthenticationConfigurationBaseImpl) ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl {
	return s
}

var _ ApiAuthenticationConfigurationBase = RawApiAuthenticationConfigurationBaseImpl{}

// RawApiAuthenticationConfigurationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawApiAuthenticationConfigurationBaseImpl struct {
	apiAuthenticationConfigurationBase BaseApiAuthenticationConfigurationBaseImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawApiAuthenticationConfigurationBaseImpl) ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl {
	return s.apiAuthenticationConfigurationBase
}

func (s RawApiAuthenticationConfigurationBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalApiAuthenticationConfigurationBaseImplementation(input []byte) (ApiAuthenticationConfigurationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApiAuthenticationConfigurationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.basicAuthentication") {
		var out BasicAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BasicAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.clientCertificateAuthentication") {
		var out ClientCertificateAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClientCertificateAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.pkcs12Certificate") {
		var out Pkcs12Certificate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Pkcs12Certificate: %+v", err)
		}
		return out, nil
	}

	var parent BaseApiAuthenticationConfigurationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApiAuthenticationConfigurationBaseImpl: %+v", err)
	}

	return RawApiAuthenticationConfigurationBaseImpl{
		apiAuthenticationConfigurationBase: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
