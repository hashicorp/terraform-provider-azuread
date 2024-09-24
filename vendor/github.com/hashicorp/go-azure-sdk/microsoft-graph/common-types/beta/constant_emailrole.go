package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailRole string

const (
	EmailRole_Recipient EmailRole = "recipient"
	EmailRole_Sender    EmailRole = "sender"
	EmailRole_Unknown   EmailRole = "unknown"
)

func PossibleValuesForEmailRole() []string {
	return []string{
		string(EmailRole_Recipient),
		string(EmailRole_Sender),
		string(EmailRole_Unknown),
	}
}

func (s *EmailRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailRole(input string) (*EmailRole, error) {
	vals := map[string]EmailRole{
		"recipient": EmailRole_Recipient,
		"sender":    EmailRole_Sender,
		"unknown":   EmailRole_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailRole(input)
	return &out, nil
}
