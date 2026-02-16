package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailType string

const (
	EmailType_Main     EmailType = "main"
	EmailType_Other    EmailType = "other"
	EmailType_Personal EmailType = "personal"
	EmailType_Unknown  EmailType = "unknown"
	EmailType_Work     EmailType = "work"
)

func PossibleValuesForEmailType() []string {
	return []string{
		string(EmailType_Main),
		string(EmailType_Other),
		string(EmailType_Personal),
		string(EmailType_Unknown),
		string(EmailType_Work),
	}
}

func (s *EmailType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailType(input string) (*EmailType, error) {
	vals := map[string]EmailType{
		"main":     EmailType_Main,
		"other":    EmailType_Other,
		"personal": EmailType_Personal,
		"unknown":  EmailType_Unknown,
		"work":     EmailType_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailType(input)
	return &out, nil
}
