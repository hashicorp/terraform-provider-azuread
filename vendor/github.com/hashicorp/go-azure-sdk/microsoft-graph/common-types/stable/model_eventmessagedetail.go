package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageDetail interface {
	EventMessageDetail() BaseEventMessageDetailImpl
}

var _ EventMessageDetail = BaseEventMessageDetailImpl{}

type BaseEventMessageDetailImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEventMessageDetailImpl) EventMessageDetail() BaseEventMessageDetailImpl {
	return s
}

var _ EventMessageDetail = RawEventMessageDetailImpl{}

// RawEventMessageDetailImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEventMessageDetailImpl struct {
	eventMessageDetail BaseEventMessageDetailImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawEventMessageDetailImpl) EventMessageDetail() BaseEventMessageDetailImpl {
	return s.eventMessageDetail
}

func UnmarshalEventMessageDetailImplementation(input []byte) (EventMessageDetail, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EventMessageDetail into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.callEndedEventMessageDetail") {
		var out CallEndedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallEndedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callRecordingEventMessageDetail") {
		var out CallRecordingEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallRecordingEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callStartedEventMessageDetail") {
		var out CallStartedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallStartedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.callTranscriptEventMessageDetail") {
		var out CallTranscriptEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CallTranscriptEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelAddedEventMessageDetail") {
		var out ChannelAddedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelAddedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelDeletedEventMessageDetail") {
		var out ChannelDeletedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelDeletedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelDescriptionUpdatedEventMessageDetail") {
		var out ChannelDescriptionUpdatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelDescriptionUpdatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelRenamedEventMessageDetail") {
		var out ChannelRenamedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelRenamedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelSetAsFavoriteByDefaultEventMessageDetail") {
		var out ChannelSetAsFavoriteByDefaultEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelSetAsFavoriteByDefaultEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.channelUnsetAsFavoriteByDefaultEventMessageDetail") {
		var out ChannelUnsetAsFavoriteByDefaultEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChannelUnsetAsFavoriteByDefaultEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatRenamedEventMessageDetail") {
		var out ChatRenamedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatRenamedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conversationMemberRoleUpdatedEventMessageDetail") {
		var out ConversationMemberRoleUpdatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConversationMemberRoleUpdatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingPolicyUpdatedEventMessageDetail") {
		var out MeetingPolicyUpdatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingPolicyUpdatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.membersAddedEventMessageDetail") {
		var out MembersAddedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembersAddedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.membersDeletedEventMessageDetail") {
		var out MembersDeletedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembersDeletedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.membersJoinedEventMessageDetail") {
		var out MembersJoinedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembersJoinedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.membersLeftEventMessageDetail") {
		var out MembersLeftEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MembersLeftEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messagePinnedEventMessageDetail") {
		var out MessagePinnedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessagePinnedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.messageUnpinnedEventMessageDetail") {
		var out MessageUnpinnedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MessageUnpinnedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.tabUpdatedEventMessageDetail") {
		var out TabUpdatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TabUpdatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamArchivedEventMessageDetail") {
		var out TeamArchivedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamArchivedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamCreatedEventMessageDetail") {
		var out TeamCreatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamCreatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamDescriptionUpdatedEventMessageDetail") {
		var out TeamDescriptionUpdatedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamDescriptionUpdatedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamJoiningDisabledEventMessageDetail") {
		var out TeamJoiningDisabledEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamJoiningDisabledEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamJoiningEnabledEventMessageDetail") {
		var out TeamJoiningEnabledEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamJoiningEnabledEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamRenamedEventMessageDetail") {
		var out TeamRenamedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamRenamedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamUnarchivedEventMessageDetail") {
		var out TeamUnarchivedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamUnarchivedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppInstalledEventMessageDetail") {
		var out TeamsAppInstalledEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppInstalledEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppRemovedEventMessageDetail") {
		var out TeamsAppRemovedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppRemovedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teamsAppUpgradedEventMessageDetail") {
		var out TeamsAppUpgradedEventMessageDetail
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeamsAppUpgradedEventMessageDetail: %+v", err)
		}
		return out, nil
	}

	var parent BaseEventMessageDetailImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEventMessageDetailImpl: %+v", err)
	}

	return RawEventMessageDetailImpl{
		eventMessageDetail: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
