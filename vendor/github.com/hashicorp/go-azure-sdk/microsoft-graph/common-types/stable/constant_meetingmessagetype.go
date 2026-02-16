package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingMessageType string

const (
	MeetingMessageType_MeetingAccepted           MeetingMessageType = "meetingAccepted"
	MeetingMessageType_MeetingCancelled          MeetingMessageType = "meetingCancelled"
	MeetingMessageType_MeetingDeclined           MeetingMessageType = "meetingDeclined"
	MeetingMessageType_MeetingRequest            MeetingMessageType = "meetingRequest"
	MeetingMessageType_MeetingTenativelyAccepted MeetingMessageType = "meetingTenativelyAccepted"
	MeetingMessageType_None                      MeetingMessageType = "none"
)

func PossibleValuesForMeetingMessageType() []string {
	return []string{
		string(MeetingMessageType_MeetingAccepted),
		string(MeetingMessageType_MeetingCancelled),
		string(MeetingMessageType_MeetingDeclined),
		string(MeetingMessageType_MeetingRequest),
		string(MeetingMessageType_MeetingTenativelyAccepted),
		string(MeetingMessageType_None),
	}
}

func (s *MeetingMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingMessageType(input string) (*MeetingMessageType, error) {
	vals := map[string]MeetingMessageType{
		"meetingaccepted":           MeetingMessageType_MeetingAccepted,
		"meetingcancelled":          MeetingMessageType_MeetingCancelled,
		"meetingdeclined":           MeetingMessageType_MeetingDeclined,
		"meetingrequest":            MeetingMessageType_MeetingRequest,
		"meetingtenativelyaccepted": MeetingMessageType_MeetingTenativelyAccepted,
		"none":                      MeetingMessageType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingMessageType(input)
	return &out, nil
}
