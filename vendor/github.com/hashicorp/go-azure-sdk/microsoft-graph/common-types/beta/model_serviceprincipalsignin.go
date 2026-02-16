package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignInIdentity = ServicePrincipalSignIn{}

type ServicePrincipalSignIn struct {
	// appId of the service principal that is signing in.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Fields inherited from SignInIdentity

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServicePrincipalSignIn) SignInIdentity() BaseSignInIdentityImpl {
	return BaseSignInIdentityImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServicePrincipalSignIn{}

func (s ServicePrincipalSignIn) MarshalJSON() ([]byte, error) {
	type wrapper ServicePrincipalSignIn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePrincipalSignIn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePrincipalSignIn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.servicePrincipalSignIn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePrincipalSignIn: %+v", err)
	}

	return encoded, nil
}
