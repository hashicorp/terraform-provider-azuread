package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSetting interface {
	EndUserNotificationSetting() BaseEndUserNotificationSettingImpl
}

var _ EndUserNotificationSetting = BaseEndUserNotificationSettingImpl{}

type BaseEndUserNotificationSettingImpl struct {
	// Notification preference. Possible values are: unknown, microsoft, custom, unknownFutureValue.
	NotificationPreference *EndUserNotificationPreference `json:"notificationPreference,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Positive reinforcement detail.
	PositiveReinforcement *PositiveReinforcementNotification `json:"positiveReinforcement,omitempty"`

	// End user notification type. Possible values are: unknown, noTraining, trainingSelected, noNotification,
	// unknownFutureValue.
	SettingType *EndUserNotificationSettingType `json:"settingType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEndUserNotificationSettingImpl) EndUserNotificationSetting() BaseEndUserNotificationSettingImpl {
	return s
}

var _ EndUserNotificationSetting = RawEndUserNotificationSettingImpl{}

// RawEndUserNotificationSettingImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEndUserNotificationSettingImpl struct {
	endUserNotificationSetting BaseEndUserNotificationSettingImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawEndUserNotificationSettingImpl) EndUserNotificationSetting() BaseEndUserNotificationSettingImpl {
	return s.endUserNotificationSetting
}

func UnmarshalEndUserNotificationSettingImplementation(input []byte) (EndUserNotificationSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EndUserNotificationSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.noTrainingNotificationSetting") {
		var out NoTrainingNotificationSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NoTrainingNotificationSetting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingNotificationSetting") {
		var out TrainingNotificationSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingNotificationSetting: %+v", err)
		}
		return out, nil
	}

	var parent BaseEndUserNotificationSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEndUserNotificationSettingImpl: %+v", err)
	}

	return RawEndUserNotificationSettingImpl{
		endUserNotificationSetting: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
