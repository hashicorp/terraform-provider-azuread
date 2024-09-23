package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEvidenceRole string

const (
	SecurityEvidenceRole_Added             SecurityEvidenceRole = "added"
	SecurityEvidenceRole_Attacked          SecurityEvidenceRole = "attacked"
	SecurityEvidenceRole_Attacker          SecurityEvidenceRole = "attacker"
	SecurityEvidenceRole_CommandAndControl SecurityEvidenceRole = "commandAndControl"
	SecurityEvidenceRole_Compromised       SecurityEvidenceRole = "compromised"
	SecurityEvidenceRole_Contextual        SecurityEvidenceRole = "contextual"
	SecurityEvidenceRole_Created           SecurityEvidenceRole = "created"
	SecurityEvidenceRole_Destination       SecurityEvidenceRole = "destination"
	SecurityEvidenceRole_Edited            SecurityEvidenceRole = "edited"
	SecurityEvidenceRole_Loaded            SecurityEvidenceRole = "loaded"
	SecurityEvidenceRole_PolicyViolator    SecurityEvidenceRole = "policyViolator"
	SecurityEvidenceRole_Scanned           SecurityEvidenceRole = "scanned"
	SecurityEvidenceRole_Source            SecurityEvidenceRole = "source"
	SecurityEvidenceRole_Suspicious        SecurityEvidenceRole = "suspicious"
	SecurityEvidenceRole_Unknown           SecurityEvidenceRole = "unknown"
)

func PossibleValuesForSecurityEvidenceRole() []string {
	return []string{
		string(SecurityEvidenceRole_Added),
		string(SecurityEvidenceRole_Attacked),
		string(SecurityEvidenceRole_Attacker),
		string(SecurityEvidenceRole_CommandAndControl),
		string(SecurityEvidenceRole_Compromised),
		string(SecurityEvidenceRole_Contextual),
		string(SecurityEvidenceRole_Created),
		string(SecurityEvidenceRole_Destination),
		string(SecurityEvidenceRole_Edited),
		string(SecurityEvidenceRole_Loaded),
		string(SecurityEvidenceRole_PolicyViolator),
		string(SecurityEvidenceRole_Scanned),
		string(SecurityEvidenceRole_Source),
		string(SecurityEvidenceRole_Suspicious),
		string(SecurityEvidenceRole_Unknown),
	}
}

func (s *SecurityEvidenceRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEvidenceRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEvidenceRole(input string) (*SecurityEvidenceRole, error) {
	vals := map[string]SecurityEvidenceRole{
		"added":             SecurityEvidenceRole_Added,
		"attacked":          SecurityEvidenceRole_Attacked,
		"attacker":          SecurityEvidenceRole_Attacker,
		"commandandcontrol": SecurityEvidenceRole_CommandAndControl,
		"compromised":       SecurityEvidenceRole_Compromised,
		"contextual":        SecurityEvidenceRole_Contextual,
		"created":           SecurityEvidenceRole_Created,
		"destination":       SecurityEvidenceRole_Destination,
		"edited":            SecurityEvidenceRole_Edited,
		"loaded":            SecurityEvidenceRole_Loaded,
		"policyviolator":    SecurityEvidenceRole_PolicyViolator,
		"scanned":           SecurityEvidenceRole_Scanned,
		"source":            SecurityEvidenceRole_Source,
		"suspicious":        SecurityEvidenceRole_Suspicious,
		"unknown":           SecurityEvidenceRole_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEvidenceRole(input)
	return &out, nil
}
