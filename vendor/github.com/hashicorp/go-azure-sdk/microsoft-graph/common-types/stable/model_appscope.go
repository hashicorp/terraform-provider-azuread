package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AppScope{}

type AppScope struct {
	// Provides the display name of the app-specific resource represented by the app scope. Provided for display purposes
	// since appScopeId is often an immutable, non-human-readable id. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Describes the type of app-specific resource represented by the app scope and is provided for display purposes, so a
	// user interface can convey to the user the kind of app specific resource represented by the app scope. Read-only.
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

func (s AppScope) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppScope{}

func (s AppScope) MarshalJSON() ([]byte, error) {
	type wrapper AppScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppScope: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "type")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppScope: %+v", err)
	}

	return encoded, nil
}
