package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionAuthenticationConfiguration = AzureAdPopTokenAuthentication{}

type AzureAdPopTokenAuthentication struct {

	// Fields inherited from CustomExtensionAuthenticationConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AzureAdPopTokenAuthentication) CustomExtensionAuthenticationConfiguration() BaseCustomExtensionAuthenticationConfigurationImpl {
	return BaseCustomExtensionAuthenticationConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureAdPopTokenAuthentication{}

func (s AzureAdPopTokenAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper AzureAdPopTokenAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureAdPopTokenAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureAdPopTokenAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureAdPopTokenAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureAdPopTokenAuthentication: %+v", err)
	}

	return encoded, nil
}
