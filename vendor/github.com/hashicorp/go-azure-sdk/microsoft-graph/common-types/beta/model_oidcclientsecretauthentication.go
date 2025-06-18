package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OidcClientAuthentication = OidcClientSecretAuthentication{}

type OidcClientSecretAuthentication struct {
	// The client secret obtained from configuring the client application on the external OpenID Connect identity provider.
	// The property includes the client secret and enables the identity provider to use either the clientsecretpost
	// authentication method.
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Fields inherited from OidcClientAuthentication

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OidcClientSecretAuthentication) OidcClientAuthentication() BaseOidcClientAuthenticationImpl {
	return BaseOidcClientAuthenticationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OidcClientSecretAuthentication{}

func (s OidcClientSecretAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper OidcClientSecretAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OidcClientSecretAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OidcClientSecretAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.oidcClientSecretAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OidcClientSecretAuthentication: %+v", err)
	}

	return encoded, nil
}
