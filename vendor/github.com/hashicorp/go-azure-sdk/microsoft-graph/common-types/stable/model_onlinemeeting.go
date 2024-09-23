package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnlineMeetingBase = OnlineMeeting{}

type OnlineMeeting struct {
	// The content stream of the attendee report of a Microsoft Teams live event. Read-only.
	AttendeeReport nullable.Type[string] `json:"attendeeReport,omitempty"`

	// Settings related to a live event.
	BroadcastSettings *BroadcastMeetingSettings `json:"broadcastSettings,omitempty"`

	// The meeting creation time in UTC. Read-only.
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	// The meeting end time in UTC. Required when you create an online meeting.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Indicates whether this meeting is a Teams live event.
	IsBroadcast nullable.Type[bool] `json:"isBroadcast,omitempty"`

	// The participants associated with the online meeting, including the organizer and the attendees.
	Participants *MeetingParticipants `json:"participants,omitempty"`

	// The recordings of an online meeting. Read-only.
	Recordings *[]CallRecording `json:"recordings,omitempty"`

	// The meeting start time in UTC.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The transcripts of an online meeting. Read-only.
	Transcripts *[]CallTranscript `json:"transcripts,omitempty"`

	// Fields inherited from OnlineMeetingBase

	// Indicates whether attendees can turn on their camera.
	AllowAttendeeToEnableCamera nullable.Type[bool] `json:"allowAttendeeToEnableCamera,omitempty"`

	// Indicates whether attendees can turn on their microphone.
	AllowAttendeeToEnableMic nullable.Type[bool] `json:"allowAttendeeToEnableMic,omitempty"`

	// Specifies the mode of the meeting chat.
	AllowMeetingChat *MeetingChatMode `json:"allowMeetingChat,omitempty"`

	// Specifies if participants are allowed to rename themselves in an instance of the meeting.
	AllowParticipantsToChangeName nullable.Type[bool] `json:"allowParticipantsToChangeName,omitempty"`

	// Indicates if Teams reactions are enabled for the meeting.
	AllowTeamworkReactions nullable.Type[bool] `json:"allowTeamworkReactions,omitempty"`

	// Specifies who can be a presenter in a meeting.
	AllowedPresenters *OnlineMeetingPresenters `json:"allowedPresenters,omitempty"`

	// The attendance reports of an online meeting. Read-only.
	AttendanceReports *[]MeetingAttendanceReport `json:"attendanceReports,omitempty"`

	// The phone access (dial-in) information for an online meeting. Read-only.
	AudioConferencing *AudioConferencing `json:"audioConferencing,omitempty"`

	// The chat information associated with this online meeting.
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`

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

func (s OnlineMeeting) OnlineMeetingBase() BaseOnlineMeetingBaseImpl {
	return BaseOnlineMeetingBaseImpl{
		AllowAttendeeToEnableCamera:    s.AllowAttendeeToEnableCamera,
		AllowAttendeeToEnableMic:       s.AllowAttendeeToEnableMic,
		AllowMeetingChat:               s.AllowMeetingChat,
		AllowParticipantsToChangeName:  s.AllowParticipantsToChangeName,
		AllowTeamworkReactions:         s.AllowTeamworkReactions,
		AllowedPresenters:              s.AllowedPresenters,
		AttendanceReports:              s.AttendanceReports,
		AudioConferencing:              s.AudioConferencing,
		ChatInfo:                       s.ChatInfo,
		IsEntryExitAnnounced:           s.IsEntryExitAnnounced,
		JoinInformation:                s.JoinInformation,
		JoinMeetingIdSettings:          s.JoinMeetingIdSettings,
		JoinWebUrl:                     s.JoinWebUrl,
		LobbyBypassSettings:            s.LobbyBypassSettings,
		RecordAutomatically:            s.RecordAutomatically,
		ShareMeetingChatHistoryDefault: s.ShareMeetingChatHistoryDefault,
		Subject:                        s.Subject,
		VideoTeleconferenceId:          s.VideoTeleconferenceId,
		WatermarkProtection:            s.WatermarkProtection,
		Id:                             s.Id,
		ODataId:                        s.ODataId,
		ODataType:                      s.ODataType,
	}
}

func (s OnlineMeeting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnlineMeeting{}

func (s OnlineMeeting) MarshalJSON() ([]byte, error) {
	type wrapper OnlineMeeting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnlineMeeting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnlineMeeting: %+v", err)
	}

	delete(decoded, "attendeeReport")
	delete(decoded, "creationDateTime")
	delete(decoded, "recordings")
	delete(decoded, "transcripts")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onlineMeeting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnlineMeeting: %+v", err)
	}

	return encoded, nil
}
