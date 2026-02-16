package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailSyncSchedule string

const (
	EmailSyncSchedule_AsMessagesArrive EmailSyncSchedule = "asMessagesArrive"
	EmailSyncSchedule_BasedOnMyUsage   EmailSyncSchedule = "basedOnMyUsage"
	EmailSyncSchedule_FifteenMinutes   EmailSyncSchedule = "fifteenMinutes"
	EmailSyncSchedule_Manual           EmailSyncSchedule = "manual"
	EmailSyncSchedule_SixtyMinutes     EmailSyncSchedule = "sixtyMinutes"
	EmailSyncSchedule_ThirtyMinutes    EmailSyncSchedule = "thirtyMinutes"
	EmailSyncSchedule_UserDefined      EmailSyncSchedule = "userDefined"
)

func PossibleValuesForEmailSyncSchedule() []string {
	return []string{
		string(EmailSyncSchedule_AsMessagesArrive),
		string(EmailSyncSchedule_BasedOnMyUsage),
		string(EmailSyncSchedule_FifteenMinutes),
		string(EmailSyncSchedule_Manual),
		string(EmailSyncSchedule_SixtyMinutes),
		string(EmailSyncSchedule_ThirtyMinutes),
		string(EmailSyncSchedule_UserDefined),
	}
}

func (s *EmailSyncSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailSyncSchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailSyncSchedule(input string) (*EmailSyncSchedule, error) {
	vals := map[string]EmailSyncSchedule{
		"asmessagesarrive": EmailSyncSchedule_AsMessagesArrive,
		"basedonmyusage":   EmailSyncSchedule_BasedOnMyUsage,
		"fifteenminutes":   EmailSyncSchedule_FifteenMinutes,
		"manual":           EmailSyncSchedule_Manual,
		"sixtyminutes":     EmailSyncSchedule_SixtyMinutes,
		"thirtyminutes":    EmailSyncSchedule_ThirtyMinutes,
		"userdefined":      EmailSyncSchedule_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailSyncSchedule(input)
	return &out, nil
}
