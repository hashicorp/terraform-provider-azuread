package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsServiceRole string

const (
	CallRecordsServiceRole_AudioTeleconferencerController                          CallRecordsServiceRole = "audioTeleconferencerController"
	CallRecordsServiceRole_ConferencingAnnouncementService                         CallRecordsServiceRole = "conferencingAnnouncementService"
	CallRecordsServiceRole_ConferencingAttendant                                   CallRecordsServiceRole = "conferencingAttendant"
	CallRecordsServiceRole_CustomBot                                               CallRecordsServiceRole = "customBot"
	CallRecordsServiceRole_ExchangeUnifiedMessagingService                         CallRecordsServiceRole = "exchangeUnifiedMessagingService"
	CallRecordsServiceRole_Gateway                                                 CallRecordsServiceRole = "gateway"
	CallRecordsServiceRole_MediaController                                         CallRecordsServiceRole = "mediaController"
	CallRecordsServiceRole_MediationServer                                         CallRecordsServiceRole = "mediationServer"
	CallRecordsServiceRole_MediationServerCloudConnectorEdition                    CallRecordsServiceRole = "mediationServerCloudConnectorEdition"
	CallRecordsServiceRole_ResponseGroupService                                    CallRecordsServiceRole = "responseGroupService"
	CallRecordsServiceRole_ResponseGroupServiceAnnouncementService                 CallRecordsServiceRole = "responseGroupServiceAnnouncementService"
	CallRecordsServiceRole_SkypeForBusinessApplicationSharingMcu                   CallRecordsServiceRole = "skypeForBusinessApplicationSharingMcu"
	CallRecordsServiceRole_SkypeForBusinessAttendant                               CallRecordsServiceRole = "skypeForBusinessAttendant"
	CallRecordsServiceRole_SkypeForBusinessAudioVideoMcu                           CallRecordsServiceRole = "skypeForBusinessAudioVideoMcu"
	CallRecordsServiceRole_SkypeForBusinessAutoAttendant                           CallRecordsServiceRole = "skypeForBusinessAutoAttendant"
	CallRecordsServiceRole_SkypeForBusinessCallQueues                              CallRecordsServiceRole = "skypeForBusinessCallQueues"
	CallRecordsServiceRole_SkypeForBusinessMicrosoftTeamsGateway                   CallRecordsServiceRole = "skypeForBusinessMicrosoftTeamsGateway"
	CallRecordsServiceRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform CallRecordsServiceRole = "skypeForBusinessUnifiedCommunicationApplicationPlatform"
	CallRecordsServiceRole_SkypeTranslator                                         CallRecordsServiceRole = "skypeTranslator"
	CallRecordsServiceRole_Unknown                                                 CallRecordsServiceRole = "unknown"
	CallRecordsServiceRole_Voicemail                                               CallRecordsServiceRole = "voicemail"
)

func PossibleValuesForCallRecordsServiceRole() []string {
	return []string{
		string(CallRecordsServiceRole_AudioTeleconferencerController),
		string(CallRecordsServiceRole_ConferencingAnnouncementService),
		string(CallRecordsServiceRole_ConferencingAttendant),
		string(CallRecordsServiceRole_CustomBot),
		string(CallRecordsServiceRole_ExchangeUnifiedMessagingService),
		string(CallRecordsServiceRole_Gateway),
		string(CallRecordsServiceRole_MediaController),
		string(CallRecordsServiceRole_MediationServer),
		string(CallRecordsServiceRole_MediationServerCloudConnectorEdition),
		string(CallRecordsServiceRole_ResponseGroupService),
		string(CallRecordsServiceRole_ResponseGroupServiceAnnouncementService),
		string(CallRecordsServiceRole_SkypeForBusinessApplicationSharingMcu),
		string(CallRecordsServiceRole_SkypeForBusinessAttendant),
		string(CallRecordsServiceRole_SkypeForBusinessAudioVideoMcu),
		string(CallRecordsServiceRole_SkypeForBusinessAutoAttendant),
		string(CallRecordsServiceRole_SkypeForBusinessCallQueues),
		string(CallRecordsServiceRole_SkypeForBusinessMicrosoftTeamsGateway),
		string(CallRecordsServiceRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform),
		string(CallRecordsServiceRole_SkypeTranslator),
		string(CallRecordsServiceRole_Unknown),
		string(CallRecordsServiceRole_Voicemail),
	}
}

func (s *CallRecordsServiceRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsServiceRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsServiceRole(input string) (*CallRecordsServiceRole, error) {
	vals := map[string]CallRecordsServiceRole{
		"audioteleconferencercontroller":                          CallRecordsServiceRole_AudioTeleconferencerController,
		"conferencingannouncementservice":                         CallRecordsServiceRole_ConferencingAnnouncementService,
		"conferencingattendant":                                   CallRecordsServiceRole_ConferencingAttendant,
		"custombot":                                               CallRecordsServiceRole_CustomBot,
		"exchangeunifiedmessagingservice":                         CallRecordsServiceRole_ExchangeUnifiedMessagingService,
		"gateway":                                                 CallRecordsServiceRole_Gateway,
		"mediacontroller":                                         CallRecordsServiceRole_MediaController,
		"mediationserver":                                         CallRecordsServiceRole_MediationServer,
		"mediationservercloudconnectoredition":                    CallRecordsServiceRole_MediationServerCloudConnectorEdition,
		"responsegroupservice":                                    CallRecordsServiceRole_ResponseGroupService,
		"responsegroupserviceannouncementservice":                 CallRecordsServiceRole_ResponseGroupServiceAnnouncementService,
		"skypeforbusinessapplicationsharingmcu":                   CallRecordsServiceRole_SkypeForBusinessApplicationSharingMcu,
		"skypeforbusinessattendant":                               CallRecordsServiceRole_SkypeForBusinessAttendant,
		"skypeforbusinessaudiovideomcu":                           CallRecordsServiceRole_SkypeForBusinessAudioVideoMcu,
		"skypeforbusinessautoattendant":                           CallRecordsServiceRole_SkypeForBusinessAutoAttendant,
		"skypeforbusinesscallqueues":                              CallRecordsServiceRole_SkypeForBusinessCallQueues,
		"skypeforbusinessmicrosoftteamsgateway":                   CallRecordsServiceRole_SkypeForBusinessMicrosoftTeamsGateway,
		"skypeforbusinessunifiedcommunicationapplicationplatform": CallRecordsServiceRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform,
		"skypetranslator":                                         CallRecordsServiceRole_SkypeTranslator,
		"unknown":                                                 CallRecordsServiceRole_Unknown,
		"voicemail":                                               CallRecordsServiceRole_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsServiceRole(input)
	return &out, nil
}
