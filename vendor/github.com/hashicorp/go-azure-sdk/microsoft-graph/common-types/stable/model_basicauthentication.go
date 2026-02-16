package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApiAuthenticationConfigurationBase = BasicAuthentication{}

type BasicAuthentication struct {
	// The password. It isn't returned in the responses.
	Password nullable.Type[string] `json:"password,omitempty"`

	// The username.
	Username nullable.Type[string] `json:"username,omitempty"`

	// Fields inherited from ApiAuthenticationConfigurationBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BasicAuthentication) ApiAuthenticationConfigurationBase() BaseApiAuthenticationConfigurationBaseImpl {
	return BaseApiAuthenticationConfigurationBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BasicAuthentication{}

func (s BasicAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper BasicAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasicAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasicAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.basicAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasicAuthentication: %+v", err)
	}

	return encoded, nil
}
