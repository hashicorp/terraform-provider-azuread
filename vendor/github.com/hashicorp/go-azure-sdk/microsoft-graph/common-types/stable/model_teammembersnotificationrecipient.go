package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamworkNotificationRecipient = TeamMembersNotificationRecipient{}

type TeamMembersNotificationRecipient struct {
	// The unique identifier for the team whose members should receive the notification.
	TeamId *string `json:"teamId,omitempty"`

	// Fields inherited from TeamworkNotificationRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeamMembersNotificationRecipient) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return BaseTeamworkNotificationRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TeamMembersNotificationRecipient{}

func (s TeamMembersNotificationRecipient) MarshalJSON() ([]byte, error) {
	type wrapper TeamMembersNotificationRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeamMembersNotificationRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamMembersNotificationRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teamMembersNotificationRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeamMembersNotificationRecipient: %+v", err)
	}

	return encoded, nil
}
