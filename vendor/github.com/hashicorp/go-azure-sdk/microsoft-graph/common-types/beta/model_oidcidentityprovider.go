package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityProviderBase = OidcIdentityProvider{}

type OidcIdentityProvider struct {
	ClientAuthentication OidcClientAuthentication `json:"clientAuthentication"`

	// The client ID for the application obtained when registering the application with the identity provider.
	ClientId *string `json:"clientId,omitempty"`

	// After the OIDC provider sends an ID token back to Microsoft Entra External ID, Microsoft Entra External ID needs to
	// be able to map the claims from the received token to the claims that Microsoft Entra ID recognizes and uses. This
	// complex type captures that mapping.
	InboundClaimMapping *OidcInboundClaimMappingOverride `json:"inboundClaimMapping,omitempty"`

	// The issuer URI. Issuer URI is a case-sensitive URL using https scheme contains scheme, host, and optionally, port
	// number and path components and no query or fragment components. Note: Configuring other Microsoft Entra tenants as an
	// external identity provider is currently not supported. As a result, the microsoftonline.com domain in the issuer URI
	// is not accepted.
	Issuer *string `json:"issuer,omitempty"`

	ResponseType *OidcResponseType `json:"responseType,omitempty"`

	// Scope defines the information and permissions you are looking to gather from your custom identity provider.
	Scope *string `json:"scope,omitempty"`

	// The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider
	// describes a metadata document that contains most of the information required to perform sign-in. This includes
	// information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect
	// metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Note: The metadata
	// document should, at minimum, contain the following properties: issuer, authorizationendpoint, tokenendpoint,
	// tokenendpointauthmethodssupported, responsetypessupported, subjecttypessupported and jwks_uri. Visit OpenID Connect
	// Discovery specifications for more details.
	WellKnownEndpoint nullable.Type[string] `json:"wellKnownEndpoint,omitempty"`

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

func (s OidcIdentityProvider) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s OidcIdentityProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OidcIdentityProvider{}

func (s OidcIdentityProvider) MarshalJSON() ([]byte, error) {
	type wrapper OidcIdentityProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OidcIdentityProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OidcIdentityProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.oidcIdentityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OidcIdentityProvider: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OidcIdentityProvider{}

func (s *OidcIdentityProvider) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClientId            *string                          `json:"clientId,omitempty"`
		InboundClaimMapping *OidcInboundClaimMappingOverride `json:"inboundClaimMapping,omitempty"`
		Issuer              *string                          `json:"issuer,omitempty"`
		ResponseType        *OidcResponseType                `json:"responseType,omitempty"`
		Scope               *string                          `json:"scope,omitempty"`
		WellKnownEndpoint   nullable.Type[string]            `json:"wellKnownEndpoint,omitempty"`
		DisplayName         nullable.Type[string]            `json:"displayName,omitempty"`
		Id                  *string                          `json:"id,omitempty"`
		ODataId             *string                          `json:"@odata.id,omitempty"`
		ODataType           *string                          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClientId = decoded.ClientId
	s.InboundClaimMapping = decoded.InboundClaimMapping
	s.Issuer = decoded.Issuer
	s.ResponseType = decoded.ResponseType
	s.Scope = decoded.Scope
	s.WellKnownEndpoint = decoded.WellKnownEndpoint
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OidcIdentityProvider into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["clientAuthentication"]; ok {
		impl, err := UnmarshalOidcClientAuthenticationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ClientAuthentication' for 'OidcIdentityProvider': %+v", err)
		}
		s.ClientAuthentication = impl
	}

	return nil
}
