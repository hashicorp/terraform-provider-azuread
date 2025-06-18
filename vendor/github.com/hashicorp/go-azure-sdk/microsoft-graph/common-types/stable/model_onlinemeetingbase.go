package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBase interface {
	Entity
	OnlineMeetingBase() BaseOnlineMeetingBaseImpl
}

var _ OnlineMeetingBase = BaseOnlineMeetingBaseImpl{}

type BaseOnlineMeetingBaseImpl struct {
	// Indicates whether attendees can turn on their camera.
	AllowAttendeeToEnableCamera nullable.Type[bool] `json:"allowAttendeeToEnableCamera,omitempty"`

	// Indicates whether attendees can turn on their microphone.
	AllowAttendeeToEnableMic nullable.Type[bool] `json:"allowAttendeeToEnableMic,omitempty"`

	// Indicates whether breakout rooms are enabled for the meeting.
	AllowBreakoutRooms nullable.Type[bool] `json:"allowBreakoutRooms,omitempty"`

	// Indicates whether live share is enabled for the meeting. Possible values are: enabled, disabled, unknownFutureValue.
	AllowLiveShare *MeetingLiveShareOptions `json:"allowLiveShare,omitempty"`

	// Specifies the mode of the meeting chat.
	AllowMeetingChat *MeetingChatMode `json:"allowMeetingChat,omitempty"`

	// Specifies if participants are allowed to rename themselves in an instance of the meeting.
	AllowParticipantsToChangeName nullable.Type[bool] `json:"allowParticipantsToChangeName,omitempty"`

	// Indicates whether PowerPoint live is enabled for the meeting.
	AllowPowerPointSharing nullable.Type[bool] `json:"allowPowerPointSharing,omitempty"`

	// Indicates whether recording is enabled for the meeting.
	AllowRecording nullable.Type[bool] `json:"allowRecording,omitempty"`

	// Indicates if Teams reactions are enabled for the meeting.
	AllowTeamworkReactions nullable.Type[bool] `json:"allowTeamworkReactions,omitempty"`

	// Indicates whether transcription is enabled for the meeting.
	AllowTranscription nullable.Type[bool] `json:"allowTranscription,omitempty"`

	// Indicates whether whiteboard is enabled for the meeting.
	AllowWhiteboard nullable.Type[bool] `json:"allowWhiteboard,omitempty"`

	// Specifies the users who can admit from the lobby. Possible values are: organizerAndCoOrganizersAndPresenters,
	// organizerAndCoOrganizers, unknownFutureValue.
	AllowedLobbyAdmitters *AllowedLobbyAdmitterRoles `json:"allowedLobbyAdmitters,omitempty"`

	// Specifies who can be a presenter in a meeting.
	AllowedPresenters *OnlineMeetingPresenters `json:"allowedPresenters,omitempty"`

	// The attendance reports of an online meeting. Read-only.
	AttendanceReports *[]MeetingAttendanceReport `json:"attendanceReports,omitempty"`

	// The phone access (dial-in) information for an online meeting. Read-only.
	AudioConferencing *AudioConferencing `json:"audioConferencing,omitempty"`

	// The chat information associated with this online meeting.
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`

	// Specifies the configuration settings for meeting chat restrictions.
	ChatRestrictions *ChatRestrictions `json:"chatRestrictions,omitempty"`

	// Indicates whether to announce when callers join or leave.
	IsEntryExitAnnounced nullable.Type[bool] `json:"isEntryExitAnnounced,omitempty"`

	// The join information in the language and locale variant specified in 'Accept-Language' request HTTP header.
	// Read-only.
	JoinInformation *ItemBody `json:"joinInformation,omitempty"`

	// Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is
	// created, the joinMeetingIdSettings can't be modified. To make any changes to this property, you must cancel this
	// meeting and create a new one.
	JoinMeetingIdSettings *JoinMeetingIdSettings `json:"joinMeetingIdSettings,omitempty"`

	// The join URL of the online meeting. Read-only.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

	// Specifies which participants can bypass the meeting lobby.
	LobbyBypassSettings *LobbyBypassSettings `json:"lobbyBypassSettings,omitempty"`

	// Indicates whether to record the meeting automatically.
	RecordAutomatically nullable.Type[bool] `json:"recordAutomatically,omitempty"`

	// Specifies whether meeting chat history is shared with participants. Possible values are: all, none,
	// unknownFutureValue.
	ShareMeetingChatHistoryDefault *MeetingChatHistoryDefaultMode `json:"shareMeetingChatHistoryDefault,omitempty"`

	// The subject of the online meeting.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The video teleconferencing ID. Read-only.
	VideoTeleconferenceId nullable.Type[string] `json:"videoTeleconferenceId,omitempty"`

	// Specifies whether the client application should apply a watermark to a content type.
	WatermarkProtection *WatermarkProtectionValues `json:"watermarkProtection,omitempty"`

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

func (s BaseOnlineMeetingBaseImpl) OnlineMeetingBase() BaseOnlineMeetingBaseImpl {
	return s
}

func (s BaseOnlineMeetingBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OnlineMeetingBase = RawOnlineMeetingBaseImpl{}

// RawOnlineMeetingBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnlineMeetingBaseImpl struct {
	onlineMeetingBase BaseOnlineMeetingBaseImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawOnlineMeetingBaseImpl) OnlineMeetingBase() BaseOnlineMeetingBaseImpl {
	return s.onlineMeetingBase
}

func (s RawOnlineMeetingBaseImpl) Entity() BaseEntityImpl {
	return s.onlineMeetingBase.Entity()
}

var _ json.Marshaler = BaseOnlineMeetingBaseImpl{}

func (s BaseOnlineMeetingBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOnlineMeetingBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOnlineMeetingBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOnlineMeetingBaseImpl: %+v", err)
	}

	delete(decoded, "attendanceReports")
	delete(decoded, "audioConferencing")
	delete(decoded, "joinInformation")
	delete(decoded, "joinWebUrl")
	delete(decoded, "videoTeleconferenceId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onlineMeetingBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOnlineMeetingBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOnlineMeetingBaseImplementation(input []byte) (OnlineMeetingBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnlineMeetingBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onlineMeeting") {
		var out OnlineMeeting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnlineMeeting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventSession") {
		var out VirtualEventSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventSession: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnlineMeetingBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnlineMeetingBaseImpl: %+v", err)
	}

	return RawOnlineMeetingBaseImpl{
		onlineMeetingBase: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
