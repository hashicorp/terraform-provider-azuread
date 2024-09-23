package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityProvider{}

type IdentityProvider struct {
	// The client ID for the application. This is the client ID obtained when registering the application with the identity
	// provider. Required. Not nullable.
	ClientId nullable.Type[string] `json:"clientId,omitempty"`

	// The client secret for the application. This is the client secret obtained when registering the application with the
	// identity provider. This is write-only. A read operation will return . Required. Not nullable.
	ClientSecret nullable.Type[string] `json:"clientSecret,omitempty"`

	// The display name of the identity provider. Not nullable.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft,
	// Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s IdentityProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityProvider{}

func (s IdentityProvider) MarshalJSON() ([]byte, error) {
	type wrapper IdentityProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityProvider: %+v", err)
	}

	return encoded, nil
}
