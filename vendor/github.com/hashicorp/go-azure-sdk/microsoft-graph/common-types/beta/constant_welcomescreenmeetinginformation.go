package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WelcomeScreenMeetingInformation string

const (
	WelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject WelcomeScreenMeetingInformation = "showOrganizerAndTimeAndSubject"
	WelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly       WelcomeScreenMeetingInformation = "showOrganizerAndTimeOnly"
	WelcomeScreenMeetingInformation_UserDefined                    WelcomeScreenMeetingInformation = "userDefined"
)

func PossibleValuesForWelcomeScreenMeetingInformation() []string {
	return []string{
		string(WelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject),
		string(WelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly),
		string(WelcomeScreenMeetingInformation_UserDefined),
	}
}

func (s *WelcomeScreenMeetingInformation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWelcomeScreenMeetingInformation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWelcomeScreenMeetingInformation(input string) (*WelcomeScreenMeetingInformation, error) {
	vals := map[string]WelcomeScreenMeetingInformation{
		"showorganizerandtimeandsubject": WelcomeScreenMeetingInformation_ShowOrganizerAndTimeAndSubject,
		"showorganizerandtimeonly":       WelcomeScreenMeetingInformation_ShowOrganizerAndTimeOnly,
		"userdefined":                    WelcomeScreenMeetingInformation_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WelcomeScreenMeetingInformation(input)
	return &out, nil
}
