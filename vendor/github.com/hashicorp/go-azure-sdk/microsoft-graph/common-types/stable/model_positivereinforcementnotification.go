package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseEndUserNotification = PositiveReinforcementNotification{}

type PositiveReinforcementNotification struct {
	// Delivery preference. Possible values are: unknown, deliverImmedietly, deliverAfterCampaignEnd, unknownFutureValue.
	DeliveryPreference *NotificationDeliveryPreference `json:"deliveryPreference,omitempty"`

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

func (s PositiveReinforcementNotification) BaseEndUserNotification() BaseBaseEndUserNotificationImpl {
	return BaseBaseEndUserNotificationImpl{
		DefaultLanguage:     s.DefaultLanguage,
		EndUserNotification: s.EndUserNotification,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

var _ json.Marshaler = PositiveReinforcementNotification{}

func (s PositiveReinforcementNotification) MarshalJSON() ([]byte, error) {
	type wrapper PositiveReinforcementNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PositiveReinforcementNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PositiveReinforcementNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.positiveReinforcementNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PositiveReinforcementNotification: %+v", err)
	}

	return encoded, nil
}
