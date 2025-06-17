package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SamlOrWsFedProvider = InternalDomainFederation{}

type InternalDomainFederation struct {
	// URL of the endpoint used by active clients when authenticating with federated domains set up for single sign-on in
	// Microsoft Entra ID. Corresponds to the ActiveLogOnUri property of the Set-EntraDomainFederationSettings PowerShell
	// cmdlet.
	ActiveSignInUri nullable.Type[string] `json:"activeSignInUri,omitempty"`

	DefaultInteractiveAuthenticationMethod nullable.Type[string] `json:"defaultInteractiveAuthenticationMethod,omitempty"`

	// Determines whether Microsoft Entra ID accepts the MFA performed by the federated IdP when a federated user accesses
	// an application that is governed by a conditional access policy that requires MFA. The possible values are:
	// acceptIfMfaDoneByFederatedIdp, enforceMfaByFederatedIdp, rejectMfaByFederatedIdp, unknownFutureValue. For more
	// information, see federatedIdpMfaBehavior values.
	FederatedIdpMfaBehavior *FederatedIdpMfaBehavior `json:"federatedIdpMfaBehavior,omitempty"`

	// If true, when SAML authentication requests are sent to the federated SAML IdP, Microsoft Entra ID will sign those
	// requests using the OrgID signing key. If false (default), the SAML authentication requests sent to the federated IdP
	// aren't signed.
	IsSignedAuthenticationRequestRequired nullable.Type[bool] `json:"isSignedAuthenticationRequestRequired,omitempty"`

	// Fallback token signing certificate that can also be used to sign tokens, for example when the primary signing
	// certificate expires. Formatted as Base64 encoded strings of the public portion of the federated IdP's token signing
	// certificate. Needs to be compatible with the X509Certificate2 class. Much like the signingCertificate, the
	// nextSigningCertificate property is used if a rollover is required outside of the auto-rollover update, a new
	// federation service is being set up, or if the new token signing certificate isn't present in the federation
	// properties after the federation service certificate has been updated.
	NextSigningCertificate nullable.Type[string] `json:"nextSigningCertificate,omitempty"`

	OpenIdConnectDiscoveryEndpoint nullable.Type[string] `json:"openIdConnectDiscoveryEndpoint,omitempty"`
	PasswordChangeUri              nullable.Type[string] `json:"passwordChangeUri,omitempty"`

	// URI that clients are redirected to for resetting their password.
	PasswordResetUri nullable.Type[string] `json:"passwordResetUri,omitempty"`

	// Sets the preferred behavior for the sign-in prompt. The possible values are: translateToFreshPasswordAuthentication,
	// nativeSupport, disabled, unknownFutureValue.
	PromptLoginBehavior *PromptLoginBehavior `json:"promptLoginBehavior,omitempty"`

	// URI that clients are redirected to when they sign out of Microsoft Entra services. Corresponds to the LogOffUri
	// property of the Set-EntraDomainFederationSettings PowerShell cmdlet.
	SignOutUri nullable.Type[string] `json:"signOutUri,omitempty"`

	// Provides status and timestamp of the last update of the signing certificate.
	SigningCertificateUpdateStatus *SigningCertificateUpdateStatus `json:"signingCertificateUpdateStatus,omitempty"`

	// Fields inherited from SamlOrWsFedProvider

	// Issuer URI of the federation server.
	IssuerUri nullable.Type[string] `json:"issuerUri,omitempty"`

	// URI of the metadata exchange endpoint used for authentication from rich client applications.
	MetadataExchangeUri nullable.Type[string] `json:"metadataExchangeUri,omitempty"`

	// URI that web-based clients are directed to when signing in to Microsoft Entra services.
	PassiveSignInUri nullable.Type[string] `json:"passiveSignInUri,omitempty"`

	// Preferred authentication protocol. Supported values include saml or wsfed.
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

func (s InternalDomainFederation) SamlOrWsFedProvider() BaseSamlOrWsFedProviderImpl {
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

func (s InternalDomainFederation) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s InternalDomainFederation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InternalDomainFederation{}

func (s InternalDomainFederation) MarshalJSON() ([]byte, error) {
	type wrapper InternalDomainFederation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InternalDomainFederation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InternalDomainFederation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.internalDomainFederation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InternalDomainFederation: %+v", err)
	}

	return encoded, nil
}
