package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseEndUserNotification interface {
	BaseEndUserNotification() BaseBaseEndUserNotificationImpl
}

var _ BaseEndUserNotification = BaseBaseEndUserNotificationImpl{}

type BaseBaseEndUserNotificationImpl struct {
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

func (s BaseBaseEndUserNotificationImpl) BaseEndUserNotification() BaseBaseEndUserNotificationImpl {
	return s
}

var _ BaseEndUserNotification = RawBaseEndUserNotificationImpl{}

// RawBaseEndUserNotificationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBaseEndUserNotificationImpl struct {
	baseEndUserNotification BaseBaseEndUserNotificationImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawBaseEndUserNotificationImpl) BaseEndUserNotification() BaseBaseEndUserNotificationImpl {
	return s.baseEndUserNotification
}

func UnmarshalBaseEndUserNotificationImplementation(input []byte) (BaseEndUserNotification, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEndUserNotification into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.positiveReinforcementNotification") {
		var out PositiveReinforcementNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PositiveReinforcementNotification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.simulationNotification") {
		var out SimulationNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SimulationNotification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.trainingReminderNotification") {
		var out TrainingReminderNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TrainingReminderNotification: %+v", err)
		}
		return out, nil
	}

	var parent BaseBaseEndUserNotificationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBaseEndUserNotificationImpl: %+v", err)
	}

	return RawBaseEndUserNotificationImpl{
		baseEndUserNotification: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
