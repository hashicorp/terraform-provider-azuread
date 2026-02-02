package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInIdentity interface {
	SignInIdentity() BaseSignInIdentityImpl
}

var _ SignInIdentity = BaseSignInIdentityImpl{}

type BaseSignInIdentityImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSignInIdentityImpl) SignInIdentity() BaseSignInIdentityImpl {
	return s
}

var _ SignInIdentity = RawSignInIdentityImpl{}

// RawSignInIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSignInIdentityImpl struct {
	signInIdentity BaseSignInIdentityImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawSignInIdentityImpl) SignInIdentity() BaseSignInIdentityImpl {
	return s.signInIdentity
}

func UnmarshalSignInIdentityImplementation(input []byte) (SignInIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SignInIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalSignIn") {
		var out ServicePrincipalSignIn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalSignIn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSignIn") {
		var out UserSignIn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSignIn: %+v", err)
		}
		return out, nil
	}

	var parent BaseSignInIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSignInIdentityImpl: %+v", err)
	}

	return RawSignInIdentityImpl{
		signInIdentity: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
