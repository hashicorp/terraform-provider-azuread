package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRequestType string

const (
	MeetingRequestType_FullUpdate          MeetingRequestType = "fullUpdate"
	MeetingRequestType_InformationalUpdate MeetingRequestType = "informationalUpdate"
	MeetingRequestType_NewMeetingRequest   MeetingRequestType = "newMeetingRequest"
	MeetingRequestType_None                MeetingRequestType = "none"
	MeetingRequestType_Outdated            MeetingRequestType = "outdated"
	MeetingRequestType_PrincipalWantsCopy  MeetingRequestType = "principalWantsCopy"
	MeetingRequestType_SilentUpdate        MeetingRequestType = "silentUpdate"
)

func PossibleValuesForMeetingRequestType() []string {
	return []string{
		string(MeetingRequestType_FullUpdate),
		string(MeetingRequestType_InformationalUpdate),
		string(MeetingRequestType_NewMeetingRequest),
		string(MeetingRequestType_None),
		string(MeetingRequestType_Outdated),
		string(MeetingRequestType_PrincipalWantsCopy),
		string(MeetingRequestType_SilentUpdate),
	}
}

func (s *MeetingRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeetingRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeetingRequestType(input string) (*MeetingRequestType, error) {
	vals := map[string]MeetingRequestType{
		"fullupdate":          MeetingRequestType_FullUpdate,
		"informationalupdate": MeetingRequestType_InformationalUpdate,
		"newmeetingrequest":   MeetingRequestType_NewMeetingRequest,
		"none":                MeetingRequestType_None,
		"outdated":            MeetingRequestType_Outdated,
		"principalwantscopy":  MeetingRequestType_PrincipalWantsCopy,
		"silentupdate":        MeetingRequestType_SilentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeetingRequestType(input)
	return &out, nil
}
