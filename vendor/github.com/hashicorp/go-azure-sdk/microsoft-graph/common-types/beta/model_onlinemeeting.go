package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OnlineMeetingBase = OnlineMeeting{}

type OnlineMeeting struct {
	// The AI insights generated for an online meeting. Read-only.
	AiInsights *[]CallAiInsight `json:"aiInsights,omitempty"`

	// The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
	AlternativeRecording nullable.Type[string] `json:"alternativeRecording,omitempty"`

	// The content stream of the attendee report of a Teams live event. Read-only.
	AttendeeReport nullable.Type[string] `json:"attendeeReport,omitempty"`

	BroadcastRecording nullable.Type[string] `json:"broadcastRecording,omitempty"`

	// Settings related to a live event.
	BroadcastSettings *BroadcastMeetingSettings `json:"broadcastSettings,omitempty"`

	// The list of meeting capabilities. Possible values are: questionAndAnswer,unknownFutureValue.
	Capabilities *[]MeetingCapabilities `json:"capabilities,omitempty"`

	// The meeting creation time in UTC. Read-only.
	CreationDateTime nullable.Type[string] `json:"creationDateTime,omitempty"`

	// The meeting end time in UTC. Required when you create an online meeting.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The external ID. A custom ID. Optional.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Indicates whether this event is a Teams live event.
	IsBroadcast nullable.Type[bool] `json:"isBroadcast,omitempty"`

	JoinUrl nullable.Type[string] `json:"joinUrl,omitempty"`

	// The attendance report of the latest online meeting session. Read-only.
	MeetingAttendanceReport *MeetingAttendanceReport `json:"meetingAttendanceReport,omitempty"`

	// The ID of the meeting template.
	MeetingTemplateId nullable.Type[string] `json:"meetingTemplateId,omitempty"`

	// The participants associated with the online meeting, including the organizer and the attendees.
	Participants *MeetingParticipants `json:"participants,omitempty"`

	// The content stream of the recording of a Teams live event. Read-only.
	Recording nullable.Type[string] `json:"recording,omitempty"`

	// The recordings of an online meeting. Read-only.
	Recordings *[]CallRecording `json:"recordings,omitempty"`

	// The registration that is enabled for an online meeting. One online meeting can only have one registration enabled.
	Registration *MeetingRegistration `json:"registration,omitempty"`

	// The meeting start time in UTC. Required when you create an online meeting.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The transcripts of an online meeting. Read-only.
	Transcripts *[]CallTranscript `json:"transcripts,omitempty"`

	// Fields inherited from OnlineMeetingBase

	// Indicates whether attendees can turn on their camera.
	AllowAttendeeToEnableCamera nullable.Type[bool] `json:"allowAttendeeToEnableCamera,omitempty"`

	// Indicates whether attendees can turn on their microphone.
	AllowAttendeeToEnableMic nullable.Type[bool] `json:"allowAttendeeToEnableMic,omitempty"`

	// Indicates whether breakout rooms are enabled for the meeting.
	AllowBreakoutRooms nullable.Type[bool] `json:"allowBreakoutRooms,omitempty"`

	// Indicates whether copying and sharing meeting content is enabled for the meeting.
	AllowCopyingAndSharingMeetingContent nullable.Type[bool] `json:"allowCopyingAndSharingMeetingContent,omitempty"`

	// Indicates whether live share is enabled for the meeting. Possible values are: enabled, disabled, unknownFutureValue.
	AllowLiveShare *MeetingLiveShareOptions `json:"allowLiveShare,omitempty"`

	// Specifies the mode of meeting chat. Possible values are: enabled, disabled, limited, unknownFutureValue.
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

	// Specifies who can be a presenter in a meeting. Possible values are: everyone, organization, roleIsPresenter,
	// organizer, unknownFutureValue.
	AllowedPresenters *OnlineMeetingPresenters `json:"allowedPresenters,omitempty"`

	// Specifies whose identity is anonymized in the meeting. Possible values are: attendee. The attendee value can't be
	// removed through a PATCH operation once added.
	AnonymizeIdentityForRoles *[]OnlineMeetingRole `json:"anonymizeIdentityForRoles,omitempty"`

	// The attendance reports of an online meeting. Read-only.
	AttendanceReports *[]MeetingAttendanceReport `json:"attendanceReports,omitempty"`

	// The phone access (dial-in) information for an online meeting. Read-only.
	AudioConferencing *AudioConferencing `json:"audioConferencing,omitempty"`

	// The chat information associated with this online meeting.
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`

	// Specifies the configuration settings for meeting chat restrictions.
	ChatRestrictions *ChatRestrictions `json:"chatRestrictions,omitempty"`

	// Indicates whether end-to-end encryption (E2EE) is enabled for the online meeting.
	IsEndToEndEncryptionEnabled nullable.Type[bool] `json:"isEndToEndEncryptionEnabled,omitempty"`

	// Indicates whether to announce when callers join or leave.
	IsEntryExitAnnounced nullable.Type[bool] `json:"isEntryExitAnnounced,omitempty"`

	// The join information in the language and locale variant specified in 'Accept-Language' request HTTP header.
	// Read-only.
	JoinInformation *ItemBody `json:"joinInformation,omitempty"`

	// Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is
	// created, the joinMeetingIdSettings can't be modified. To make any changes to this property, the meeting needs to be
	// canceled and a new one needs to be created.
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
		AllowAttendeeToEnableCamera:          s.AllowAttendeeToEnableCamera,
		AllowAttendeeToEnableMic:             s.AllowAttendeeToEnableMic,
		AllowBreakoutRooms:                   s.AllowBreakoutRooms,
		AllowCopyingAndSharingMeetingContent: s.AllowCopyingAndSharingMeetingContent,
		AllowLiveShare:                       s.AllowLiveShare,
		AllowMeetingChat:                     s.AllowMeetingChat,
		AllowParticipantsToChangeName:        s.AllowParticipantsToChangeName,
		AllowPowerPointSharing:               s.AllowPowerPointSharing,
		AllowRecording:                       s.AllowRecording,
		AllowTeamworkReactions:               s.AllowTeamworkReactions,
		AllowTranscription:                   s.AllowTranscription,
		AllowWhiteboard:                      s.AllowWhiteboard,
		AllowedLobbyAdmitters:                s.AllowedLobbyAdmitters,
		AllowedPresenters:                    s.AllowedPresenters,
		AnonymizeIdentityForRoles:            s.AnonymizeIdentityForRoles,
		AttendanceReports:                    s.AttendanceReports,
		AudioConferencing:                    s.AudioConferencing,
		ChatInfo:                             s.ChatInfo,
		ChatRestrictions:                     s.ChatRestrictions,
		IsEndToEndEncryptionEnabled:          s.IsEndToEndEncryptionEnabled,
		IsEntryExitAnnounced:                 s.IsEntryExitAnnounced,
		JoinInformation:                      s.JoinInformation,
		JoinMeetingIdSettings:                s.JoinMeetingIdSettings,
		JoinWebUrl:                           s.JoinWebUrl,
		LobbyBypassSettings:                  s.LobbyBypassSettings,
		RecordAutomatically:                  s.RecordAutomatically,
		ShareMeetingChatHistoryDefault:       s.ShareMeetingChatHistoryDefault,
		Subject:                              s.Subject,
		VideoTeleconferenceId:                s.VideoTeleconferenceId,
		WatermarkProtection:                  s.WatermarkProtection,
		Id:                                   s.Id,
		ODataId:                              s.ODataId,
		ODataType:                            s.ODataType,
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

	delete(decoded, "aiInsights")
	delete(decoded, "alternativeRecording")
	delete(decoded, "attendeeReport")
	delete(decoded, "creationDateTime")
	delete(decoded, "meetingAttendanceReport")
	delete(decoded, "recording")
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
