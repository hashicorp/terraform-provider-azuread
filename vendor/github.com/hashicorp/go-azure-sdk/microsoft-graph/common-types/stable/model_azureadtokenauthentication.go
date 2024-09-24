package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomExtensionAuthenticationConfiguration = AzureAdTokenAuthentication{}

type AzureAdTokenAuthentication struct {
	// The appID of the Microsoft Entra application to use to authenticate an app with a custom extension.
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// Fields inherited from CustomExtensionAuthenticationConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AzureAdTokenAuthentication) CustomExtensionAuthenticationConfiguration() BaseCustomExtensionAuthenticationConfigurationImpl {
	return BaseCustomExtensionAuthenticationConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureAdTokenAuthentication{}

func (s AzureAdTokenAuthentication) MarshalJSON() ([]byte, error) {
	type wrapper AzureAdTokenAuthentication
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureAdTokenAuthentication: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureAdTokenAuthentication: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureAdTokenAuthentication"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureAdTokenAuthentication: %+v", err)
	}

	return encoded, nil
}
