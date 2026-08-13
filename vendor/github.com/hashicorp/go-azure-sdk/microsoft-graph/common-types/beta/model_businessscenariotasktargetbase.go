package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskTargetBase interface {
	BusinessScenarioTaskTargetBase() BaseBusinessScenarioTaskTargetBaseImpl
}

var _ BusinessScenarioTaskTargetBase = BaseBusinessScenarioTaskTargetBaseImpl{}

type BaseBusinessScenarioTaskTargetBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	TaskTargetKind *PlannerTaskTargetKind `json:"taskTargetKind,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseBusinessScenarioTaskTargetBaseImpl) BusinessScenarioTaskTargetBase() BaseBusinessScenarioTaskTargetBaseImpl {
	return s
}

var _ BusinessScenarioTaskTargetBase = RawBusinessScenarioTaskTargetBaseImpl{}

// RawBusinessScenarioTaskTargetBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawBusinessScenarioTaskTargetBaseImpl struct {
	businessScenarioTaskTargetBase BaseBusinessScenarioTaskTargetBaseImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawBusinessScenarioTaskTargetBaseImpl) BusinessScenarioTaskTargetBase() BaseBusinessScenarioTaskTargetBaseImpl {
	return s.businessScenarioTaskTargetBase
}

func (s RawBusinessScenarioTaskTargetBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalBusinessScenarioTaskTargetBaseImplementation(input []byte) (BusinessScenarioTaskTargetBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessScenarioTaskTargetBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.businessScenarioGroupTarget") {
		var out BusinessScenarioGroupTarget
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BusinessScenarioGroupTarget: %+v", err)
		}
		return out, nil
	}

	var parent BaseBusinessScenarioTaskTargetBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBusinessScenarioTaskTargetBaseImpl: %+v", err)
	}

	return RawBusinessScenarioTaskTargetBaseImpl{
		businessScenarioTaskTargetBase: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
