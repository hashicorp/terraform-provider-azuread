package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileHashAlgorithm string

const (
	SecurityFileHashAlgorithm_Md5      SecurityFileHashAlgorithm = "md5"
	SecurityFileHashAlgorithm_Sha1     SecurityFileHashAlgorithm = "sha1"
	SecurityFileHashAlgorithm_Sha256   SecurityFileHashAlgorithm = "sha256"
	SecurityFileHashAlgorithm_Sha256ac SecurityFileHashAlgorithm = "sha256ac"
	SecurityFileHashAlgorithm_Unknown  SecurityFileHashAlgorithm = "unknown"
)

func PossibleValuesForSecurityFileHashAlgorithm() []string {
	return []string{
		string(SecurityFileHashAlgorithm_Md5),
		string(SecurityFileHashAlgorithm_Sha1),
		string(SecurityFileHashAlgorithm_Sha256),
		string(SecurityFileHashAlgorithm_Sha256ac),
		string(SecurityFileHashAlgorithm_Unknown),
	}
}

func (s *SecurityFileHashAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileHashAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileHashAlgorithm(input string) (*SecurityFileHashAlgorithm, error) {
	vals := map[string]SecurityFileHashAlgorithm{
		"md5":      SecurityFileHashAlgorithm_Md5,
		"sha1":     SecurityFileHashAlgorithm_Sha1,
		"sha256":   SecurityFileHashAlgorithm_Sha256,
		"sha256ac": SecurityFileHashAlgorithm_Sha256ac,
		"unknown":  SecurityFileHashAlgorithm_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileHashAlgorithm(input)
	return &out, nil
}
