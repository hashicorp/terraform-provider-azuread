package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AzurePermissionsDefinitionAction = AzureActionPermissionsDefinitionAction{}

type AzureActionPermissionsDefinitionAction struct {
	// List of actions relating to the Azure permission.
	Actions *[]string `json:"actions,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AzureActionPermissionsDefinitionAction) AzurePermissionsDefinitionAction() BaseAzurePermissionsDefinitionActionImpl {
	return BaseAzurePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AzureActionPermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AzureActionPermissionsDefinitionAction{}

func (s AzureActionPermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper AzureActionPermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureActionPermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureActionPermissionsDefinitionAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azureActionPermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureActionPermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
