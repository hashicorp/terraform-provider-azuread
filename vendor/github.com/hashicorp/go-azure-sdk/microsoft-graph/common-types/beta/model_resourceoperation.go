package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ResourceOperation{}

type ResourceOperation struct {
	// Type of action this operation is going to perform. The actionName should be concise and limited to as few words as
	// possible.
	ActionName nullable.Type[string] `json:"actionName,omitempty"`

	// Description of the resource operation. The description is used in mouse-over text for the operation when shown in the
	// Azure Portal.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Determines whether the Permission is validated for Scopes defined per Role Assignment. This property is read-only.
	EnabledForScopeValidation *bool `json:"enabledForScopeValidation,omitempty"`

	// Resource category to which this Operation belongs. This property is read-only.
	Resource nullable.Type[string] `json:"resource,omitempty"`

	// Name of the Resource this operation is performed on.
	ResourceName nullable.Type[string] `json:"resourceName,omitempty"`

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

func (s ResourceOperation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ResourceOperation{}

func (s ResourceOperation) MarshalJSON() ([]byte, error) {
	type wrapper ResourceOperation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceOperation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceOperation: %+v", err)
	}

	delete(decoded, "enabledForScopeValidation")
	delete(decoded, "resource")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.resourceOperation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceOperation: %+v", err)
	}

	return encoded, nil
}
