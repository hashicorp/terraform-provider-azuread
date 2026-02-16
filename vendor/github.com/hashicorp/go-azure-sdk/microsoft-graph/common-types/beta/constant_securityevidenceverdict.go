package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEvidenceVerdict string

const (
	SecurityEvidenceVerdict_Malicious      SecurityEvidenceVerdict = "malicious"
	SecurityEvidenceVerdict_NoThreatsFound SecurityEvidenceVerdict = "noThreatsFound"
	SecurityEvidenceVerdict_Suspicious     SecurityEvidenceVerdict = "suspicious"
	SecurityEvidenceVerdict_Unknown        SecurityEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityEvidenceVerdict() []string {
	return []string{
		string(SecurityEvidenceVerdict_Malicious),
		string(SecurityEvidenceVerdict_NoThreatsFound),
		string(SecurityEvidenceVerdict_Suspicious),
		string(SecurityEvidenceVerdict_Unknown),
	}
}

func (s *SecurityEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEvidenceVerdict(input string) (*SecurityEvidenceVerdict, error) {
	vals := map[string]SecurityEvidenceVerdict{
		"malicious":      SecurityEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityEvidenceVerdict_Suspicious,
		"unknown":        SecurityEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEvidenceVerdict(input)
	return &out, nil
}
