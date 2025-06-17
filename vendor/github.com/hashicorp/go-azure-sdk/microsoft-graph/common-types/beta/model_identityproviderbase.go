package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderBase interface {
	Entity
	IdentityProviderBase() BaseIdentityProviderBaseImpl
}

var _ IdentityProviderBase = BaseIdentityProviderBaseImpl{}

type BaseIdentityProviderBaseImpl struct {
	// The display name of the identity provider.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

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

func (s BaseIdentityProviderBaseImpl) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return s
}

func (s BaseIdentityProviderBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IdentityProviderBase = RawIdentityProviderBaseImpl{}

// RawIdentityProviderBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityProviderBaseImpl struct {
	identityProviderBase BaseIdentityProviderBaseImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawIdentityProviderBaseImpl) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return s.identityProviderBase
}

func (s RawIdentityProviderBaseImpl) Entity() BaseEntityImpl {
	return s.identityProviderBase.Entity()
}

var _ json.Marshaler = BaseIdentityProviderBaseImpl{}

func (s BaseIdentityProviderBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIdentityProviderBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIdentityProviderBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIdentityProviderBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityProviderBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIdentityProviderBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIdentityProviderBaseImplementation(input []byte) (IdentityProviderBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityProviderBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appleManagedIdentityProvider") {
		var out AppleManagedIdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleManagedIdentityProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.builtInIdentityProvider") {
		var out BuiltInIdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BuiltInIdentityProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oidcIdentityProvider") {
		var out OidcIdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OidcIdentityProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.openIdConnectIdentityProvider") {
		var out OpenIdConnectIdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenIdConnectIdentityProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlOrWsFedProvider") {
		var out SamlOrWsFedProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlOrWsFedProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.socialIdentityProvider") {
		var out SocialIdentityProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SocialIdentityProvider: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityProviderBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityProviderBaseImpl: %+v", err)
	}

	return RawIdentityProviderBaseImpl{
		identityProviderBase: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
