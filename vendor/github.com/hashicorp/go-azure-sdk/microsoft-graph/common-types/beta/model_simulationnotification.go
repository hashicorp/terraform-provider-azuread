package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseEndUserNotification = SimulationNotification{}

type SimulationNotification struct {
	// Target user type. Possible values are: unknown, clicked, compromised, allUsers, unknownFutureValue.
	TargettedUserType *TargettedUserType `json:"targettedUserType,omitempty"`

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

func (s SimulationNotification) BaseEndUserNotification() BaseBaseEndUserNotificationImpl {
	return BaseBaseEndUserNotificationImpl{
		DefaultLanguage:     s.DefaultLanguage,
		EndUserNotification: s.EndUserNotification,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
	}
}

var _ json.Marshaler = SimulationNotification{}

func (s SimulationNotification) MarshalJSON() ([]byte, error) {
	type wrapper SimulationNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SimulationNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SimulationNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.simulationNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SimulationNotification: %+v", err)
	}

	return encoded, nil
}
