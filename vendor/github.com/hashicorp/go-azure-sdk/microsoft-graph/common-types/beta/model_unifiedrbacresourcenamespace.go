package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedRbacResourceNamespace{}

type UnifiedRbacResourceNamespace struct {
	// Name of the resource namespace. Typically, the same name as the id property, such as microsoft.aad.b2c. Required.
	// Supports $filter (eq, startsWith).
	Name string `json:"name"`

	// Operations that an authorized principal is allowed to perform.
	ResourceActions *[]UnifiedRbacResourceAction `json:"resourceActions,omitempty"`

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

func (s UnifiedRbacResourceNamespace) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedRbacResourceNamespace{}

func (s UnifiedRbacResourceNamespace) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedRbacResourceNamespace
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedRbacResourceNamespace: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedRbacResourceNamespace: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedRbacResourceNamespace"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedRbacResourceNamespace: %+v", err)
	}

	return encoded, nil
}
