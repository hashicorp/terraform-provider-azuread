package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApiAuthenticationConfigurationBase = Pkcs12Certificate{}

type Pkcs12Certificate struct {
	// The password for the pfx file. Required. If no password is used, you must still provide a value of ''.
	Password nullable.Type[string] `json:"password,omitempty"`

	// Represents the pfx content that is sent. The value should be a base-64 encoded version of the actual certificate
	// content. Required.
	Pkcs12Value nullable.Type[string] `json:"pkcs12Value,omitempty"`

	// Fields inherited from ApiAuthenticationConfigurationBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Pkcs12Certificate) ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl {
	return BaseApiAuthenticationConfigurationBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Pkcs12Certificate{}

func (s Pkcs12Certificate) MarshalJSON() ([]byte, error) {
	type wrapper Pkcs12Certificate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Pkcs12Certificate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Pkcs12Certificate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.pkcs12Certificate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Pkcs12Certificate: %+v", err)
	}

	return encoded, nil
}
