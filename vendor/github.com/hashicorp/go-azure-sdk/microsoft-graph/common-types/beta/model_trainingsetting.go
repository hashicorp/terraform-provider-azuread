package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingSetting interface {
	TrainingSetting() BaseTrainingSettingImpl
}

var _ TrainingSetting = BaseTrainingSettingImpl{}

type BaseTrainingSettingImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of setting. Possible values are: microsoftCustom, microsoftManaged, noTraining, custom, unknownFutureValue.
	SettingType *TrainingSettingType `json:"settingType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseTrainingSettingImpl) TrainingSetting() BaseTrainingSettingImpl {
	return s
}

var _ TrainingSetting = RawTrainingSettingImpl{}

// RawTrainingSettingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTrainingSettingImpl struct {
	trainingSetting BaseTrainingSettingImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawTrainingSettingImpl) TrainingSetting() BaseTrainingSettingImpl {
	return s.trainingSetting
}

func UnmarshalTrainingSettingImplementation(input []byte) (TrainingSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TrainingSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customTrainingSetting") {
		var out CustomTrainingSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomTrainingSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftCustomTrainingSetting") {
		var out MicrosoftCustomTrainingSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftCustomTrainingSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftManagedTrainingSetting") {
		var out MicrosoftManagedTrainingSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftManagedTrainingSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.microsoftTrainingAssignmentMapping") {
		var out MicrosoftTrainingAssignmentMapping
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftTrainingAssignmentMapping: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.noTrainingSetting") {
		var out NoTrainingSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoTrainingSetting: %+v", err)
		}
		return out, nil
	}

	var parent BaseTrainingSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTrainingSettingImpl: %+v", err)
	}

	return RawTrainingSettingImpl{
		trainingSetting: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
