package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseEndUserNotification = TrainingReminderNotification{}

type TrainingReminderNotification struct {
	// Configurable frequency for the reminder email introduced during simulation creation. Possible values are: unknown,
	// weekly, biWeekly, unknownFutureValue.
	DeliveryFrequency *NotificationDeliveryFrequency `json:"deliveryFrequency,omitempty"`

	// Fields inherited from BaseEndUserNotification

	// The default language for the end user notification.
	DefaultLanguage nullable.Type[string] `json:"defaultLanguage,omitempty"`

	EndUserNotification *EndUserNotification `json:"endUserNotification,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TrainingReminderNotification) BaseEndUserNotification() BaseBaseEndUserNotificationImpl {
	return BaseBaseEndUserNotificationImpl{
		DefaultLanguage:     s.DefaultLanguage,
		EndUserNotification: s.EndUserNotification,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

var _ json.Marshaler = TrainingReminderNotification{}

func (s TrainingReminderNotification) MarshalJSON() ([]byte, error) {
	type wrapper TrainingReminderNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TrainingReminderNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TrainingReminderNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.trainingReminderNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TrainingReminderNotification: %+v", err)
	}

	return encoded, nil
}
