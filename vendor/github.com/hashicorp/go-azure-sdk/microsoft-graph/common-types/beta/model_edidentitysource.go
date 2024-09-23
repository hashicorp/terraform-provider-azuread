package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PermissionsDefinitionIdentitySource = EdIdentitySource{}

type EdIdentitySource struct {

	// Fields inherited from PermissionsDefinitionIdentitySource

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EdIdentitySource) PermissionsDefinitionIdentitySource() BasePermissionsDefinitionIdentitySourceImpl {
	return BasePermissionsDefinitionIdentitySourceImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdIdentitySource{}

func (s EdIdentitySource) MarshalJSON() ([]byte, error) {
	type wrapper EdIdentitySource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdIdentitySource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdIdentitySource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.edIdentitySource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdIdentitySource: %+v", err)
	}

	return encoded, nil
}
