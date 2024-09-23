package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityProviderBase = SocialIdentityProvider{}

type SocialIdentityProvider struct {
	// The identifier for the client application obtained when registering the application with the identity provider.
	// Required.
	ClientId nullable.Type[string] `json:"clientId,omitempty"`

	// The client secret for the application that is obtained when the application is registered with the identity provider.
	// This is write-only. A read operation returns . Required.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// For a B2B scenario, possible values: Google, Facebook. For a B2C scenario, possible values: Microsoft, Google,
	// Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat. Required.
	IdentityProviderType nullable.Type[string] `json:"identityProviderType,omitempty"`

	// Fields inherited from IdentityProviderBase

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

func (s SocialIdentityProvider) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s SocialIdentityProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SocialIdentityProvider{}

func (s SocialIdentityProvider) MarshalJSON() ([]byte, error) {
	type wrapper SocialIdentityProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SocialIdentityProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SocialIdentityProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.socialIdentityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SocialIdentityProvider: %+v", err)
	}

	return encoded, nil
}
