package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EndUserNotificationSetting = NoTrainingNotificationSetting{}

type NoTrainingNotificationSetting struct {
	// The notification for the user who is part of the simulation.
	SimulationNotification *SimulationNotification `json:"simulationNotification,omitempty"`

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

func (s NoTrainingNotificationSetting) EndUserNotificationSetting() BaseEndUserNotificationSettingImpl {
	return BaseEndUserNotificationSettingImpl{
		NotificationPreference: s.NotificationPreference,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		PositiveReinforcement:  s.PositiveReinforcement,
		SettingType:            s.SettingType,
	}
}

var _ json.Marshaler = NoTrainingNotificationSetting{}

func (s NoTrainingNotificationSetting) MarshalJSON() ([]byte, error) {
	type wrapper NoTrainingNotificationSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NoTrainingNotificationSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NoTrainingNotificationSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.noTrainingNotificationSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NoTrainingNotificationSetting: %+v", err)
	}

	return encoded, nil
}
