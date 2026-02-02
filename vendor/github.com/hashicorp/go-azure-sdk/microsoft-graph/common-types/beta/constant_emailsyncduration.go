package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailSyncDuration string

const (
	EmailSyncDuration_OneDay      EmailSyncDuration = "oneDay"
	EmailSyncDuration_OneMonth    EmailSyncDuration = "oneMonth"
	EmailSyncDuration_OneWeek     EmailSyncDuration = "oneWeek"
	EmailSyncDuration_ThreeDays   EmailSyncDuration = "threeDays"
	EmailSyncDuration_TwoWeeks    EmailSyncDuration = "twoWeeks"
	EmailSyncDuration_Unlimited   EmailSyncDuration = "unlimited"
	EmailSyncDuration_UserDefined EmailSyncDuration = "userDefined"
)

func PossibleValuesForEmailSyncDuration() []string {
	return []string{
		string(EmailSyncDuration_OneDay),
		string(EmailSyncDuration_OneMonth),
		string(EmailSyncDuration_OneWeek),
		string(EmailSyncDuration_ThreeDays),
		string(EmailSyncDuration_TwoWeeks),
		string(EmailSyncDuration_Unlimited),
		string(EmailSyncDuration_UserDefined),
	}
}

func (s *EmailSyncDuration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailSyncDuration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailSyncDuration(input string) (*EmailSyncDuration, error) {
	vals := map[string]EmailSyncDuration{
		"oneday":      EmailSyncDuration_OneDay,
		"onemonth":    EmailSyncDuration_OneMonth,
		"oneweek":     EmailSyncDuration_OneWeek,
		"threedays":   EmailSyncDuration_ThreeDays,
		"twoweeks":    EmailSyncDuration_TwoWeeks,
		"unlimited":   EmailSyncDuration_Unlimited,
		"userdefined": EmailSyncDuration_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailSyncDuration(input)
	return &out, nil
}
