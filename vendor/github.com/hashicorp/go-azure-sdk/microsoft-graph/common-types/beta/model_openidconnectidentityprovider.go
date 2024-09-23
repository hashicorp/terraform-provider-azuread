package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityProviderBase = OpenIdConnectIdentityProvider{}

type OpenIdConnectIdentityProvider struct {
	// After the OIDC provider sends an ID token back to Microsoft Entra ID, Microsoft Entra ID needs to be able to map the
	// claims from the received token to the claims that Microsoft Entra ID recognizes and uses. This complex type captures
	// that mapping. Required.
	ClaimsMapping ClaimsMapping `json:"claimsMapping"`

	// The client identifier for the application obtained when registering the application with the identity provider.
	// Required.
	ClientId nullable.Type[string] `json:"clientId,omitempty"`

	// The client secret for the application obtained when registering the application with the identity provider. The
	// clientSecret has a dependency on responseType. When responseType is code, a secret is required for the auth code
	// exchange.When responseType is idtoken the secret is not required because there is no code exchange. The idtoken is
	// returned directly from the authorization response. This is write-only. A read operation returns .
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of
	// having the user make a selection among the list of available identity providers.
	DomainHint nullable.Type[string] `json:"domainHint,omitempty"`

	// The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider
	// describes a metadata document that contains most of the information required to perform sign-in. This includes
	// information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect
	// metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Provide the
	// metadata URL for the OpenID Connect identity provider you add. Read-only. Required.
	MetadataUrl nullable.Type[string] `json:"metadataUrl,omitempty"`

	ResponseMode *OpenIdConnectResponseMode  `json:"responseMode,omitempty"`
	ResponseType *OpenIdConnectResponseTypes `json:"responseType,omitempty"`

	// Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID
	// Connect requests must contain the openid scope value in order to receive the ID token from the identity provider.
	// Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes
	// can be appended, separated by a space. For more details about the scope limitations, see RFC6749 Section 3.3.
	// Required.
	Scope nullable.Type[string] `json:"scope,omitempty"`

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

func (s OpenIdConnectIdentityProvider) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s OpenIdConnectIdentityProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OpenIdConnectIdentityProvider{}

func (s OpenIdConnectIdentityProvider) MarshalJSON() ([]byte, error) {
	type wrapper OpenIdConnectIdentityProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenIdConnectIdentityProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenIdConnectIdentityProvider: %+v", err)
	}

	delete(decoded, "metadataUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openIdConnectIdentityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenIdConnectIdentityProvider: %+v", err)
	}

	return encoded, nil
}
