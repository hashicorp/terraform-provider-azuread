package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkNotificationRecipient interface {
	TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl
}

var _ TeamworkNotificationRecipient = BaseTeamworkNotificationRecipientImpl{}

type BaseTeamworkNotificationRecipientImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseTeamworkNotificationRecipientImpl) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return s
}

var _ TeamworkNotificationRecipient = RawTeamworkNotificationRecipientImpl{}

// RawTeamworkNotificationRecipientImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawTeamworkNotificationRecipientImpl struct {
	teamworkNotificationRecipient BaseTeamworkNotificationRecipientImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawTeamworkNotificationRecipientImpl) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return s.teamworkNotificationRecipient
}

func (s RawTeamworkNotificationRecipientImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalTeamworkNotificationRecipientImplementation(input []byte) (TeamworkNotificationRecipient, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeamworkNotificationRecipient into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aadUserNotificationRecipient") {
		var out AadUserNotificationRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AadUserNotificationRecipient: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelMembersNotificationRecipient") {
		var out ChannelMembersNotificationRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelMembersNotificationRecipient: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMembersNotificationRecipient") {
		var out ChatMembersNotificationRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMembersNotificationRecipient: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamMembersNotificationRecipient") {
		var out TeamMembersNotificationRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamMembersNotificationRecipient: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeamworkNotificationRecipientImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeamworkNotificationRecipientImpl: %+v", err)
	}

	return RawTeamworkNotificationRecipientImpl{
		teamworkNotificationRecipient: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
