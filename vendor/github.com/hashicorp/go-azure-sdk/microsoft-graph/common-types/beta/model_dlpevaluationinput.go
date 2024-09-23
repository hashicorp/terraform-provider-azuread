package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpEvaluationInput interface {
	DlpEvaluationInput() BaseDlpEvaluationInputImpl
}

var _ DlpEvaluationInput = BaseDlpEvaluationInputImpl{}

type BaseDlpEvaluationInputImpl struct {
	CurrentLabel             *CurrentLabel              `json:"currentLabel,omitempty"`
	DiscoveredSensitiveTypes *[]DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDlpEvaluationInputImpl) DlpEvaluationInput() BaseDlpEvaluationInputImpl {
	return s
}

var _ DlpEvaluationInput = RawDlpEvaluationInputImpl{}

// RawDlpEvaluationInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDlpEvaluationInputImpl struct {
	dlpEvaluationInput BaseDlpEvaluationInputImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawDlpEvaluationInputImpl) DlpEvaluationInput() BaseDlpEvaluationInputImpl {
	return s.dlpEvaluationInput
}

func UnmarshalDlpEvaluationInputImplementation(input []byte) (DlpEvaluationInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DlpEvaluationInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.dlpEvaluationWindowsDevicesInput") {
		var out DlpEvaluationWindowsDevicesInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DlpEvaluationWindowsDevicesInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseDlpEvaluationInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDlpEvaluationInputImpl: %+v", err)
	}

	return RawDlpEvaluationInputImpl{
		dlpEvaluationInput: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
