package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamworkNotificationRecipient = AadUserNotificationRecipient{}

type AadUserNotificationRecipient struct {
	// Microsoft Entra user identifier. Use the List users method to get this ID.
	UserId *string `json:"userId,omitempty"`

	// Fields inherited from TeamworkNotificationRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AadUserNotificationRecipient) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return BaseTeamworkNotificationRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AadUserNotificationRecipient{}

func (s AadUserNotificationRecipient) MarshalJSON() ([]byte, error) {
	type wrapper AadUserNotificationRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AadUserNotificationRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AadUserNotificationRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.aadUserNotificationRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AadUserNotificationRecipient: %+v", err)
	}

	return encoded, nil
}
