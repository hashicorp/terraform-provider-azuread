package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppIntentAndState{}

type MobileAppIntentAndState struct {
	// Device identifier created or collected by Intune.
	ManagedDeviceIdentifier nullable.Type[string] `json:"managedDeviceIdentifier,omitempty"`

	// The list of payload intents and states for the tenant.
	MobileAppList *[]MobileAppIntentAndStateDetail `json:"mobileAppList,omitempty"`

	// Identifier for the user that tried to enroll the device.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MobileAppIntentAndState) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppIntentAndState{}

func (s MobileAppIntentAndState) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppIntentAndState
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppIntentAndState: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppIntentAndState: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppIntentAndState"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppIntentAndState: %+v", err)
	}

	return encoded, nil
}
