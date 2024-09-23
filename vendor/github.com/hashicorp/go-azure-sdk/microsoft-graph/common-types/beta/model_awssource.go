package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthorizationSystemIdentitySource = AwsSource{}

type AwsSource struct {
	// AWS account ID.
	AccountId nullable.Type[string] `json:"accountId,omitempty"`

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

func (s AwsSource) AuthorizationSystemIdentitySource() BaseAuthorizationSystemIdentitySourceImpl {
	return BaseAuthorizationSystemIdentitySourceImpl{
		IdentityProviderType: s.IdentityProviderType,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

var _ json.Marshaler = AwsSource{}

func (s AwsSource) MarshalJSON() ([]byte, error) {
	type wrapper AwsSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsSource: %+v", err)
	}

	return encoded, nil
}
