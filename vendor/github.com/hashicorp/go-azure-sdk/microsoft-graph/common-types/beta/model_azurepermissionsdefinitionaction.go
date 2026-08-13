package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzurePermissionsDefinitionAction interface {
	PermissionsDefinitionAction
	AzurePermissionsDefinitionAction() BaseAzurePermissionsDefinitionActionImpl
}

var _ AzurePermissionsDefinitionAction = BaseAzurePermissionsDefinitionActionImpl{}

type BaseAzurePermissionsDefinitionActionImpl struct {

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAzurePermissionsDefinitionActionImpl) AzurePermissionsDefinitionAction() BaseAzurePermissionsDefinitionActionImpl {
	return s
}

func (s BaseAzurePermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AzurePermissionsDefinitionAction = RawAzurePermissionsDefinitionActionImpl{}

// RawAzurePermissionsDefinitionActionImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAzurePermissionsDefinitionActionImpl struct {
	azurePermissionsDefinitionAction BaseAzurePermissionsDefinitionActionImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawAzurePermissionsDefinitionActionImpl) AzurePermissionsDefinitionAction() BaseAzurePermissionsDefinitionActionImpl {
	return s.azurePermissionsDefinitionAction
}

func (s RawAzurePermissionsDefinitionActionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func (s RawAzurePermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return s.azurePermissionsDefinitionAction.PermissionsDefinitionAction()
}

var _ json.Marshaler = BaseAzurePermissionsDefinitionActionImpl{}

func (s BaseAzurePermissionsDefinitionActionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAzurePermissionsDefinitionActionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAzurePermissionsDefinitionActionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAzurePermissionsDefinitionActionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.azurePermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAzurePermissionsDefinitionActionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAzurePermissionsDefinitionActionImplementation(input []byte) (AzurePermissionsDefinitionAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AzurePermissionsDefinitionAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.azureActionPermissionsDefinitionAction") {
		var out AzureActionPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureActionPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureRolePermissionsDefinitionAction") {
		var out AzureRolePermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureRolePermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseAzurePermissionsDefinitionActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAzurePermissionsDefinitionActionImpl: %+v", err)
	}

	return RawAzurePermissionsDefinitionActionImpl{
		azurePermissionsDefinitionAction: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
