package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Call{}

type Call struct {
	// The list of active modalities. Possible values are: unknown, audio, video, videoBasedScreenSharing, data. Read-only.
	ActiveModalities *[]Modality `json:"activeModalities,omitempty"`

	// The participant that answered the call. Read-only.
	AnsweredBy *ParticipantInfo `json:"answeredBy,omitempty"`

	AudioRoutingGroups *[]AudioRoutingGroup `json:"audioRoutingGroups,omitempty"`

	// A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in
	// a P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
	CallChainId nullable.Type[string] `json:"callChainId,omitempty"`

	// Contains the optional features for the call.
	CallOptions CallOptions `json:"callOptions"`

	// The routing information on how the call was retargeted. Read-only.
	CallRoutes *[]CallRoute `json:"callRoutes,omitempty"`

	// The callback URL on which callbacks will be delivered. Must be https.
	CallbackUri *string `json:"callbackUri,omitempty"`

	// The chat information. Required information for meeting scenarios.
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`

	ContentSharingSessions *[]ContentSharingSession `json:"contentSharingSessions,omitempty"`

	// The direction of the call. The possible values are incoming or outgoing. Read-only.
	Direction *CallDirection `json:"direction,omitempty"`

	// The context associated with an incoming call. Read-only. Server generated.
	IncomingContext *IncomingContext `json:"incomingContext,omitempty"`

	// The media configuration. Required information for creating peer to peer calls or joining meetings.
	MediaConfig MediaConfig `json:"mediaConfig"`

	// Read-only. The call media state.
	MediaState *CallMediaState `json:"mediaState,omitempty"`

	// Contains the capabilities of a meeting. Read-only.
	MeetingCapability *MeetingCapability `json:"meetingCapability,omitempty"`

	// The meeting information. Required information for meeting scenarios.
	MeetingInfo MeetingInfo `json:"meetingInfo"`

	MyParticipantId nullable.Type[string] `json:"myParticipantId,omitempty"`
	Operations      *[]CommsOperation     `json:"operations,omitempty"`
	Participants    *[]Participant        `json:"participants,omitempty"`

	// The list of requested modalities. Possible values are: unknown, audio, video, videoBasedScreenSharing, data.
	RequestedModalities *[]Modality `json:"requestedModalities,omitempty"`

	// The result information. For example, the result can hold termination reason. Read-only.
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`

	// Ringing timeout in seconds for outgoing peer to peer calls. The max value for this attribute is 115 seconds.
	RingingTimeoutInSeconds nullable.Type[int64] `json:"ringingTimeoutInSeconds,omitempty"`

	// This property is applicable for peer to peer calls only. Possible values are: none, noMissedCall,
	// disableForwardingExceptPhone, disableForwarding, preferSkypeForBusiness, unknownFutureValue.
	RoutingPolicies *[]RoutingPolicy `json:"routingPolicies,omitempty"`

	// The originator of the call.
	Source *ParticipantInfo `json:"source,omitempty"`

	// The call state. Possible values are: incoming, establishing, ringing, established, hold, transferring,
	// transferAccepted, redirecting, terminating, terminated. Read-only.
	State *CallState `json:"state,omitempty"`

	// The subject of the conversation.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// The targets of the call. Required information for creating peer to peer call.
	Targets *[]InvitationParticipantInfo `json:"targets,omitempty"`

	// Read-only. tenantId in Microsoft Entra ID.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	TerminationReason nullable.Type[string] `json:"terminationReason,omitempty"`
	ToneInfo          *ToneInfo             `json:"toneInfo,omitempty"`

	// The transcription information for the call. Read-only.
	Transcription *CallTranscriptionInfo `json:"transcription,omitempty"`

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

func (s Call) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Call{}

func (s Call) MarshalJSON() ([]byte, error) {
	type wrapper Call
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Call: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Call: %+v", err)
	}

	delete(decoded, "activeModalities")
	delete(decoded, "answeredBy")
	delete(decoded, "callRoutes")
	delete(decoded, "direction")
	delete(decoded, "incomingContext")
	delete(decoded, "mediaState")
	delete(decoded, "meetingCapability")
	delete(decoded, "resultInfo")
	delete(decoded, "state")
	delete(decoded, "tenantId")
	delete(decoded, "transcription")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.call"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Call: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Call{}

func (s *Call) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActiveModalities        *[]Modality                  `json:"activeModalities,omitempty"`
		AnsweredBy              *ParticipantInfo             `json:"answeredBy,omitempty"`
		AudioRoutingGroups      *[]AudioRoutingGroup         `json:"audioRoutingGroups,omitempty"`
		CallChainId             nullable.Type[string]        `json:"callChainId,omitempty"`
		CallRoutes              *[]CallRoute                 `json:"callRoutes,omitempty"`
		CallbackUri             *string                      `json:"callbackUri,omitempty"`
		ChatInfo                *ChatInfo                    `json:"chatInfo,omitempty"`
		ContentSharingSessions  *[]ContentSharingSession     `json:"contentSharingSessions,omitempty"`
		Direction               *CallDirection               `json:"direction,omitempty"`
		IncomingContext         *IncomingContext             `json:"incomingContext,omitempty"`
		MediaState              *CallMediaState              `json:"mediaState,omitempty"`
		MeetingCapability       *MeetingCapability           `json:"meetingCapability,omitempty"`
		MyParticipantId         nullable.Type[string]        `json:"myParticipantId,omitempty"`
		Participants            *[]Participant               `json:"participants,omitempty"`
		RequestedModalities     *[]Modality                  `json:"requestedModalities,omitempty"`
		ResultInfo              *ResultInfo                  `json:"resultInfo,omitempty"`
		RingingTimeoutInSeconds nullable.Type[int64]         `json:"ringingTimeoutInSeconds,omitempty"`
		RoutingPolicies         *[]RoutingPolicy             `json:"routingPolicies,omitempty"`
		Source                  *ParticipantInfo             `json:"source,omitempty"`
		State                   *CallState                   `json:"state,omitempty"`
		Subject                 nullable.Type[string]        `json:"subject,omitempty"`
		Targets                 *[]InvitationParticipantInfo `json:"targets,omitempty"`
		TenantId                nullable.Type[string]        `json:"tenantId,omitempty"`
		TerminationReason       nullable.Type[string]        `json:"terminationReason,omitempty"`
		ToneInfo                *ToneInfo                    `json:"toneInfo,omitempty"`
		Transcription           *CallTranscriptionInfo       `json:"transcription,omitempty"`
		Id                      *string                      `json:"id,omitempty"`
		ODataId                 *string                      `json:"@odata.id,omitempty"`
		ODataType               *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActiveModalities = decoded.ActiveModalities
	s.AnsweredBy = decoded.AnsweredBy
	s.AudioRoutingGroups = decoded.AudioRoutingGroups
	s.CallChainId = decoded.CallChainId
	s.CallRoutes = decoded.CallRoutes
	s.CallbackUri = decoded.CallbackUri
	s.ChatInfo = decoded.ChatInfo
	s.ContentSharingSessions = decoded.ContentSharingSessions
	s.Direction = decoded.Direction
	s.IncomingContext = decoded.IncomingContext
	s.MediaState = decoded.MediaState
	s.MeetingCapability = decoded.MeetingCapability
	s.MyParticipantId = decoded.MyParticipantId
	s.Participants = decoded.Participants
	s.RequestedModalities = decoded.RequestedModalities
	s.ResultInfo = decoded.ResultInfo
	s.RingingTimeoutInSeconds = decoded.RingingTimeoutInSeconds
	s.RoutingPolicies = decoded.RoutingPolicies
	s.Source = decoded.Source
	s.State = decoded.State
	s.Subject = decoded.Subject
	s.Targets = decoded.Targets
	s.TenantId = decoded.TenantId
	s.TerminationReason = decoded.TerminationReason
	s.ToneInfo = decoded.ToneInfo
	s.Transcription = decoded.Transcription
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Call into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["callOptions"]; ok {
		impl, err := UnmarshalCallOptionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CallOptions' for 'Call': %+v", err)
		}
		s.CallOptions = impl
	}

	if v, ok := temp["mediaConfig"]; ok {
		impl, err := UnmarshalMediaConfigImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MediaConfig' for 'Call': %+v", err)
		}
		s.MediaConfig = impl
	}

	if v, ok := temp["meetingInfo"]; ok {
		impl, err := UnmarshalMeetingInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MeetingInfo' for 'Call': %+v", err)
		}
		s.MeetingInfo = impl
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]CommsOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCommsOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'Call': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
