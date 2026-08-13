package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInContext interface {
	SignInContext() BaseSignInContextImpl
}

var _ SignInContext = BaseSignInContextImpl{}

type BaseSignInContextImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSignInContextImpl) SignInContext() BaseSignInContextImpl {
	return s
}

var _ SignInContext = RawSignInContextImpl{}

// RawSignInContextImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSignInContextImpl struct {
	signInContext BaseSignInContextImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawSignInContextImpl) SignInContext() BaseSignInContextImpl {
	return s.signInContext
}

func (s RawSignInContextImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSignInContextImplementation(input []byte) (SignInContext, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SignInContext into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.applicationContext") {
		var out ApplicationContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationContext: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.authContext") {
		var out AuthContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuthContext: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userActionContext") {
		var out UserActionContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserActionContext: %+v", err)
		}
		return out, nil
	}

	var parent BaseSignInContextImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSignInContextImpl: %+v", err)
	}

	return RawSignInContextImpl{
		signInContext: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
