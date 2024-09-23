package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApiAuthenticationConfigurationBase = ClientCertificateAuthentication{}

type ClientCertificateAuthentication struct {
	// The list of certificates uploaded for this API connector.
	CertificateList *[]Pkcs12CertificateInformation `json:"certificateList,omitempty"`

	// Fields inherited from ApiAuthenticationConfigurationBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ClientCertificateAuthentication) ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl {
	return BaseApiAuthenticationConfigurationBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ClientCertificateAuthentication{}

func (s ClientCertificateAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper ClientCertificateAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClientCertificateAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClientCertificateAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.clientCertificateAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClientCertificateAuthentication: %+v", err)
	}

	return encoded, nil
}
