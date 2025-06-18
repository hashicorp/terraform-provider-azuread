package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OidcClientAuthentication interface {
	OidcClientAuthentication() BaseOidcClientAuthenticationImpl
}

var _ OidcClientAuthentication = BaseOidcClientAuthenticationImpl{}

type BaseOidcClientAuthenticationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOidcClientAuthenticationImpl) OidcClientAuthentication() BaseOidcClientAuthenticationImpl {
	return s
}

var _ OidcClientAuthentication = RawOidcClientAuthenticationImpl{}

// RawOidcClientAuthenticationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOidcClientAuthenticationImpl struct {
	oidcClientAuthentication BaseOidcClientAuthenticationImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawOidcClientAuthenticationImpl) OidcClientAuthentication() BaseOidcClientAuthenticationImpl {
	return s.oidcClientAuthentication
}

func UnmarshalOidcClientAuthenticationImplementation(input []byte) (OidcClientAuthentication, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OidcClientAuthentication into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.oidcClientSecretAuthentication") {
		var out OidcClientSecretAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OidcClientSecretAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oidcPrivateJwtKeyClientAuthentication") {
		var out OidcPrivateJwtKeyClientAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OidcPrivateJwtKeyClientAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseOidcClientAuthenticationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOidcClientAuthenticationImpl: %+v", err)
	}

	return RawOidcClientAuthenticationImpl{
		oidcClientAuthentication: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
