package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityProviderBase = BuiltInIdentityProvider{}

type BuiltInIdentityProvider struct {
	// The identity provider type. For a B2B scenario, possible values: AADSignup, MicrosoftAccount, EmailOTP. Required.
	IdentityProviderType nullable.Type[string] `json:"identityProviderType,omitempty"`

	State *IdentityProviderState `json:"state,omitempty"`

	// Fields inherited from IdentityProviderBase

	// The display name of the identity provider.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BuiltInIdentityProvider) IdentityProviderBase() BaseIdentityProviderBaseImpl {
	return BaseIdentityProviderBaseImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s BuiltInIdentityProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BuiltInIdentityProvider{}

func (s BuiltInIdentityProvider) MarshalJSON() ([]byte, error) {
	type wrapper BuiltInIdentityProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BuiltInIdentityProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BuiltInIdentityProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.builtInIdentityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BuiltInIdentityProvider: %+v", err)
	}

	return encoded, nil
}
