package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ITunesPairingMode string

const (
	ITunesPairingMode_Allow               ITunesPairingMode = "allow"
	ITunesPairingMode_Disallow            ITunesPairingMode = "disallow"
	ITunesPairingMode_RequiresCertificate ITunesPairingMode = "requiresCertificate"
)

func PossibleValuesForITunesPairingMode() []string {
	return []string{
		string(ITunesPairingMode_Allow),
		string(ITunesPairingMode_Disallow),
		string(ITunesPairingMode_RequiresCertificate),
	}
}

func (s *ITunesPairingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseITunesPairingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseITunesPairingMode(input string) (*ITunesPairingMode, error) {
	vals := map[string]ITunesPairingMode{
		"allow":               ITunesPairingMode_Allow,
		"disallow":            ITunesPairingMode_Disallow,
		"requirescertificate": ITunesPairingMode_RequiresCertificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ITunesPairingMode(input)
	return &out, nil
}
