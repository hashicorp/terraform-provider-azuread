package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnAuthenticationMethodLoadStartHandler = OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp{}

type OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp struct {
	IdentityProviders *[]IdentityProviderBase `json:"identityProviders,omitempty"`

	// Fields inherited from OnAuthenticationMethodLoadStartHandler

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp) OnAuthenticationMethodLoadStartHandler() BaseOnAuthenticationMethodLoadStartHandlerImpl {
	return BaseOnAuthenticationMethodLoadStartHandlerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp{}

func (s OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp) MarshalJSON() ([]byte, error) {
	type wrapper OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp{}

func (s *OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identityProviders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IdentityProviders into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentityProviderBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityProviderBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IdentityProviders' for 'OnAuthenticationMethodLoadStartExternalUsersSelfServiceSignUp': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IdentityProviders = &output
	}

	return nil
}
