package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedProvider interface {
	Entity
	IdentityProviderBase
	SamlOrWsFedProvider() BaseSamlOrWsFedProviderImpl
}

var _ SamlOrWsFedProvider = BaseSamlOrWsFedProviderImpl{}

type BaseSamlOrWsFedProviderImpl struct {
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

func (s BaseSamlOrWsFedProviderImpl) SamlOrWsFedProvider() BaseSamlOrWsFedProviderImpl {
	return s
}

func (s BaseSamlOrWsFedProviderImpl) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s BaseSamlOrWsFedProviderImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SamlOrWsFedProvider = RawSamlOrWsFedProviderImpl{}

// RawSamlOrWsFedProviderImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSamlOrWsFedProviderImpl struct {
	samlOrWsFedProvider BaseSamlOrWsFedProviderImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawSamlOrWsFedProviderImpl) SamlOrWsFedProvider() BaseSamlOrWsFedProviderImpl {
	return s.samlOrWsFedProvider
}

func (s RawSamlOrWsFedProviderImpl) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return s.samlOrWsFedProvider.IdentityProviderBase()
}

func (s RawSamlOrWsFedProviderImpl) Entity() BaseEntityImpl {
	return s.samlOrWsFedProvider.Entity()
}

var _ json.Marshaler = BaseSamlOrWsFedProviderImpl{}

func (s BaseSamlOrWsFedProviderImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSamlOrWsFedProviderImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSamlOrWsFedProviderImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSamlOrWsFedProviderImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.samlOrWsFedProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSamlOrWsFedProviderImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSamlOrWsFedProviderImplementation(input []byte) (SamlOrWsFedProvider, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SamlOrWsFedProvider into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.internalDomainFederation") {
		var out InternalDomainFederation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternalDomainFederation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlOrWsFedExternalDomainFederation") {
		var out SamlOrWsFedExternalDomainFederation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlOrWsFedExternalDomainFederation: %+v", err)
		}
		return out, nil
	}

	var parent BaseSamlOrWsFedProviderImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSamlOrWsFedProviderImpl: %+v", err)
	}

	return RawSamlOrWsFedProviderImpl{
		samlOrWsFedProvider: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
