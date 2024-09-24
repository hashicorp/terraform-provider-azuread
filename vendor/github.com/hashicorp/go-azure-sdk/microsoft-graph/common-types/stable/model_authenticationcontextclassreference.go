package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationContextClassReference{}

type AuthenticationContextClassReference struct {
	// A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be
	// used to provide secondary text to describe the authentication context class reference when building user-facing admin
	// experiences. For example, a selection UX.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name is the friendly name of the authenticationContextClassReference object. This value should be used to
	// identify the authentication context class reference when building user-facing admin experiences. For example, a
	// selection UX.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for
	// use by apps. When it's set to false, it shouldn't be shown in authentication context selection UX, or used to protect
	// app resources. It's shown and available for Conditional Access policy authoring. The default value is false. Supports
	// $filter (eq).
	IsAvailable nullable.Type[bool] `json:"isAvailable,omitempty"`

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

func (s AuthenticationContextClassReference) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationContextClassReference{}

func (s AuthenticationContextClassReference) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationContextClassReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationContextClassReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationContextClassReference: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationContextClassReference"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationContextClassReference: %+v", err)
	}

	return encoded, nil
}
