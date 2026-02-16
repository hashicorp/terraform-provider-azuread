package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndUserNotificationSetting = TrainingNotificationSetting{}

type TrainingNotificationSetting struct {
	// Training assignment details.
	TrainingAssignment BaseEndUserNotification `json:"trainingAssignment"`

	// Training reminder details.
	TrainingReminder *TrainingReminderNotification `json:"trainingReminder,omitempty"`

	// Fields inherited from EndUserNotificationSetting

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

func (s TrainingNotificationSetting) EndUserNotificationSetting() BaseEndUserNotificationSettingImpl {
	return BaseEndUserNotificationSettingImpl{
		NotificationPreference: s.NotificationPreference,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		PositiveReinforcement:  s.PositiveReinforcement,
		SettingType:            s.SettingType,
	}
}

var _ json.Marshaler = TrainingNotificationSetting{}

func (s TrainingNotificationSetting) MarshalJSON() ([]byte, error) {
	type wrapper TrainingNotificationSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TrainingNotificationSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TrainingNotificationSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trainingNotificationSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TrainingNotificationSetting: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TrainingNotificationSetting{}

func (s *TrainingNotificationSetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TrainingReminder       *TrainingReminderNotification      `json:"trainingReminder,omitempty"`
		NotificationPreference *EndUserNotificationPreference     `json:"notificationPreference,omitempty"`
		ODataId                *string                            `json:"@odata.id,omitempty"`
		ODataType              *string                            `json:"@odata.type,omitempty"`
		PositiveReinforcement  *PositiveReinforcementNotification `json:"positiveReinforcement,omitempty"`
		SettingType            *EndUserNotificationSettingType    `json:"settingType,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TrainingReminder = decoded.TrainingReminder
	s.NotificationPreference = decoded.NotificationPreference
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PositiveReinforcement = decoded.PositiveReinforcement
	s.SettingType = decoded.SettingType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TrainingNotificationSetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["trainingAssignment"]; ok {
		impl, err := UnmarshalBaseEndUserNotificationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingAssignment' for 'TrainingNotificationSetting': %+v", err)
		}
		s.TrainingAssignment = impl
	}

	return nil
}
