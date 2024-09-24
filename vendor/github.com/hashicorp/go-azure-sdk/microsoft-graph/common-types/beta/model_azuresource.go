package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystemIdentitySource = AzureSource{}

type AzureSource struct {
	// Azure subscription ID.
	SubscriptionId nullable.Type[string] `json:"subscriptionId,omitempty"`

	// Fields inherited from AuthorizationSystemIdentitySource

	// Type of identity provider. Read-only.
	IdentityProviderType nullable.Type[string] `json:"identityProviderType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AzureSource) AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl {
	return BaseAuthorizationSystemIdentitySourceImpl{
		IdentityProviderType: s.IdentityProviderType,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = AzureSource{}

func (s AzureSource) MarshalJSON() ([]byte, error) {
	type wrapper AzureSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureSource: %+v", err)
	}

	return encoded, nil
}
