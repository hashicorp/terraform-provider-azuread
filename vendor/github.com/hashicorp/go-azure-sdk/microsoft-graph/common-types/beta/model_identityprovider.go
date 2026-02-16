package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProvider interface {
	Entity
	IdentityProvider() BaseIdentityProviderImpl
}

var _ IdentityProvider = BaseIdentityProviderImpl{}

type BaseIdentityProviderImpl struct {
	// The client ID for the application obtained when registering the application with the identity provider. This is a
	// required field. Required. Not nullable.
	ClientId nullable.Type[string] `json:"clientId,omitempty"`

	// The client secret for the application obtained when registering the application with the identity provider. This is
	// write-only. A read operation returns . This is a required field. Required. Not nullable.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The display name of the identity provider. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft,
	// Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo,QQ, WeChat, OpenIDConnect. Not nullable.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s BaseIdentityProviderImpl) IdentityProvider() BaseIdentityProviderImpl {
	return s
}

func (s BaseIdentityProviderImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IdentityProvider = RawIdentityProviderImpl{}

// RawIdentityProviderImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityProviderImpl struct {
	identityProvider BaseIdentityProviderImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawIdentityProviderImpl) IdentityProvider() BaseIdentityProviderImpl {
	return s.identityProvider
}

func (s RawIdentityProviderImpl) Entity() BaseEntityImpl {
	return s.identityProvider.Entity()
}

var _ json.Marshaler = BaseIdentityProviderImpl{}

func (s BaseIdentityProviderImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIdentityProviderImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIdentityProviderImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIdentityProviderImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIdentityProviderImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIdentityProviderImplementation(input []byte) (IdentityProvider, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityProvider into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.openIdConnectProvider") {
		var out OpenIdConnectProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenIdConnectProvider: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityProviderImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityProviderImpl: %+v", err)
	}

	return RawIdentityProviderImpl{
		identityProvider: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
