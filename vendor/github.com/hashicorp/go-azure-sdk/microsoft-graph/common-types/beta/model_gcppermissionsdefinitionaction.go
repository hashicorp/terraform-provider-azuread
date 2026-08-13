package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpPermissionsDefinitionAction interface {
	PermissionsDefinitionAction
	GcpPermissionsDefinitionAction() BaseGcpPermissionsDefinitionActionImpl
}

var _ GcpPermissionsDefinitionAction = BaseGcpPermissionsDefinitionActionImpl{}

type BaseGcpPermissionsDefinitionActionImpl struct {

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseGcpPermissionsDefinitionActionImpl) GcpPermissionsDefinitionAction() BaseGcpPermissionsDefinitionActionImpl {
	return s
}

func (s BaseGcpPermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GcpPermissionsDefinitionAction = RawGcpPermissionsDefinitionActionImpl{}

// RawGcpPermissionsDefinitionActionImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawGcpPermissionsDefinitionActionImpl struct {
	gcpPermissionsDefinitionAction BaseGcpPermissionsDefinitionActionImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawGcpPermissionsDefinitionActionImpl) GcpPermissionsDefinitionAction() BaseGcpPermissionsDefinitionActionImpl {
	return s.gcpPermissionsDefinitionAction
}

func (s RawGcpPermissionsDefinitionActionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func (s RawGcpPermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return s.gcpPermissionsDefinitionAction.PermissionsDefinitionAction()
}

var _ json.Marshaler = BaseGcpPermissionsDefinitionActionImpl{}

func (s BaseGcpPermissionsDefinitionActionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGcpPermissionsDefinitionActionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGcpPermissionsDefinitionActionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGcpPermissionsDefinitionActionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.gcpPermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGcpPermissionsDefinitionActionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalGcpPermissionsDefinitionActionImplementation(input []byte) (GcpPermissionsDefinitionAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GcpPermissionsDefinitionAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpActionPermissionsDefinitionAction") {
		var out GcpActionPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpActionPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpRolePermissionsDefinitionAction") {
		var out GcpRolePermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpRolePermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseGcpPermissionsDefinitionActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGcpPermissionsDefinitionActionImpl: %+v", err)
	}

	return RawGcpPermissionsDefinitionActionImpl{
		gcpPermissionsDefinitionAction: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
