package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SamlOrWsFedProvider = SamlOrWsFedExternalDomainFederation{}

type SamlOrWsFedExternalDomainFederation struct {
	// Collection of domain names of the external organizations that the tenant is federating with. Supports $filter (eq).
	Domains *[]ExternalDomainName `json:"domains,omitempty"`

	// Fields inherited from SamlOrWsFedProvider

	// Issuer URI of the federation server.
	IssuerUri nullable.Type[string] `json:"issuerUri,omitempty"`

	// URI of the metadata exchange endpoint used for authentication from rich client applications.
	MetadataExchangeUri nullable.Type[string] `json:"metadataExchangeUri,omitempty"`

	// URI that web-based clients are directed to when signing in to Microsoft Entra services.
	PassiveSignInUri nullable.Type[string] `json:"passiveSignInUri,omitempty"`

	// Preferred authentication protocol. The possible values are: wsFed, saml, unknownFutureValue.
	PreferredAuthenticationProtocol *AuthenticationProtocol `json:"preferredAuthenticationProtocol,omitempty"`

	// Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a
	// Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible
	// with the X509Certificate2 class. This property is used in the following scenarios: if a rollover is required outside
	// of the autorollover update a new federation service is being set up if the new token signing certificate isn't
	// present in the federation properties after the federation service certificate has been updated. Microsoft Entra ID
	// updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the
	// federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available,
	// Microsoft Entra ID monitors the metadata daily and will update the federation settings for the domain when a new
	// certificate is available.
	SigningCertificate nullable.Type[string] `json:"signingCertificate,omitempty"`

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

func (s SamlOrWsFedExternalDomainFederation) SamlOrWsFedProvider() BaseSamlOrWsFedProviderImpl {
	return BaseSamlOrWsFedProviderImpl{
		IssuerUri:                       s.IssuerUri,
		MetadataExchangeUri:             s.MetadataExchangeUri,
		PassiveSignInUri:                s.PassiveSignInUri,
		PreferredAuthenticationProtocol: s.PreferredAuthenticationProtocol,
		SigningCertificate:              s.SigningCertificate,
		DisplayName:                     s.DisplayName,
		Id:                              s.Id,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
	}
}

func (s SamlOrWsFedExternalDomainFederation) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s SamlOrWsFedExternalDomainFederation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SamlOrWsFedExternalDomainFederation{}

func (s SamlOrWsFedExternalDomainFederation) MarshalJSON() ([]byte, error) {
	type wrapper SamlOrWsFedExternalDomainFederation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SamlOrWsFedExternalDomainFederation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SamlOrWsFedExternalDomainFederation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.samlOrWsFedExternalDomainFederation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SamlOrWsFedExternalDomainFederation: %+v", err)
	}

	return encoded, nil
}
