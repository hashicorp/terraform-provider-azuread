package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CoachmarkLocationType string

const (
	CoachmarkLocationType_DisplayName CoachmarkLocationType = "displayName"
	CoachmarkLocationType_ExternalTag CoachmarkLocationType = "externalTag"
	CoachmarkLocationType_FromEmail   CoachmarkLocationType = "fromEmail"
	CoachmarkLocationType_MessageBody CoachmarkLocationType = "messageBody"
	CoachmarkLocationType_Subject     CoachmarkLocationType = "subject"
	CoachmarkLocationType_Unknown     CoachmarkLocationType = "unknown"
)

func PossibleValuesForCoachmarkLocationType() []string {
	return []string{
		string(CoachmarkLocationType_DisplayName),
		string(CoachmarkLocationType_ExternalTag),
		string(CoachmarkLocationType_FromEmail),
		string(CoachmarkLocationType_MessageBody),
		string(CoachmarkLocationType_Subject),
		string(CoachmarkLocationType_Unknown),
	}
}

func (s *CoachmarkLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCoachmarkLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCoachmarkLocationType(input string) (*CoachmarkLocationType, error) {
	vals := map[string]CoachmarkLocationType{
		"displayname": CoachmarkLocationType_DisplayName,
		"externaltag": CoachmarkLocationType_ExternalTag,
		"fromemail":   CoachmarkLocationType_FromEmail,
		"messagebody": CoachmarkLocationType_MessageBody,
		"subject":     CoachmarkLocationType_Subject,
		"unknown":     CoachmarkLocationType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CoachmarkLocationType(input)
	return &out, nil
}
