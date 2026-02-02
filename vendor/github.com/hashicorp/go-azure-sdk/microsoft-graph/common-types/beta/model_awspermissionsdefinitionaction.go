package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsPermissionsDefinitionAction interface {
	PermissionsDefinitionAction
	AwsPermissionsDefinitionAction() BaseAwsPermissionsDefinitionActionImpl
}

var _ AwsPermissionsDefinitionAction = BaseAwsPermissionsDefinitionActionImpl{}

type BaseAwsPermissionsDefinitionActionImpl struct {

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAwsPermissionsDefinitionActionImpl) AwsPermissionsDefinitionAction() BaseAwsPermissionsDefinitionActionImpl {
	return s
}

func (s BaseAwsPermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AwsPermissionsDefinitionAction = RawAwsPermissionsDefinitionActionImpl{}

// RawAwsPermissionsDefinitionActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAwsPermissionsDefinitionActionImpl struct {
	awsPermissionsDefinitionAction BaseAwsPermissionsDefinitionActionImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawAwsPermissionsDefinitionActionImpl) AwsPermissionsDefinitionAction() BaseAwsPermissionsDefinitionActionImpl {
	return s.awsPermissionsDefinitionAction
}

func (s RawAwsPermissionsDefinitionActionImpl) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return s.awsPermissionsDefinitionAction.PermissionsDefinitionAction()
}

var _ json.Marshaler = BaseAwsPermissionsDefinitionActionImpl{}

func (s BaseAwsPermissionsDefinitionActionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAwsPermissionsDefinitionActionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAwsPermissionsDefinitionActionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAwsPermissionsDefinitionActionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsPermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAwsPermissionsDefinitionActionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAwsPermissionsDefinitionActionImplementation(input []byte) (AwsPermissionsDefinitionAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsPermissionsDefinitionAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsActionsPermissionsDefinitionAction") {
		var out AwsActionsPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsActionsPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.awsPolicyPermissionsDefinitionAction") {
		var out AwsPolicyPermissionsDefinitionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsPolicyPermissionsDefinitionAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseAwsPermissionsDefinitionActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAwsPermissionsDefinitionActionImpl: %+v", err)
	}

	return RawAwsPermissionsDefinitionActionImpl{
		awsPermissionsDefinitionAction: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
