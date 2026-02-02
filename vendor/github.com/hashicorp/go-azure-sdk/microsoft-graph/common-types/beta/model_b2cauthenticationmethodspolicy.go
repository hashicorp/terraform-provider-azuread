package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = B2cAuthenticationMethodsPolicy{}

type B2cAuthenticationMethodsPolicy struct {
	// The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
	IsEmailPasswordAuthenticationEnabled *bool `json:"isEmailPasswordAuthenticationEnabled,omitempty"`

	// The tenant admin can configure local accounts using phone number if the phone number and one-time password
	// authentication method is enabled.
	IsPhoneOneTimePasswordAuthenticationEnabled *bool `json:"isPhoneOneTimePasswordAuthenticationEnabled,omitempty"`

	// The tenant admin can configure local accounts using username if the username and password authentication method is
	// enabled.
	IsUserNameAuthenticationEnabled *bool `json:"isUserNameAuthenticationEnabled,omitempty"`

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

func (s B2cAuthenticationMethodsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = B2cAuthenticationMethodsPolicy{}

func (s B2cAuthenticationMethodsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper B2cAuthenticationMethodsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling B2cAuthenticationMethodsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling B2cAuthenticationMethodsPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.b2cAuthenticationMethodsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling B2cAuthenticationMethodsPolicy: %+v", err)
	}

	return encoded, nil
}
