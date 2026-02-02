package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OidcClientAuthentication = OidcPrivateJwtKeyClientAuthentication{}

type OidcPrivateJwtKeyClientAuthentication struct {

	// Fields inherited from OidcClientAuthentication

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OidcPrivateJwtKeyClientAuthentication) OidcClientAuthentication() BaseOidcClientAuthenticationImpl {
	return BaseOidcClientAuthenticationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OidcPrivateJwtKeyClientAuthentication{}

func (s OidcPrivateJwtKeyClientAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper OidcPrivateJwtKeyClientAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OidcPrivateJwtKeyClientAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OidcPrivateJwtKeyClientAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.oidcPrivateJwtKeyClientAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OidcPrivateJwtKeyClientAuthentication: %+v", err)
	}

	return encoded, nil
}
