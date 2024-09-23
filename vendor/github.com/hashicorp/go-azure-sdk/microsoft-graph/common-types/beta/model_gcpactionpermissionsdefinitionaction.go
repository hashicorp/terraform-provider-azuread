package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GcpPermissionsDefinitionAction = GcpActionPermissionsDefinitionAction{}

type GcpActionPermissionsDefinitionAction struct {
	// List of actions.
	Actions *[]string `json:"actions,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s GcpActionPermissionsDefinitionAction) GcpPermissionsDefinitionAction() BaseGcpPermissionsDefinitionActionImpl {
	return BaseGcpPermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s GcpActionPermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpActionPermissionsDefinitionAction{}

func (s GcpActionPermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper GcpActionPermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpActionPermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpActionPermissionsDefinitionAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpActionPermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpActionPermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
