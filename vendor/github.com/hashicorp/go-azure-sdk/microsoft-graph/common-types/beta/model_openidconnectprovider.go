package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityProvider = OpenIdConnectProvider{}

type OpenIdConnectProvider struct {
	// After the OIDC provider sends an ID token back to Microsoft Entra ID, Microsoft Entra ID needs to be able to map the
	// claims from the received token to the claims that Microsoft Entra ID recognizes and uses. This complex type captures
	// that mapping. It's a required property.
	ClaimsMapping *ClaimsMapping `json:"claimsMapping,omitempty"`

	// The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of
	// having the user make a selection among the list of available identity providers.
	DomainHint nullable.Type[string] `json:"domainHint,omitempty"`

	// The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider
	// describes a metadata document that contains most of the information required to perform sign-in. This includes
	// information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect
	// metadata document is always located at an endpoint that ends in a well-known/openid-configuration. For the OpenID
	// Connect identity provider you're looking to add, you need to provide the metadata URL. It's a required property and
	// is read only after creation.
	MetadataUrl nullable.Type[string] `json:"metadataUrl,omitempty"`

	ResponseMode *OpenIdConnectResponseMode  `json:"responseMode,omitempty"`
	ResponseType *OpenIdConnectResponseTypes `json:"responseType,omitempty"`

	// Scope defines the information and permissions you're looking to gather from your custom identity provider. OpenID
	// Connect requests must contain the openid scope value in order to receive the ID token from the identity provider.
	// Without the ID token, users aren't able to sign in to Azure AD B2C using the custom identity provider. Other scopes
	// can be appended separated by space. For more information about the scope limitations, see RFC6749 Section 3.3. It's a
	// required property.
	Scope nullable.Type[string] `json:"scope,omitempty"`

	// Fields inherited from IdentityProvider

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

func (s OpenIdConnectProvider) IdentityProvider() BaseIdentityProviderImpl {
	return BaseIdentityProviderImpl{
		ClientId:     s.ClientId,
		ClientSecret: s.ClientSecret,
		Name:         s.Name,
		Type:         s.Type,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s OpenIdConnectProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OpenIdConnectProvider{}

func (s OpenIdConnectProvider) MarshalJSON() ([]byte, error) {
	type wrapper OpenIdConnectProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenIdConnectProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenIdConnectProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openIdConnectProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenIdConnectProvider: %+v", err)
	}

	return encoded, nil
}
