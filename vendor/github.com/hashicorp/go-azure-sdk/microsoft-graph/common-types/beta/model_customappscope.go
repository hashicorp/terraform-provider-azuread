package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppScope = CustomAppScope{}

type CustomAppScope struct {
	CustomAttributes *CustomAppScopeAttributesDictionary `json:"customAttributes,omitempty"`

	// Fields inherited from AppScope

	// Provides the display name of the app-specific resource represented by the app scope. Provided for display purposes
	// since appScopeId is often an immutable, non-human-readable ID. Read only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Describes the type of app-specific resource represented by the app scope. For display purposes, so a user interface
	// can convey to the user the kind of app specific resource represented by the app scope. Read only.
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

func (s CustomAppScope) AppScope() BaseAppScopeImpl {
	return BaseAppScopeImpl{
		DisplayName: s.DisplayName,
		Type:        s.Type,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s CustomAppScope) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomAppScope{}

func (s CustomAppScope) MarshalJSON() ([]byte, error) {
	type wrapper CustomAppScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomAppScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomAppScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customAppScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomAppScope: %+v", err)
	}

	return encoded, nil
}
