package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GcpPermissionsDefinitionAction = GcpRolePermissionsDefinitionAction{}

type GcpRolePermissionsDefinitionAction struct {
	Roles *[]PermissionsDefinitionGcpRole `json:"roles,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s GcpRolePermissionsDefinitionAction) GcpPermissionsDefinitionAction() BaseGcpPermissionsDefinitionActionImpl {
	return BaseGcpPermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s GcpRolePermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GcpRolePermissionsDefinitionAction{}

func (s GcpRolePermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper GcpRolePermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GcpRolePermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpRolePermissionsDefinitionAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpRolePermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GcpRolePermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
