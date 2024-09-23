package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEntityIdentifier string

const (
	SecurityFileEntityIdentifier_InitiatingProcessSHA1   SecurityFileEntityIdentifier = "initiatingProcessSHA1"
	SecurityFileEntityIdentifier_InitiatingProcessSHA256 SecurityFileEntityIdentifier = "initiatingProcessSHA256"
	SecurityFileEntityIdentifier_Sha1                    SecurityFileEntityIdentifier = "sha1"
	SecurityFileEntityIdentifier_Sha256                  SecurityFileEntityIdentifier = "sha256"
)

func PossibleValuesForSecurityFileEntityIdentifier() []string {
	return []string{
		string(SecurityFileEntityIdentifier_InitiatingProcessSHA1),
		string(SecurityFileEntityIdentifier_InitiatingProcessSHA256),
		string(SecurityFileEntityIdentifier_Sha1),
		string(SecurityFileEntityIdentifier_Sha256),
	}
}

func (s *SecurityFileEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileEntityIdentifier(input string) (*SecurityFileEntityIdentifier, error) {
	vals := map[string]SecurityFileEntityIdentifier{
		"initiatingprocesssha1":   SecurityFileEntityIdentifier_InitiatingProcessSHA1,
		"initiatingprocesssha256": SecurityFileEntityIdentifier_InitiatingProcessSHA256,
		"sha1":                    SecurityFileEntityIdentifier_Sha1,
		"sha256":                  SecurityFileEntityIdentifier_Sha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEntityIdentifier(input)
	return &out, nil
}
